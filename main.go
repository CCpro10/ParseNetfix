package main

import (
	"fmt"
	"log"
	"main/util"
)

func main() {
	util.InitCmsMap()
	fullUrl := "https://www.netflix.com/title/80168230"
	body, err := util.Fetch(fullUrl)
	log.Println(body, err)
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
