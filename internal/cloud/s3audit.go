package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3AuditResult struct {
	Bucket         string `json:"bucket"`
	PublicAccess   bool   `json:"public_access"`
	ACLGrantsExist bool   `json:"acl_grants_exist"`
	Error          string `json:"error,omitempty"`
}

func AuditS3Bucket(profile, bucket string) S3AuditResult {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile))
	if err != nil {
		return S3AuditResult{Bucket: bucket, Error: "Failed to load AWS config: " + err.Error()}
	}

	client := s3.NewFromConfig(cfg)

	// 1. Check ACL
	aclResp, err := client.GetBucketAcl(context.TODO(), &s3.GetBucketAclInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return S3AuditResult{Bucket: bucket, Error: "Unable to get ACL: " + err.Error()}
	}

	public := false
	for _, grant := range aclResp.Grants {
		if grant.Grantee != nil && grant.Grantee.URI != nil {
			if *grant.Grantee.URI == "http://acs.amazonaws.com/groups/global/AllUsers" {
				public = true
			}
		}
	}

	// 2. Check Public Access Block config
	blockResp, err := client.GetBucketPublicAccessBlock(context.TODO(), &s3.GetBucketPublicAccessBlockInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		// Not having a public access block doesn't necessarily mean it's public
	}

	publicAccess := true
	if blockResp.PublicAccessBlockConfiguration != nil {
		cfg := blockResp.PublicAccessBlockConfiguration
		publicAccess = !(cfg.BlockPublicAcls && cfg.BlockPublicPolicy && cfg.IgnorePublicAcls && cfg.RestrictPublicBuckets)
	}

	return S3AuditResult{
		Bucket:         bucket,
		PublicAccess:   publicAccess,
		ACLGrantsExist: public,
	}
}
