package main

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "bytes"
)

func PostUriString (strUrl string , postDict map[string]string) string {
    var httpReq *http.Request
    var respHtml string = ""

    postValues := url.Values{}
    for postKey, PostValue := range postDict {
        postValues.Set(postKey, PostValue)
    }

    postDataStr := postValues.Encode()
    postDataBytes := []byte(postDataStr)
    postBytesReader := bytes.NewReader(postDataBytes)
    client := &http.Client{}
    httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
    httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    httpResp, err := client.Do(httpReq)
    if err != nil {
      fmt.Println("error do")
      return ""
    }

    defer httpResp.Body.Close()

    body, _ := ioutil.ReadAll(httpResp.Body)
    respHtml = string(body)

    return respHtml

}

func main() {
    postDict := map[string]string{}
    uri := "http://www.baidu.com/"

    fmt.Println("fmt: ", PostUriString(uri, postDict))
}
