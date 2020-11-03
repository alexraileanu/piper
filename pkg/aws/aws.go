package aws

import (
    "bytes"
    "log"
    "os"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type AWS struct {
    sess   *session.Session
    client *s3.S3
}

func Initialize() *AWS {
    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String("eu-central-1"),
        Credentials: credentials.NewStaticCredentials(
            os.Getenv("AWS_SECRET_KEY_ID"),
            os.Getenv("AWS_SECRET_ACCESS_KEY"),
            "",
        ),
    }))
    client := s3.New(sess)

    return &AWS{
        sess:   sess,
        client: client,
    }
}

func (a *AWS) Post(filePath string, bucket string, key string) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("error opening file: %v\n", err)
        return
    }

    fileInfo, _ := file.Stat()
    size := fileInfo.Size()
    buffer := make([]byte, size)

    _, err = file.Read(buffer)
    if err != nil {
        log.Fatalf("error reading file contents: %v\n", err)
        return
    }

    _, err = a.client.PutObject(&s3.PutObjectInput{
        Bucket:        aws.String(bucket),
        Body:          bytes.NewReader(buffer),
        ContentLength: aws.Int64(size),
        Key:           aws.String(key),
    })
    if err != nil {
        log.Fatalf("error posting to s3: %v\n", err)
        return
    }
}
