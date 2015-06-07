package utils

import (
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
	"github.com/revel/revel"
	"io/ioutil"
	"strings"
)

func Phase(c *revel.Controller) (data map[string]interface{}) {
	//如果不是json格式的内容，直接返回
	if !strings.Contains(c.Request.ContentType, "application/json") {
		fmt.Println("Not a JSON")
		return
	}
	//else
	content, _ := ioutil.ReadAll(c.Request.Body)
	json, _ := simplejson.NewJson([]byte(content))
	// fmt.Printf("JSON is : %s\n", json)
	data, _ = json.Map()
	return
}
