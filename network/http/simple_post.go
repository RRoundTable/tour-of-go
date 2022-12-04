package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "log"
)


func main() {
    postBody, _ := json.Marshal(map[string]string{
        "name": "roundtable",
        "email": "ryu071511@gmail.com",
    })

    responseBody := bytes.NewBuffer(postBody)

    resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)

    if err != nil {
        log.Fatalf("An Error Occured %v", err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    sb := string(body)
    log.Printf(sb)
}
