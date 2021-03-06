package topic

import (
	"strconv"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/topic"
)

// Compare helper struct
type Compare struct {
	TopicNamesMatch  bool
	ClusterMatch     bool
	EnvironmentMatch bool
	PartitionsMatch  bool
	// PartitionDecrease bool
	ConfigMatch bool
}

func updateStrategy(tp v1alpha1.TopicParameters, td topic.DescribeResponse, to v1alpha1.TopicObservation) (Compare, error) {
	var compare Compare

	if tp.Topic.Name == td.TopicName {
		compare.TopicNamesMatch = true
	}

	if tp.Cluster == to.Cluster {
		compare.ClusterMatch = true
	}

	if tp.Environment == to.Environment {
		compare.EnvironmentMatch = true
	}

	if strconv.FormatInt(tp.Topic.Config.Retention, 10) == td.Config.RetentionMs {
		compare.ConfigMatch = true
	}

	numPartitions, err := strconv.Atoi(td.Config.NumPartitions)
	if err != nil {
		return compare, err
	}

	if tp.Topic.Partitions == numPartitions {
		compare.PartitionsMatch = true

	} else {
		// if tp.Topic.Partitions < numPartitions {
		// 	compare.PartitionDecrease = true
		// }
		compare.PartitionsMatch = false
	}

	return compare, nil
}

// IsDestructive helper method to determine destructive behaviour
func (tc *Compare) IsDestructive() bool {
	isDestructive := false
	if !tc.ClusterMatch {
		isDestructive = true
	}

	if !tc.EnvironmentMatch {
		isDestructive = true
	}

	if !tc.TopicNamesMatch {
		isDestructive = true
	}

	if !tc.PartitionsMatch {
		isDestructive = true
	}

	return isDestructive
}

// DestructiveActionsAllowed helper method to deny destructive actions
func DestructiveActionsAllowed(dp v1.DeletionPolicy) bool {
	return dp == "Delete"
}
