package usecase

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mediaUsecase struct {
	conn           *utils.Conn
	contextTimeout time.Duration
}

// NewMediaUsecase creates a new feedback usecase
func NewMediaUsecase(conn *utils.Conn, timeout time.Duration) domain.MediaUsecase {
	return &mediaUsecase{
		conn:           conn,
		contextTimeout: timeout,
	}
}

func (u *mediaUsecase) Store(c context.Context, file *multipart.FileHeader, buf bytes.Buffer) (err error) {
	_, err = s3.New(u.conn.AWS).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String("portal-jabar-staging"),
		Key:                  aws.String("staging/media/img/" + file.Filename),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buf.Bytes()),
		ContentLength:        aws.Int64(file.Size),
		ContentType:          aws.String(http.DetectContentType(buf.Bytes())),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	
	return
}
