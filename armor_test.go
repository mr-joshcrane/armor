package armor_test

import (
	"testing"

	"github.com/mr-joshcrane/armor"
)

func TestValidatePolicy_CanDetectAValidPolicy(t *testing.T) {
	policy := `
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "s3:GetObject",
      "Resource": "arn:aws:s3:::examplebucket/*"
    }
  ]
}`
	err := armor.Validate(policy)
	if err != nil {
		t.Error("expected no error, got:", err)
	}
}

func TestValidatePolicy_CanDetectAnInvalidPolicy(t *testing.T) {
	policy := `{"invalid": "policy"}`
	err := armor.Validate(policy)
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestValidatePolicy_CanDetectAnInvalidPolicyWithAnInvalidEffect(t *testing.T) {
	policy := `
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "MaybeMaybeNot",
			"Action": "s3:InvalidAction",
			"Resource": "arn:aws:s3:::examplebucket/*"
		}
	]
}`
	err := armor.Validate(policy)
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestValidatePolcy_CanDetectPolicyMissingRequiredFields(t *testing.T) {
	policy := `
{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Resource": "arn:aws:s3:::examplebucket/*"
		}
	]
}`

	err := armor.Validate(policy)
	if err == nil {
		t.Error("expected an error, got nil")
	}
}
