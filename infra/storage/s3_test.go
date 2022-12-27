package infra

import (
	"testing"
)

func TestNewS3Uploader(t *testing.T) {
	if got := NewS3Uploader(); got == nil {
		t.Errorf("NewS3Uploader: got nil; want *S3Uploader")
	}
}
