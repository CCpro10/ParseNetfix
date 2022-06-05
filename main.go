package main

import (
	"fmt"
	"log"
	"main/service"
	"os"
)

func main() {
	//设置网络代理
	os.Setenv("HTTP_PROXY", "127.0.0.1:19180")
	os.Setenv("HTTPS_PROXY", "127.0.0.1:19180")
	service.InitCmsMap()
	//由于网络原因,Fetch很可能是不成功的
	fullUrl := "https://www.netflix.com/title/80168230"
	body, err := service.Fetch(fullUrl)

	CheckErrLogIfNotNil(err)

	cms, err := service.ParseDetail(service.CmsMap, body)
	CheckErrLogIfNotNil(err)
	fmt.Println(cms)

}

func CheckErrLogIfNotNil(err error) {
	if err != nil {
		log.Println(err)
	}
}
