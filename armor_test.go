package armor_test

import (
	"testing"

	"github.com/qba73/armor"
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
	if !armor.Validate(policy) {
		t.Error("a valid policy flagged as invalid")
	}
}

func TestValidatePolicy_CanDetectAnInvalidPolicy(t *testing.T) {
	policy := `{"invalid": "policy"}`
	if armor.Validate(policy) {
		t.Error("an invalid policy flagged as valid")
	}
}
