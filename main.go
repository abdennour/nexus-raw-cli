package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello here")
	version := flag.String("version", "", "artifact version e.g: 1.0.0")
	flag.Parse()
	fmt.Printf("version is %s", *version)
	uploadfile(*version)
}

func uploadfile(version string) {
	data, err := os.Open("nexus-cli")
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://nexus:8081/repository/drive/cli/%s/nexus-cli", version), data)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "admin123")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
}




