package service

import (
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	body, err := Fetch("https://www.netflix.com/sg/title/80168230")
	if err != nil {
		panic(err)
	}
	log.Printf("%s", body)
}
