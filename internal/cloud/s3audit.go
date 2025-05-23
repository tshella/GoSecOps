package cloud

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3AuditResult struct {
	Bucket         string `json:"bucket"`
	PublicAccess   bool   `json:"public_access"`    // true = publicly accessible
	ACLGrantsExist bool   `json:"acl_grants_exist"` // true = public ACL found
	Error          string `json:"error,omitempty"`
}

func AuditS3Bucket(profile, bucket string) S3AuditResult {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithSharedConfigProfile(profile))
	if err != nil {
		return S3AuditResult{Bucket: bucket, Error: "Failed to load AWS config: " + err.Error()}
	}

	client := s3.NewFromConfig(cfg)

	// 1. Check ACLs
	aclResp, err := client.GetBucketAcl(ctx, &s3.GetBucketAclInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return S3AuditResult{Bucket: bucket, Error: "Unable to get ACL: " + err.Error()}
	}

	publicACL := false
	for _, grant := range aclResp.Grants {
		if grant.Grantee != nil && grant.Grantee.URI != nil {
			if *grant.Grantee.URI == "http://acs.amazonaws.com/groups/global/AllUsers" {
				publicACL = true
				break
			}
		}
	}

	// 2. Check Public Access Block config
	publicAccess := true // assume public unless proven otherwise
	blockResp, err := client.GetPublicAccessBlock(ctx, &s3.GetPublicAccessBlockInput{
		Bucket: aws.String(bucket),
	})
	if err == nil && blockResp.PublicAccessBlockConfiguration != nil {
		block := blockResp.PublicAccessBlockConfiguration

		// Dereference *bools safely
		allBlocked := aws.ToBool(block.BlockPublicAcls) &&
			aws.ToBool(block.BlockPublicPolicy) &&
			aws.ToBool(block.IgnorePublicAcls) &&
			aws.ToBool(block.RestrictPublicBuckets)

		publicAccess = !allBlocked // true means public
	}

	return S3AuditResult{
		Bucket:         bucket,
		PublicAccess:   publicAccess,
		ACLGrantsExist: publicACL,
	}
}
