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

package integration_test

import (
	"github.com/banzaicloud/go-cruise-control/integration_test/helpers"
	"github.com/banzaicloud/go-cruise-control/pkg/api"
	"github.com/banzaicloud/go-cruise-control/pkg/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resume Sampling", Label("api:pause_sampling", "api:resume_sampling", "api:state"), func() {

	BeforeEach(func() {
		By("waiting until Cruise Control is ready")
		Eventually(func() bool {
			ready, err := helpers.IsCruiseControlReady(ctx, cruisecontrol)
			Expect(err).NotTo(HaveOccurred())
			return ready
		}, CruiseControlReadyTimeout, 15).Should(BeTrue())
	})

	Describe("Resuming metric sampling in Cruise Control", func() {
		Context("after it has been paused", func() {
			It("should return no error", func() {
				By("pausing sampling")
				req := api.PauseSamplingRequestWithDefaults()
				req.Reason = "integration testing"
				resp, err := cruisecontrol.PauseSampling(ctx, req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.Failed()).To(BeFalse())

				By("waiting until monitor state is changed to paused")
				Eventually(func() bool {
					req2 := api.StateRequestWithDefaults()
					req2.Substates = []types.Substate{
						types.SubstateMonitor,
					}

					resp2, err := cruisecontrol.State(ctx, req2)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp2.Failed()).To(BeFalse())
					state := resp2.Result.MonitorState.State
					log.V(0).Info("cruise control monitor state", "state", state)

					return state == types.MonitorStatePaused
				}, 300, 15).Should(BeTrue())

				By("resuming metric sampling Cruise Control")
				req2 := api.ResumeSamplingRequestWithDefaults()
				req2.Reason = "integration testing"
				resp2, err := cruisecontrol.ResumeSampling(ctx, req2)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp2.Failed()).To(BeFalse())

				By("waiting until monitor state is changed to running")
				Eventually(func() bool {
					req3 := api.StateRequestWithDefaults()
					req3.Substates = []types.Substate{
						types.SubstateMonitor,
					}

					resp3, err := cruisecontrol.State(ctx, req3)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp3.Failed()).To(BeFalse())
					state := resp3.Result.MonitorState.State
					log.V(0).Info("cruise control monitor state", "state", state)

					return state == types.MonitorStateRunning
				}, 300, 15).Should(BeTrue())

			})
		})
	})
})
