package main

import (
	"fmt"
	"log"
	"main/util"
)

func main() {
	util.InitCmsMap()
	//由于网络原因,Fetch很可能是不成功的
	fullUrl := "https://www.netflix.com/title/80168230"
	body, err := util.Fetch(fullUrl)

	CheckErrLogIfNotNil(err)

	cms, err := util.ParseDetail(util.CmsMap, body)
	CheckErrLogIfNotNil(err)
	fmt.Println(cms)

}

func CheckErrLogIfNotNil(err error) {
	if err != nil {
		log.Println(err)
	}
}
