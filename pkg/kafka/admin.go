package kafka_operator

import (
	"log"

	"github.com/Shopify/sarama"
)

type (
	KafkaClusterAdminAction struct {
		admin sarama.ClusterAdmin
	}
	TopicDescribe struct {
		Name        string `json:"name"`
		Partition   int32  `json:"partition"`
		Replication int16  `json:"replication"`
	}
)

var DefaultKafkaClusterAdminAction = &KafkaClusterAdminAction{}

func newKafkaClusterAdmin(broker []string) *KafkaClusterAdminAction {
	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0
	admin, e := sarama.NewClusterAdmin(broker, config)
	if e != nil {
		log.Println("KafkaAdminAction: NewKafka: ", e.Error())
		return nil
	}
	return &KafkaClusterAdminAction{admin: admin}
}

func (A KafkaClusterAdminAction) GetTopic() []TopicDescribe {
	topics, e := A.admin.ListTopics()
	if e != nil {
		log.Println("KafkaAdminAction: GetTopic: ", e.Error())
		return nil
	}
	var result []TopicDescribe
	for k, v := range topics {
		result = append(result, TopicDescribe{
			Name:        k,
			Partition:   v.NumPartitions,
			Replication: v.ReplicationFactor,
		})
	}
	return result
}

func (A KafkaClusterAdminAction) CreateTopic(name string, partition int32, factor int16) TopicDescribe {
	detail := &sarama.TopicDetail{
		NumPartitions:     partition,
		ReplicationFactor: factor,
	}
	var result TopicDescribe

	e := A.admin.CreateTopic(name, detail, true)
	if e != nil {
		log.Println("KafkaAdminAction: CreatTopic: ", e.Error())
		return result
	} else {
		result = TopicDescribe{
			Name:        Topic,
			Partition:   partition,
			Replication: factor,
		}
		return result
	}
}
