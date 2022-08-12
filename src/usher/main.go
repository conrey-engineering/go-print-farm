package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/conrey-engineering/go-print-farm/src/protobufs/print"
	"github.com/google/uuid"
	"os"
	"os/signal"
	"syscall"
)

// watches for PrintFile objects on a channel, creates temporary files with contents of PrintFile
func downloadQueueHandler(prints <-chan print.PrintRequest, s3dl *s3manager.Downloader) {
	for {
		print := <-prints
		var s3file = S3File{}

		s3file.LoadFromPrintFile(*print.File)
		fp, err := s3file.Download(s3dl)
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Println(s3file)
		fmt.Println(fp)
	}
}

func main() {
	var (
		printDownloadChan = make(chan print.PrintRequest)
		// printUploadChan   = make(chan print.Print)
	)

	awsSess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:           aws.String("us-east-2"),
			Endpoint:         aws.String("http://localhost:9000"),
			Credentials:      credentials.NewStaticCredentials("testsvcaccount", "testsvcaccount", ""),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		},
	}))

	s3dl := s3manager.NewDownloader(awsSess)

	go downloadQueueHandler(printDownloadChan, s3dl)

	printFile := print.PrintFile{
		Id:         uuid.New().String(),
		Filename:   "my_first_print.stl",
		Shasum:     "xxx",
		BucketName: "print-farm",
	}
	print := print.PrintRequest{
		Id:   uuid.New().String(),
		Name: "my first print",
		File: &printFile,
	}

	printDownloadChan <- print

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		_ = <-sigs
		done <- true
	}()

	<-done
}
