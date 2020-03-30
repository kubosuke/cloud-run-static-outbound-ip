package main

import (
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
	"os"
)


func main() {
	http.HandleFunc("/ip/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	client := StaticIPClient()

	resp, err := client.Get("https://ifconfig.me/ip")
	defer resp.Body.Close()
	if err != nil {
		log.Fatal("cannot read res body", err)
	}
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(byteArray))
}

func StaticIPClient() *http.Client {
	p, _ := proxy.SOCKS5("tcp", os.Getenv("HTTP_PROXY"), nil, proxy.Direct)
	client := http.DefaultClient
	client.Transport = &http.Transport{
		Dial: p.Dial,
	}
	return client
}
