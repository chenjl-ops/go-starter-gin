package requests

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

// 统一请求json格式数据
/*
usage:

type Data struct {
	Name string `json:"name"`
    Age  int    `json:"age"`
}

type D struct {
	Code     string `json:"code"`
    Message  string `json:"message"`
    Data     []Data `json:"data"`
}

var data D

url := fmt.Sprintf("http://xxxx.xxx.xxx")
err := Request(url, &data)
if err != nil {
	fmt.Println(err)
}
fmt.Println(data)

*/
func Request(url string, data interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Println("请求失败: ", err)
		return err
	}
	defer resp.Body.Close()

	err1 := json.NewDecoder(resp.Body).Decode(data)
	if err1 != nil {
		logrus.Println("解析失败: ", err1)
		return err1
	}
	return nil
}
