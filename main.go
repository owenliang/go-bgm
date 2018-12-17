package main

import (
	"github.com/owenliang/go-bgm/backend/framework"
	"fmt"
)

func main() {
	var (
		err error
	)

	if err = framework.Run(); err != nil {
		fmt.Println(err)
	}
}
