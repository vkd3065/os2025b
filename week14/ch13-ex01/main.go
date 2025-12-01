package main

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(2000 * time.Millisecond)
	}
}

func hi() {
	fmt.Println("안녕")
	time.Sleep(10 * time.Second)
}

func main() {
	start := time.Now()
	hi()         // 새 고루틴에서 실행
	go say("메인") // 메인 고루틴에서 실행
	fmt.Println("전체 실행 시간 : ", time.Since(start))
}
