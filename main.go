package main

import (
	"context"
	"fmt"
	"net/http" // 👈 웹 서버를 만들기 위한 도구
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	// 1. S3 설정 (기존과 동일)
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: "http://localhost:4566", SigningRegion: "us-east-1"}, nil
	})
	cfg, _ := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	client := s3.NewFromConfig(cfg, func(o *s3.Options) { o.UsePathStyle = true })

	// 2. 웹 서버 핸들러 설정
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		bucketName := "my-new-test-bucket"
		fileName := "web-upload.txt"
		content := "이 파일은 웹 브라우저 요청으로 생성되었습니다!"

		// S3에 업로드
		_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
			Body:   strings.NewReader(content),
		})

		if err != nil {
			http.Error(w, "S3 업로드 실패: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "<h1>✅ 업로드 성공!</h1><p>S3에 '%s' 파일이 저장되었습니다.</p>", fileName)
	})

	// 3. 서버 시작
	fmt.Println("🌐 Go 웹 서버가 8080 포트에서 시작되었습니다!")
	fmt.Println("👉 http://localhost:8080/upload 에 접속해보세요.")
	logFatal := http.ListenAndServe(":8080", nil)
	if logFatal != nil {
		fmt.Printf("서버 중단: %v\n", logFatal)
	}
}