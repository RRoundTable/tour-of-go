package main

import (
    "io/ioutil"
    "log"
    "net/http"
)

func main() {

    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")

    if err != nil {
        log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    sb := string(body)
    log.Printf(sb)
}
