provider "aws" {
  region = "ap-northeast-2"
}

resource "aws_s3_bucket" "my_bucket" {
  bucket        = "my-unique-bucket-name-9988"
  force_destroy = true  # 이 줄을 추가해!
}

resource "aws_ecr_repository" "my_app" {
  name = "my-go-app-repo"
  force_delete = true
}

# 나중에 주소를 확인하기 위해 출력 설정
output "ecr_repository_url" {
  value = aws_ecr_repository.my_app.repository_url
}