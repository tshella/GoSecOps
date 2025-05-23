package cloud

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

type IAMFinding struct {
	User       string   `json:"user"`
	PolicyName string   `json:"policy_name"`
	Issues     []string `json:"issues"`
}

func AnalyzeIAMPolicies(profile string) ([]IAMFinding, error) {
	// Load AWS config using provided profile
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %v", err)
	}

	client := iam.NewFromConfig(cfg)
	var findings []IAMFinding

	// List IAM users
	usersResp, err := client.ListUsers(context.TODO(), &iam.ListUsersInput{})
	if err != nil {
		return nil, fmt.Errorf("error listing IAM users: %v", err)
	}

	for _, user := range usersResp.Users {
		policiesResp, err := client.ListAttachedUserPolicies(context.TODO(), &iam.ListAttachedUserPoliciesInput{
			UserName: user.UserName,
		})
		if err != nil {
			continue
		}

		for _, policy := range policiesResp.AttachedPolicies {
			policyResp, err := client.GetPolicy(context.TODO(), &iam.GetPolicyInput{
				PolicyArn: policy.PolicyArn,
			})
			if err != nil || policyResp.Policy == nil {
				continue
			}

			// We will just flag * (wildcard) here as an example
			issues := []string{}
			if strings.Contains(strings.ToLower(*policy.PolicyName), "admin") {
				issues = append(issues, "Contains 'admin' in policy name")
			}
			// Real parsing would involve downloading the policy version and checking JSON actions

			if len(issues) > 0 {
				findings = append(findings, IAMFinding{
					User:       aws.ToString(user.UserName),
					PolicyName: aws.ToString(policy.PolicyName),
					Issues:     issues,
				})
			}
		}
	}

	return findings, nil
}
