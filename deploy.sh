#!/bin/bash
echo "🚀 배포 자동화 시작!"
terraform apply -auto-approve
go build -o server main.go
echo "✅ 배포 완료! 서버를 실행합니다."
./server