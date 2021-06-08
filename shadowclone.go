package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleShadowClone(c *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(404, "can't read request body")
	}
	webhooks := read()
	for _, wh := range webhooks {
		err := transferRequest(bodyBytes, wh)
		if err != nil {
			c.String(404, "transfer fail")
			return
		}
	}
	c.String(200, "transfer success")
}

func transferRequest(bodyBytes []byte, url string) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	client := &http.Client{}
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func read() []string {
	file, err := os.Open("./webhook_list")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Println(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	webhooks := []string{}
	for scanner.Scan() {
		webhooks = append(webhooks, scanner.Text())
	}
	return webhooks
}
