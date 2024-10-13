package main

import (
	"article/pkg/config"
	"fmt"
)

func main() {
	conf, err := config.NewGlobalConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conf)
}
