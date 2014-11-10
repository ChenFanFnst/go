package main

import (
    "io/ioutil"
    "net/http"
    "net/url"
    "fmt"
    "bytes"
    "os"
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
      return strUrl
    }

    defer httpResp.Body.Close()

    body, _ := ioutil.ReadAll(httpResp.Body)
    respHtml = string(body)

    return respHtml

}

func main() {
    postDict := map[string]string{}
    userFile := "./config.txt"
    uri      := ""

    fout, err := os.Open(userFile)
    if err != nil {
        fmt.Println(userFile, err)
        return
    }
    defer fout.Close()
    buf := make([]byte, 1024)
    for {
        n, _ := fout.Read(buf)
        if 0 == n {
            break
        }
        os.Stdout.Write(buf[:n])
        uri = string(buf[:n])
        fmt.Println("fmt: ", PostUriString(uri, postDict))
    }
}
