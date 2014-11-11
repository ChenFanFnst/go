package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func PostUriString(strUrl string, postDict map[string]string) string {
	var httpReq *http.Request
	var respHtml string = ""

	postValues := url.Values{}
	for postKey, postValue := range postDict {
		postValues.Set(postKey, postValue)
		fmt.Println("key:" + postKey + ",value:" + postValue)
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

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func writeToFile(values string, outfile string) error {
	var file *os.File
	var err error

	if !isExist(outfile) {
		file, err = os.Create(outfile)
		if err != nil {
			fmt.Println("Failed to create the output file", outfile)
			return err
		}
	} else {
		file, err = os.OpenFile(outfile, os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Failed to open the output file", outfile)
			return err
		}
		fmt.Println("witofile:", values)
	}

	defer file.Close()
	file.WriteString(values + "\n")
	return nil
}

func main() {
	userFile := "config.txt"
	outFile := "results.txt"

	fout, err := os.Open(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()

	buf := bufio.NewReader(fout)
	for {
		uri := ""
		params := ""
		postDict := map[string]string{}

		line, err := buf.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		for i, v := range strings.Split(line, "\t") {
			if i == 0 {
				uri = v
			} else if i%2 == 1 {
				params = v
			} else if i%2 == 0 {
				postDict[params] = v
			}

		}

		resultXml := PostUriString(uri, postDict)
		writeToFile(resultXml, outFile)
		fmt.Println("fmt: ", resultXml)
	}
}
