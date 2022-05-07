package main

// unused-roles Completed

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal(err)
	}
	client := iam.NewFromConfig(cfg)

	output, err := client.ListRoles(context.TODO(), &iam.ListRolesInput{})
	//	output2, err2 := client.GetRole(context.TODO(), &iam.GetRoleInput{})
	if err != nil {
		log.Fatal(err)
	}

	roles_list := output.Roles
	for i := range roles_list {
		input, err := client.GetRole(context.TODO(), &iam.GetRoleInput{RoleName: aws.String(*roles_list[i].RoleName)})
		if err != nil {
			log.Fatal(err)
		}
		//	log.Printf("Role Name=%s\tRoleLastUsed=%v", *input.Role.RoleName, input.Role.RoleLastUsed)

		//	fmt.Printf("%v\n", input.Role.RoleLastUsed.LastUsedDate.Unix())
		if input.Role.RoleLastUsed.LastUsedDate == nil {
			fmt.Printf("Unused Roles Found RoleName=%s\n", *input.Role.RoleName)
		} else {
			if input.Role.RoleLastUsed.LastUsedDate.Unix() <= 1646838002 {
				fmt.Printf("%s Role has not been used for more 7 days\n", *input.Role.RoleName)
			}
			// if input.Role.RoleLastUsed == nil {
			// 		fmt.Printf("%s role is not used please delete the userrole")
			// 	}
		}

		/**	for _, object := range output.Roles {
			log.Printf("RoleName=%s", aws.ToString(object.RoleName))
		}
		fmt.Println(output.Roles)  **/

	}
}
