package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func httpClient() *http.Client {
    client := &http.Client{Timeout: 10 * time.Second}
    return client
}

func send_request(client *http.Client) {

    resp, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts/1", nil)
    if err != nil {
        log.Fatalln(err)
    }

    response, err := client.Do(resp)
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalln(err)
    }
    defer response.Body.Close()
    sb := string(body)
    log.Printf(sb)
}

func main() {
    client := httpClient()
    send_request(client)
}
