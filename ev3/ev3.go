package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	Client := elasticsearchservice.NewFromConfig(cfg)

	output, err := Client.ListDomainNames(context.TODO(), &elasticsearchservice.ListDomainNamesInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range output.DomainNames {
		//	fmt.Println(*item.DomainName)
		domainName := *item.DomainName
		//	fmt.Println(domain_name)
		input, err := Client.DescribeElasticsearchDomainConfig(context.TODO(), &elasticsearchservice.DescribeElasticsearchDomainConfigInput{DomainName: aws.String(domainName)})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(*input.DomainConfig.AccessPolicies.Options)
		//var tmp map[string]interface{}
		//err = json.Unmarshal([]byte(*input.DomainConfig.AccessPolicies.Options), &tmp)
		//if err != nil {
		//	log.Fatal(err)
		//}

		//		fmt.Println(tmp["Statement"].([]interface{})[0].(map[string]interface{})["Principal"].(map[string]interface{})["AWS"])
		//if tmp["Statement"].([]interface{})[0].(map[string]interface{})["Principal"].(map[string]interface{})["AWS"] == "*" {
		//	fmt.Printf("%s Domain Has Global Access! Not Protected", domain_name)
		//} else {
		//	fmt.Println("You are Protected")
		//}
	}

}
