package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//请求路径正则表达式
const reqPathPattern = "^/(?P<type>\\w+)/(?P<target>\\w+)/(?P<action>\\w+)(\\?(?P<query>.*))?"

//视频发布
func PublishAction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PubilshAction")
}

//视频列表
func PublishList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PublishList")
}
func serve(w http.ResponseWriter, r *http.Request) {

	regPath := regexp.MustCompile(reqPathPattern)
	//key
	fieldNames := regPath.SubexpNames()
	//value
	fields := regPath.FindStringSubmatch(r.URL.Path)
	result := make(map[string]string)
	for i, name := range fieldNames {
		if i != 0 && name != "" {
			result[name] = fields[i]
		}
	}
	reqType := strings.ToLower(result["type"])
	target := strings.ToLower(result["target"])
	action := strings.ToLower(result["action"])
	if reqType == "" || target == "" || action == "" {
		fmt.Println("Requset Error!!!")
		return
	}
	fmt.Printf("type:%s, target:%s, actioin:%s \n", reqType, target, action)

	//拼接请求路径
	req := fmt.Sprintf("/%s/%s/%s", reqType, target, action)

	//switch分发处理函数
	switch req {
	case "/douyin/publish/action":
		PublishAction(w, r)

	case "/douyin/publish/list":
		PublishList(w, r)
	}
}

func main() {
	fmt.Println("启动！！")
	http.HandleFunc("/", serve)
	//监听8080端口
	//http://192.168.137.224
	http.ListenAndServe(":8080", nil)
}
