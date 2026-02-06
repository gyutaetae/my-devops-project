package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1> Go 배포 성공!</h1>")
		fmt.Fprintf(w, "<p>이 메시지는 GitHub Actions를 통해 S3로 전달되었습니다.</p>")
	})

	// 포트 설정 (로컬 테스트용)
	port := "8080"
	fmt.Printf("서버가 %s 포트에서 시작되었습니다...\n", port)
	
	// 배포를 위해 빌드 결과만 만들 거라 실제 실행은 S3에서 직접 안 되지만, 
	// 빌드 테스트를 위해 코드를 완성해두자!
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		os.Exit(1)
	}
}