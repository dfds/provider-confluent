package topic

import (
	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
)

// IClient interface for service account client
type IClient interface {
	TopicCreate(tp v1alpha1.TopicParameters) error
	TopicDelete(tp v1alpha1.TopicParameters) error
	TopicDescribe(to v1alpha1.TopicObservation) (DescribeResponse, error)
	TopicUpdate(tp v1alpha1.TopicParameters) error
}

// Config is a configuration element for the service account client
type Config struct {
	APICredentials clients.APICredentials
}

// Client is a struct for service account client
type Client struct {
	Config Config
}

// DescribeResponse is a struct used for deserialising the response of TopicDescribe
type DescribeResponse struct {
	TopicName string `json:"topic_name"`
	Config    struct {
		CleanupPolicy                        string `json:"cleanup.policy"`
		CompressionType                      string `json:"compression.type"`
		DeleteRetentionMs                    string `json:"delete.retention.ms"`
		FileDeleteDelayMs                    string `json:"file.delete.delay.ms"`
		FlushMessages                        string `json:"flush.messages"`
		FlushMs                              string `json:"flush.ms"`
		FollowerReplicationThrottledReplicas string `json:"follower.replication.throttled.replicas"`
		IndexIntervalBytes                   string `json:"index.interval.bytes"`
		LeaderReplicationThrottledReplicas   string `json:"leader.replication.throttled.replicas"`
		MaxCompactionLagMs                   string `json:"max.compaction.lag.ms"`
		MaxMessageBytes                      string `json:"max.message.bytes"`
		MessageDownconversionEnable          string `json:"message.downconversion.enable"`
		MessageFormatVersion                 string `json:"message.format.version"`
		MessageTimestampDifferenceMaxMs      string `json:"message.timestamp.difference.max.ms"`
		MessageTimestampType                 string `json:"message.timestamp.type"`
		MinCleanableDirtyRatio               string `json:"min.cleanable.dirty.ratio"`
		MinCompactionLagMs                   string `json:"min.compaction.lag.ms"`
		MinInsyncReplicas                    string `json:"min.insync.replicas"`
		NumPartitions                        string `json:"num.partitions"`
		Preallocate                          string `json:"preallocate"`
		RetentionBytes                       string `json:"retention.bytes"`
		RetentionMs                          string `json:"retention.ms"`
		SegmentBytes                         string `json:"segment.bytes"`
		SegmentIndexBytes                    string `json:"segment.index.bytes"`
		SegmentJitterMs                      string `json:"segment.jitter.ms"`
		SegmentMs                            string `json:"segment.ms"`
		UncleanLeaderElectionEnable          string `json:"unclean.leader.election.enable"`
	} `json:"config"`
}
