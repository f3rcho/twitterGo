package routers

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/f3rcho/twitterGo/db"
	"github.com/f3rcho/twitterGo/models"
)

type readSeeker struct {
	io.Reader
}

func (rs readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.ResposenAPI {
	var r models.ResposenAPI
	r.Status = 400
	IDUser := claim.ID.Hex()

	var filename string
	var user models.User
	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))
	switch uploadType {
	case "A":
		filename = "avatars/" + IDUser + ".jpg"
		user.Avatar = filename
	case "B":
		filename = "banners/" + IDUser + ".jpg"
		user.Banner = filename
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["Content-Type"])
	if err != nil {
		r.Status = 500
		r.Message = "Error al parsear el tipo de archivo"
		return r
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)
		if err != nil {
			r.Status = 500
			r.Message = "Error decoding the archive" + err.Error()
			return r
		}

		mp := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		p, err := mp.NextPart()
		if err != nil && err != io.EOF {
			r.Status = 500
			r.Message = "Error getting the NextPart" + err.Error()
			return r
		}
		if err != io.EOF {
			if p.FileName() != "" {
				buffer := bytes.NewBuffer(nil)
				if _, err := io.Copy(buffer, p); err != nil {
					r.Status = 500
					r.Message = "Error copying the archive" + err.Error()
					return r
				}

				s, err := session.NewSession(&aws.Config{
					Region: aws.String("us-east-1"),
				})
				if err != nil {
					r.Status = 500
					r.Message = "Error creating the session" + err.Error()
					return r
				}
				uploader := s3manager.NewUploader(s)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buffer},
				})
				if err != nil {
					r.Status = 500
					r.Message = "Error uploading file" + err.Error()
					return r
				}
			}
		}
		status, err := db.UpdateUser(user, IDUser)
		if err != nil || !status {
			r.Status = 400
			r.Message = "Error updating user"
			return r
		}

	} else {
		r.Message = "Must send an image with 'Content-Type' type multipart"
		r.Status = 400
		return r
	}
	r.Status = 200
	r.Message = "File uploaded successfuly"
	return r
}
