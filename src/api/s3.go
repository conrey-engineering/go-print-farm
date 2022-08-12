package main

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/conrey-engineering/go-print-farm/src/protobufs/files"
	"os"
)

func uploadPrintFile(client *s3manager.Uploader, s files.S3File, body []byte) error {
	f, err := os.CreateTemp("/tmp", "s3-")
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	_, err = client.Upload(&s3manager.UploadInput{
		Bucket: &s.Bucket,
		Key:    &s.Name,
		Body:   f,
	})

	return err
}
