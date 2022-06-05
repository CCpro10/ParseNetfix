package main

import (
	"fmt"
	"log"
	"main/service"
	"os"
)

const LocalProxyPort = "127.0.0.1:19180"

func main() {
	SetProxyEnv(LocalProxyPort)
	service.InitCmsMap()
	//由于地区原因,Fetch很可能无法获取数据

	fullUrl := "https://www.netflix.com/title/80168230"
	body, err := service.Fetch(fullUrl)
	CheckErrLogIfNotNil(err)
	cms, err := service.ParseDetail(service.CmsMap, body)
	CheckErrLogIfNotNil(err)
	fmt.Println(cms)

	fullUrl = "https://www.netflix.com/title/80221584"
	body, err = service.Fetch(fullUrl)
	CheckErrLogIfNotNil(err)
	cms, err = service.ParseDetail(service.CmsMap, body)
	CheckErrLogIfNotNil(err)
	fmt.Println(cms)

	for id, v := range service.CmsMap {
		fmt.Println(id, ":", *v)
	}
}

func SetProxyEnv(localProxyPort string) {
	//设置网络代理
	os.Setenv("HTTP_PROXY", localProxyPort)
	os.Setenv("HTTPS_PROXY", localProxyPort)
}

func CheckErrLogIfNotNil(err error) {
	if err != nil {
		log.Println(err)
	}
}
