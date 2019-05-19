package helpers

import (
	"bytes"
	_ "fmt"
	"io/ioutil"
	"net/http"
)

var defaultResponse = map[int]string {
    200: "Success",
    400: "Bad Request",
    403: "Forbidden",
    404: "Not Found",
    405: "Method Not Allowed",
    500: "Internal Server Error",
}

func JsonPost(url string, jsonPost []byte) (int, string) {
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPost))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return resp.StatusCode, ""
    }

    body, _ := ioutil.ReadAll(resp.Body)
    return resp.StatusCode, string(body)
}

func SendGetRequest(url string) (int, string) {
    req, err := http.NewRequest("GET", url, nil)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        return resp.StatusCode, ""
    }

    body, _ := ioutil.ReadAll(resp.Body)
    return resp.StatusCode, string(body)
}

func ToRest(status int, message string, data interface{}) map[string]interface{} {
    if message == "" {
        message = defaultResponse[status]
    }

    return map[string]interface{}{
        "status": status,
        "message": message,
        "data": data,
    }
}