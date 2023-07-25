package usecase

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/config"
	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/utils"

	"github.com/jabardigitalservice/portal-jabar-services/core-service/src/domain"
)

type mediaUsecase struct {
	config         *config.Config
	conn           *utils.Conn
	contextTimeout time.Duration
}

// NewMediaUsecase creates a new feedback usecase
func NewMediaUsecase(c *config.Config, conn *utils.Conn, timeout time.Duration) domain.MediaUsecase {
	return &mediaUsecase{
		config:         c,
		conn:           conn,
		contextTimeout: timeout,
	}
}

func (u *mediaUsecase) newMediaResponse(fileName string, fileDownloadUri string, size int64) *domain.MediaResponse {
	return &domain.MediaResponse{
		FileName:        fileName,
		FileDownloadUri: fileDownloadUri,
		Size:            size,
	}
}

func (u *mediaUsecase) Store(_ context.Context, file *multipart.FileHeader, buf bytes.Buffer, domain string) (res *domain.MediaResponse, err error) {
	fileName := strings.Replace(fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename), " ", "-", -1) // fixme
	fileSize := file.Size
	filePath := u.config.AWS.Env + "/media/img/" + domain + fileName
	fileDownloadUri := u.config.AWS.Cloudfront + "/" + filePath

	_, err = s3.New(u.conn.AWS).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(u.config.AWS.Bucket),
		Key:                  aws.String(filePath),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buf.Bytes()),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(buf.Bytes())),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	return u.newMediaResponse(fileName, fileDownloadUri, fileSize), err
}

func (u *mediaUsecase) Delete(_ context.Context, body *domain.DeleteMediaRequest) (err error) {
	filePath := u.config.AWS.Env + "/media/img/" + body.Domain + body.Key

	_, err = s3.New(u.conn.AWS).DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(u.config.AWS.Bucket),
		Key:    aws.String(filePath),
	})

	if err != nil {
		return
	}

	return
}
