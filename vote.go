package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var PostURL string = "http://365jia.cn/zt/qnwmh/doVote"
var vote_data map[string]string = map[string]string{"id": "678998"}

type Msg struct {
	Status  int
	Message string
	Is_Code string
}

func main() {
	fmt.Println(PostURL)
	file, err := os.Open("./proxy.txt")
	if err != nil {
		fmt.Println("无法打开proxy.txt，请检查这个文件是否存在。", err)
		return
	}
	defer func() {
		file.Close()
	}()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxy := "http://" + strings.Split(scanner.Text(), "@")[0]
		go vote(proxy)
	}

	var info string
	for {
		fmt.Scanln(&info)
		if info == "理想要努力" {
			break
		}
	}
}

func vote(proxy string) {
	fmt.Println(proxy)
	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("无效代理", err)
		return
	}
	client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

	data := url.Values{"id": {"678998"}}
	req, err := http.NewRequest("POST", PostURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("无效发送数据", err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.114 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("无法连接", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var msg Msg
	err = json.Unmarshal(body, &msg)
	if err != nil {
		fmt.Println("解析错误", err)
		fmt.Println(string(body))
		return
	}
	//fmt.Println(string(body))
	fmt.Println(msg.Message)

}
