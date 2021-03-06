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
	"io"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(out io.Writer) (logr.Logger, func() error, error) {
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	sink := zapcore.AddSync(out)
	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	core := zapcore.NewCore(encoder, sink, level)
	logger := zap.New(core).WithOptions(
		zap.ErrorOutput(sink),
		zap.Development(),
	)
	log := zapr.NewLogger(logger)
	return log, logger.Sync, nil
}
