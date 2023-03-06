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

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"

	"github.com/banzaicloud/go-cruise-control/pkg/types"
)

const (
	EndpointRemoveBroker types.APIEndpoint = "REMOVE_BROKER"
)

type RemoveBrokerRequest struct {
	types.GenericRequestWithReason

	// Whether to allow capacity estimation when cruise-control is unable to obtain all per-broker capacity information
	AllowCapacityEstimation bool `param:"allow_capacity_estimation"`
	// List of target broker ids
	BrokerIDs []int32 `param:"brokerid"`
	// The upper bound of ongoing leadership movements
	ConcurrentLeaderMovements int32 `param:"concurrent_leader_movements,omitempty"`
	// The upper bound of ongoing replica movements going into/out of each broker
	ConcurrentPartitionMovementsPerBroker int32 `param:"concurrent_partition_movements_per_broker,omitempty"`
	// Whether to calculate proposal from available valid partitions or valid windows
	DataFrom types.ProposalDataSource `param:"data_from,omitempty"`
	// List of target broker ids
	DestinationBrokerIDs []int32 `param:"destination_broker_ids,omitempty"`
	// Whether to dry-run the request or not.
	DryRun bool `param:"dryrun"`
	// Whether to allow leader replicas to be moved to recently demoted brokers
	ExcludeRecentlyDemotedBrokers bool `param:"exclude_recently_demoted_brokers,omitempty"`
	// Whether to allow replicas to be moved to recently removed broker
	ExcludeRecentlyRemovedBrokers bool `param:"exclude_recently_removed_brokers,omitempty"`
	// Specify topic whose partition is excluded from replica movement
	ExcludedTopics string `param:"excluded_topics,omitempty"`
	// Execution progress check interval in milliseconds
	ExecutionProgressCheckIntervalMs int64 `param:"execution_progress_check_interval_ms,omitempty"`
	// True to compute proposals in fast mode, false otherwise
	FastMode bool `param:"fast_mode,omitempty"`
	// List of goals used to generate proposal, the default goals will be used if this parameter is not specified
	Goals []types.Goal `param:"goals,omitempty"`
	// Whether to use Kafka assigner mode to generate proposals
	KafkaAssigner bool `param:"kafka_assigner,omitempty"`
	// Change upper bound of ongoing inter broker partition movements in cluster
	MaxPartitionMovementsInCluster int32 `param:"max_partition_movements_in_cluster,omitempty"`
	// Replica movement strategies to use
	ReplicaMovementStrategies []types.ReplicaMovementStrategy `param:"replica_movement_strategies,omitempty"`
	// Upper bound on the bandwidth in bytes per second used to move replicas
	ReplicationThrottle int64 `param:"replication_throttle,omitempty"`
	// Review id for 2-step verification
	ReviewID int32 `param:"review_id,omitempty"`
	// Whether to allow hard goals be skipped in proposal generation
	SkipHardGoalCheck bool `param:"skip_hard_goal_check,omitempty"`
	// Whether to stop the ongoing execution (if any) and start executing the given request
	StopOngoingExecution bool `param:"stop_ongoing_execution,omitempty"`
	// Whether to only use ready goals to generate proposal
	UseReadyDefaultGoals bool `param:"use_ready_default_goals,omitempty"`
	// Return detailed state information
	Verbose bool `param:"verbose,omitempty"`
	// Whether to throttle the removed broker
	ThrottleRemovedBroker bool `param:"throttle_removed_broker,omitempty"`
}

func (s RemoveBrokerRequest) Validate() error {
	if len(s.BrokerIDs) < 1 {
		return errors.New("list of brokers must not be empty")
	}

	if s.ConcurrentPartitionMovementsPerBroker == 0 {
		return errors.New("number of concurrent partition movements per broker must be bigger then 0")
	}

	if s.ConcurrentLeaderMovements == 0 {
		return errors.New("number of concurrent leader partition movements must be bigger then 0")
	}

	if s.MaxPartitionMovementsInCluster == 0 {
		return errors.New("the maximum number of partition movements must be bigger then 0")
	}

	return nil
}

func RemoveBrokerRequestWithDefaults() *RemoveBrokerRequest {
	return &RemoveBrokerRequest{
		AllowCapacityEstimation:               true,
		ConcurrentLeaderMovements:             defaultConcurrentLeaderMovements,
		ConcurrentPartitionMovementsPerBroker: defaultConcurrentPartitionMovementsPerBroker,
		ExecutionProgressCheckIntervalMs:      defaultExecutionProgressCheckIntervalMs,
		ThrottleRemovedBroker:                 true,
		DataFrom:                              types.ProposalDataSourceValidWindows,
	}
}

type RemoveBrokerResponse struct {
	types.GenericResponse

	Result *types.OptimizationResult
}

func (r *RemoveBrokerResponse) UnmarshalResponse(resp *http.Response) error {
	if err := r.GenericResponse.UnmarshalResponse(resp); err != nil {
		return fmt.Errorf("failed to parse HTTP response metadata: %w", err)
	}

	var bodyBytes []byte
	var err error

	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	var d interface{}
	switch resp.StatusCode {
	case http.StatusOK:
		r.Result = &types.OptimizationResult{}
		d = r.Result
	case http.StatusAccepted:
		r.Progress = &types.ProgressResult{}
		d = r.Progress
	default:
		r.Error = &types.APIError{}
		d = r.Error
	}

	if err = json.Unmarshal(bodyBytes, d); err != nil {
		return fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return nil
}
