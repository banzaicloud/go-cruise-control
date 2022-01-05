/*
Copyright © 2021 Cisco and/or its affiliates. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package envtest

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"github.com/docker/cli/cli/config"
	"github.com/docker/compose/v2/pkg/api"
	"github.com/docker/compose/v2/pkg/compose"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

const (
	cruiseControlContainerName = "cruisecontrol"

	stateRunning       = "running"
	healthStateHealthy = "healthy"
)

type Environment struct {
	ctx      context.Context
	composer api.Service
	project  *types.Project
	reuse    bool
}

func (e *Environment) Start() error {
	if e.reuse {
		return nil
	}
	opts := api.UpOptions{
		Create: api.CreateOptions{
			RemoveOrphans: true,
			QuietPull:     true,
		},
		Start: api.StartOptions{
			// Attach: formatter.NewLogConsumer(ctx, os.Stdout, false, false),
			Wait: true,
		},
	}
	return e.composer.Up(e.ctx, e.project, opts)
}

func (e *Environment) Stop() error {
	if e.reuse {
		return nil
	}
	opts := api.DownOptions{
		RemoveOrphans: true,
		Project:       e.project,
		Volumes:       true,
	}
	return e.composer.Down(e.ctx, e.project.Name, opts)
}

func (e *Environment) StartService(service string) error {
	services := []string{service}
	ps, err := e.composer.Ps(e.ctx, e.project.Name, api.PsOptions{Services: services})
	if err != nil {
		return err
	}
	if len(ps) == 1 && ps[0].State == "running" {
		return nil
	}

	timeout := 5 * time.Second
	opts := api.RestartOptions{
		Timeout:  &timeout,
		Services: services,
	}
	return e.composer.Restart(e.ctx, e.project, opts)
}

func (e *Environment) StopService(service string) error {
	timeout := 5 * time.Second
	opts := api.StopOptions{
		Timeout:  &timeout,
		Services: []string{service},
	}

	return e.composer.Stop(e.ctx, e.project, opts)
}

func (e Environment) Ready() (bool, error) {
	services := make([]string, len(e.project.Services))
	for i, srv := range e.project.Services {
		services[i] = srv.Name
	}

	ps, err := e.composer.Ps(e.ctx, e.project.Name, api.PsOptions{Services: services})
	if err != nil {
		return false, err
	}

	if len(services) != len(ps) {
		return false, errors.Errorf("there are missing containers in the environment. expected %d, got %d\n",
			len(services), len(ps))
	}

	for _, c := range ps {
		if c.State != stateRunning && c.Health != healthStateHealthy {
			return false, nil
		}
	}
	return true, nil
}

func (e Environment) CruiseControlURL() (*url.URL, error) {
	var cruiseControl types.ServiceConfig
	var ok bool

	for _, srv := range e.project.Services {
		if srv.Name == cruiseControlContainerName {
			cruiseControl = srv
			ok = true
			break
		}
	}

	if !ok {
		return nil, errors.Errorf("container with name %s is not available", cruiseControlContainerName)
	}

	if len(cruiseControl.Ports) < 1 {
		return nil, errors.Errorf("container with name %s has no ports published", cruiseControlContainerName)
	}

	port := cruiseControl.Ports[0].Published
	hostIP := cruiseControl.Ports[0].HostIP
	if hostIP == "" {
		hostIP = "127.0.0.1"
	}

	return &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%d", hostIP, port),
		Path:   "kafkacruisecontrol",
	}, nil
}

func New(ctx context.Context, o *cli.ProjectOptions, reuse bool) (*Environment, error) {
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	dockerConfig, err := config.Load(config.Dir())
	if err != nil {
		return nil, err
	}

	project, err := cli.ProjectFromOptions(o)
	if err != nil {
		return nil, err
	}

	return &Environment{
		ctx:      ctx,
		composer: compose.NewComposeService(dockerClient, dockerConfig),
		project:  project,
		reuse:    reuse,
	}, nil
}
