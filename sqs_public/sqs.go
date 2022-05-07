package main

// SQS Publically Accessible
import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	Client := sqs.NewFromConfig(cfg)

	output, err := Client.ListQueues(context.TODO(), &sqs.ListQueuesInput{})
	if err != nil {
		log.Fatal(err)
	}

	sqs_queues := output.QueueUrls
	for i := range sqs_queues {
		fmt.Println(sqs_queues[i])
		input, err := Client.GetQueueAttributes(context.TODO(), &sqs.GetQueueAttributesInput{QueueUrl: &sqs_queues[i], AttributeNames: []types.QueueAttributeName{"Policy"}})
		if err != nil {
			log.Fatal(err)
		}
		policy_name := input.Attributes["Policy"]

		var tmp map[string]interface{}

		err = json.Unmarshal([]byte(policy_name), &tmp)
		//fmt.Println(tmp["Statement"].([]interface{})[0].(map[string]interface{})["Principal"].(map[string]interface{})["AWS"])

		if tmp["Statement"].([]interface{})[0].(map[string]interface{})["Principal"].(map[string]interface{})["AWS"] == "*" {
			fmt.Println("SQS is Publicly Accessible")
		} else {
			fmt.Println("Don't Worry Everything is good")
		}
		// slice_name := policy_name[108:131]
		// if slice_name == "\"Principal\":{\"AWS\":\"*\"}" {
		// 	fmt.Println("SQS is Publicly Accessible")
		// } else {
		// 	fmt.Println("Don't Worry Everything is good")
		// }
		//	fmt.Println(slice_name)
		//	fmt.Println(input.Attributes["Policy"])
		// log.Printf("Queue Atrribute=%v", input.ResultMetadata)
	}
}
