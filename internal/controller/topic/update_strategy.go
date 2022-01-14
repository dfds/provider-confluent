package topic

import (
	"fmt"
	"strconv"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients/topic"
)

type TopicCompare struct {
	TopicNamesMatch   bool
	ClusterMatch      bool
	EnvironmentMatch  bool
	PartitionsMatch   bool
	PartitionDecrease bool
	ConfigMatch       bool
}

func updateStrategy(tp v1alpha1.TopicParameters, td topic.TopicDescribeResponse, to v1alpha1.TopicObservation) (TopicCompare, error) {
	var compare TopicCompare

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

	fmt.Println("Observed:", td.Config.NumPartitions)
	fmt.Println("Spec:", tp.Topic.Partitions)
	numPartitions, err := strconv.Atoi(td.Config.NumPartitions)
	if err != nil {
		return compare, err
	}
	fmt.Println("Converted observed:", numPartitions)

	fmt.Println("ARE WE FUCKEDDDDDD:", tp.Topic.Partitions == numPartitions)
	if tp.Topic.Partitions == numPartitions {
		compare.PartitionsMatch = true

	} else {
		if tp.Topic.Partitions < numPartitions {
			compare.PartitionDecrease = true
		}
		compare.PartitionsMatch = false
	}

	return compare, nil
}

func (tc *TopicCompare) IsDestructive() bool {
	fmt.Println("WWHHHATTTT IS GOING ON:", tc)
	if !(tc.ClusterMatch || tc.EnvironmentMatch || tc.TopicNamesMatch || tc.PartitionsMatch) {
		return true
	}
	return true
}
