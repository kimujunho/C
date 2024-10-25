package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1
	fmt.Println("정답(디버그용):", answer)

	fmt.Print("숫자 입력: ")
	in := bufio.NewReader(os.Stdin)
	i, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)
	guess, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("입력한 숫자:", guess)

	if answer == guess {
		fmt.Println("정답입니다!")
	} else if answer < guess {
		fmt.Println("입력하신 수는 정답보다 작은 수 입니다. LOW")
	} else {
		fmt.Println("입력하신 수는 정답보다 큰 수 입니다. HIGH")
	}

}
