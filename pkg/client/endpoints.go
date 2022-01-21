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

package client

import (
	"net/http"
	"time"

	"github.com/banzaicloud/go-cruise-control/pkg/api"
)

func (c *Client) AddBroker(r *api.AddBrokerRequest) (*api.AddBrokerResponse, error) {
	resp := &api.AddBrokerResponse{}
	return resp, c.request(r, resp, api.EndpointAddBroker, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) Admin(r *api.AdminRequest) (*api.AdminResponse, error) {
	resp := &api.AdminResponse{}
	return resp, c.request(r, resp, api.EndpointAdmin, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) Bootstrap(r *api.BootstrapRequest) (*api.BootstrapResponse, error) {
	resp := &api.BootstrapResponse{}
	return resp, c.request(r, resp, api.EndpointBootstrap, http.MethodGet, 5*time.Minute)
}

func (c *Client) DemoteBroker(r *api.DemoteBrokerRequest) (*api.DemoteBrokerResponse, error) {
	resp := &api.DemoteBrokerResponse{}
	return resp, c.request(r, resp, api.EndpointDemoteBroker, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) FixOfflineReplicas(r *api.FixOfflineReplicasRequest) (*api.FixOfflineReplicasResponse, error) {
	resp := &api.FixOfflineReplicasResponse{}
	return resp, c.request(r, resp, api.EndpointFixOfflineReplicas, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) KafkaClusterLoad(r *api.KafkaClusterLoadRequest) (*api.KafkaClusterLoadResponse, error) {
	resp := &api.KafkaClusterLoadResponse{}
	return resp, c.request(r, resp, api.EndpointKafkaClusterLoad, http.MethodGet, DefaultRequestTimeout)
}

func (c *Client) KafkaClusterState(r *api.KafkaClusterStateRequest) (*api.KafkaClusterStateResponse, error) {
	resp := &api.KafkaClusterStateResponse{}
	return resp, c.request(r, resp, api.EndpointKafkaClusterState, http.MethodGet, DefaultRequestTimeout)
}

func (c *Client) KafkaPartitionLoad(r *api.KafkaPartitionLoadRequest) (*api.KafkaPartitionLoadResponse, error) {
	resp := &api.KafkaPartitionLoadResponse{}
	return resp, c.request(r, resp, api.EndpointKafkaPartitionLoad, http.MethodGet, DefaultRequestTimeout)
}

func (c *Client) PauseSampling(r *api.PauseSamplingRequest) (*api.PauseSamplingResponse, error) {
	resp := &api.PauseSamplingResponse{}
	return resp, c.request(r, resp, api.EndpointPauseSampling, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) Proposals(r *api.ProposalsRequest) (*api.ProposalsResponse, error) {
	resp := &api.ProposalsResponse{}
	return resp, c.request(r, resp, api.EndpointProposals, http.MethodGet, DefaultRequestTimeout)
}

func (c *Client) Rebalance(r *api.RebalanceRequest) (*api.RebalanceResponse, error) {
	resp := &api.RebalanceResponse{}
	return resp, c.request(r, resp, api.EndpointRebalance, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) RemoveBroker(r *api.RemoveBrokerRequest) (*api.RemoveBrokerResponse, error) {
	resp := &api.RemoveBrokerResponse{}
	return resp, c.request(r, resp, api.EndpointRemoveBroker, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) ResumeSampling(r *api.ResumeSamplingRequest) (*api.ResumeSamplingResponse, error) {
	resp := &api.ResumeSamplingResponse{}
	return resp, c.request(r, resp, api.EndpointResumeSampling, http.MethodPost, DefaultRequestTimeout)
}

func (c *Client) Review(r *api.ReviewRequest) (*api.ReviewResponse, error) {
	resp := &api.ReviewResponse{}
	return resp, c.request(r, resp, api.EndpointReview, http.MethodPost, DefaultRequestTimeout)
}

// ReviewBoard returns a list of Cruise Control requests with their review state.
func (c *Client) ReviewBoard(r *api.ReviewBoardRequest) (*api.ReviewBoardResponse, error) {
	resp := &api.ReviewBoardResponse{}
	return resp, c.request(r, resp, api.EndpointReviewBoard, http.MethodGet, DefaultRequestTimeout)
}

// Rightsize allows manually invoke provisioner rightsizing of the cluster.
func (c *Client) Rightsize(r *api.RightsizeRequest) (*api.RightsizeResponse, error) {
	resp := &api.RightsizeResponse{}
	return resp, c.request(r, resp, api.EndpointRightsize, http.MethodPost, DefaultRequestTimeout)
}

// State reports back the Cruise Control state.
func (c *Client) State(r *api.StateRequest) (*api.StateResponse, error) {
	resp := &api.StateResponse{}
	return resp, c.request(r, resp, api.EndpointState, http.MethodGet, DefaultRequestTimeout)
}

// StopProposalExecution invoke stopping of ongoing proposal execution in Cruise Control.
func (c *Client) StopProposalExecution(r *api.StopProposalExecutionRequest) (*api.StopProposalExecutionResponse, error) {
	resp := &api.StopProposalExecutionResponse{}
	return resp, c.request(r, resp, api.EndpointStopProposalExecution, http.MethodPost, DefaultRequestTimeout)
}

// TopicConfiguration allows changing Kafka topic configuration using Cruise Control.
func (c *Client) TopicConfiguration(r *api.TopicConfigurationRequest) (*api.TopicConfigurationResponse, error) {
	resp := &api.TopicConfigurationResponse{}
	return resp, c.request(r, resp, api.EndpointTopicConfiguration, http.MethodPost, DefaultRequestTimeout)
}

// Train Cruise Control to better model broker cpu usage.
func (c *Client) Train(r *api.TrainRequest) (*api.TrainResponse, error) {
	resp := &api.TrainResponse{}
	return resp, c.request(r, resp, api.EndpointTrain, http.MethodGet, DefaultRequestTimeout)
}

// UserTasks returns the list of recent tasks performed by Cruise Control.
func (c *Client) UserTasks(r *api.UserTasksRequest) (*api.UserTasksResponse, error) {
	resp := &api.UserTasksResponse{}
	return resp, c.request(r, resp, api.EndpointUserTasks, http.MethodGet, DefaultRequestTimeout)
}
