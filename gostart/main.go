package main

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	log.Println("Go Start!")

	client := resty.New()

	res, err := client.R().Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res.String())

}
