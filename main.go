package main

import (
	"com.higginslee/go_proxy_pool/service"
	"com.higginslee/go_proxy_pool/utils"
)

func main() {
	// init config
	err := utils.Init()
	if err != nil {
		panic(err)
	}
	// run crawel main
	service.Crawel()
}
