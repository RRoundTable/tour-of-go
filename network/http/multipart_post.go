package main

import (
	"io"
    "bytes"
	"log"
	"net/http"
    "mime/multipart"
	"os"
	"strings"
	"time"
)

func httpClient() *http.Client {
    client := &http.Client{Timeout: 10 * time.Second}
    return client
}


func createFrom(form map[string]string) (string, io.Reader, error) {
    body := new(bytes.Buffer)
    mp := multipart.NewWriter(body)
    defer mp.Close()
    for key, val := range form {
        if strings.HasPrefix(val, "@") {
            val = val[1:]
            file, err := os.Open(val)
            if err != nil { return "", nil, err }
            defer file.Close()

            part, err := mp.CreateFormFile(key, val)
            if err != nil { return "", nil, err}
            io.Copy(part, file)
        }   else {
            mp.WriteField(key, val)
        }
    }
    return mp.FormDataContentType(), body, nil
}



func sendPost(client *http.Client, url, path string) {
    form := map[string]string{"file": "@sample.jpg", "key": "file"}
    ct, body, err := createFrom(form)
    if err != nil {
        log.Fatal(err)
    }
    req, err := http.NewRequest(http.MethodPost, url, body)
    req.Header.Add("Content-Type", ct)
    resp, err := client.Do(req)
    log.Print(resp)
}

func main() {
    client := httpClient()
    sendPost(
        client,
        "http://localhost:10000/find-bbox",
        "/Users/wontakryu/annotation-ai/tour-of-go/network/http/sample.jpg",
    )
}
