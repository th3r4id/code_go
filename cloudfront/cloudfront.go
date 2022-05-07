package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	Client := cloudfront.NewFromConfig(cfg)
	//	Client2 := s3.NewFromConfig(cfg)
	output, err := Client.ListDistributions(context.TODO(), &cloudfront.ListDistributionsInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range output.DistributionList.Items {
		cloud_front := *item.Origins
		for _, rand := range cloud_front.Items {
			//fmt.Println(*rand.DomainName)
			//	fmt.Println(*rand.Id)
			paramStr := strings.Split(*rand.DomainName, ".")
			for _, param := range paramStr {
				//	fmt.Println(param)
				if param == "s3" {
					//	fmt.Println(*rand.DomainName)

					_, err := Client.GetCloudFrontOriginAccessIdentity(context.TODO(), &cloudfront.GetCloudFrontOriginAccessIdentityInput{})
					if err == nil {
						fmt.Println("Error")
					}
					fmt.Println("No Policy Found Restriced access is not configured for", *rand.DomainName)
					// for _, items := range *output2.CloudFrontOriginAccessIdentity.Id {
					// 	s3_oai := items
					// 	fmt.Println(s3_oai)
					// }

				}

				// if output2.Origins.Origi
				// cloud_out := output2.CloudFrontOriginAccessIdentity.Id
				// fmt.Println(*cloud_out)
			}
		}

	}
}
