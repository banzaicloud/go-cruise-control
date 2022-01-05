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
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rightsize", Label("api:rightsize", "api:user_tasks", "api:state"),
	Serial,
	func() {

		BeforeEach(func() {
			By("waiting until Cruise Control is ready")
			Eventually(func() bool {
				ready, err := helpers.IsCruiseControlReady(ctx, cruisecontrol)
				Expect(err).NotTo(HaveOccurred())
				return ready
			}, CruiseControlReadyTimeout, 15).Should(BeTrue())
		})

		Describe("Rightsizing topic in Kafka cluster", func() {
			Context("by increasing partition count", func() {
				It("should return no error", func() {
					By("sending rightsize request to Cruise Control")
					req := api.RightsizeRequestWithDefaults()
					req.Topic = AirPortsTopicName
					req.PartitionCount = AirportsTopicNewPartitionSize

					resp, err := cruisecontrol.Rightsize(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.Failed()).To(BeFalse())

					By("waiting until the rightsize task finished")
					Eventually(func() bool {
						finished, err := helpers.HasUserTaskFinished(ctx, cruisecontrol, resp.TaskID)
						Expect(err).NotTo(HaveOccurred())
						return finished
					}, 300, 15).Should(BeTrue())

				})
			})
		})
	})
