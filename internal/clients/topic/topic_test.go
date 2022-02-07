package topic

import (
	"testing"

	"github.com/dfds/provider-confluent/apis/topic/v1alpha1"
	"github.com/dfds/provider-confluent/internal/clients"
	"github.com/stretchr/testify/assert"
	"go.dfds.cloud/utils/config"
)

var (
	testConfig = Config{
		APICredentials: clients.APICredentials{},
	}
	cluster           = config.GetEnvValue("CONFLUENT_CLUSTER_ID", "")
	environment       = config.GetEnvValue("CONFLUENT_ENVIRONMENT", "")
	name              = "crossplane-test-topic"
	partitions        = 1
	retention   int64 = 259200000
	topicConfig       = v1alpha1.TopicConfig{Name: name, Partitions: partitions, Config: v1alpha1.Config{Retention: retention}}
	topic             = v1alpha1.TopicParameters{Cluster: cluster, Environment: environment, Topic: topicConfig}
	client            = NewClient(testConfig)
)

func TestTopicLifecycle(t *testing.T) {
	clients.SkipCI(t)
	assert := assert.New(t)

	_, err := client.TopicDescribe(v1alpha1.TopicObservation{Cluster: cluster, Environment: environment, Name: name})
	if err != nil {
		assert.Equal(err.Error(), ErrUnknownTopic)
	} else {
		t.Errorf("expected topic unknow error but got: %s", err.Error())
	}

	_, err = client.TopicDescribe(v1alpha1.TopicObservation{Cluster: cluster, Environment: environment, Name: ""})
	if err != nil {
		assert.Equal(err.Error(), ErrInvalidInput)
	} else {
		t.Errorf("expected invalid input error but got: %s", err.Error())
	}

	err = client.TopicCreate(topic)
	if err != nil {
		t.Errorf("topic creation failed with unknow error: %s", err.Error())
	}

	_, err = client.TopicDescribe(v1alpha1.TopicObservation{Cluster: cluster, Environment: environment, Name: name})
	if err != nil {
		t.Errorf("cannot find previously created topic: %s", err.Error())
	}

	topic.Topic.Config.Retention = 604800000
	err = client.TopicUpdate(topic)
	if err != nil {
		t.Errorf("cannot update topic retion: %s", err.Error())
	}

	err = client.TopicDelete(topic)
	if err != nil {
		t.Errorf("cannot delete topic (manual clean up may be required): %s", err.Error())
	}

}
