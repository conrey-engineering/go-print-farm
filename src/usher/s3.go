package main

import (
	// "errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/conrey-engineering/go-print-farm/src/protobufs/print"
	"os"
)

type S3File struct {
	Bucket   string
	Name     string
	Shasum   string
	FilePath string
}

// Load from PrintFile object because its a protobuf and helpers aren't a thing in this project yet
func (s *S3File) LoadFromPrintFile(p print.PrintFile) {
	s.Bucket = p.BucketName
	s.Name = p.Filename
	s.Shasum = p.Shasum
}

// Pulls file down from s3 into a temporary location
// this is better than loading an STL into memory for future uploading
func (s *S3File) Download(client *s3manager.Downloader) (*os.File, error) {
	file, err := os.CreateTemp("/tmp", "s3-")
	if err != nil {
		return nil, err
	}

	s.FilePath = file.Name()

	// fmt.Println("test")
	// fmt.Println(file)

	_, err = client.Download(
		file,
		&s3.GetObjectInput{
			Bucket: &s.Bucket,
			Key:    &s.Name,
		},
	)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
			case s3.ErrCodeInvalidObjectState:
				fmt.Println(s3.ErrCodeInvalidObjectState, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		}
		return nil, err
	}

	// if result.ChecksumSHA256 != nil {
	// 	if *result.ChecksumSHA256 != s.Shasum {
	// 		errors.New("Checksum mismatch")
	// 	}
	// }

	return file, nil
}
