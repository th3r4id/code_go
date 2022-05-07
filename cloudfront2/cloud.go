package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	Client := cloudfront.NewFromConfig(cfg)
	Client_load := elasticloadbalancingv2.NewFromConfig(cfg)
	output, err := Client.ListDistributions(context.TODO(), &cloudfront.ListDistributionsInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range *&output.DistributionList.Items {
		load_cloud := item.Origins.Items
		for _, load_items := range load_cloud {
			cloud_dns := *load_items.DomainName
			//	fmt.Println(cloud_dns)

			paramStr := strings.Split(cloud_dns, ".")
			//	fmt.Println(paramStr)
			for _, param := range paramStr {
				//	fmt.Println(param)
				if param == "elb" {
					//	fmt.Println(cloud_dns)
					output2, err := Client_load.DescribeLoadBalancers(context.TODO(), &elasticloadbalancingv2.DescribeLoadBalancersInput{})
					if err != nil {
						log.Fatal(err)
					}
					for _, item := range output2.LoadBalancers {
						load_b := *item.DNSName
						//	fmt.Println(load_b)
						if cloud_dns == load_b {
							fmt.Printf("%s We got your back\n", cloud_dns)
						} else {
							fmt.Printf("%s Takeover Possible\n", cloud_dns)
						}

					}
				}
			}
		}
	}
}
