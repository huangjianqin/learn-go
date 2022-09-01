package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Human struct {
	//必须大写, 不然json module访问不到
	Name string
	Age  int
	E    Email
}

type Email struct {
	Addr       string
	CreateTime time.Time
}

func main() {
	json1()

	json2()
}

func json1() {
	human := Human{"abc", 10, Email{"abss@123.com", time.Now()}}
	bytes, _ := json.Marshal(human)
	fmt.Printf("%v\n", string(bytes))
}

func json2() {
	bs := []byte(`{"Name":"abc","Age":10,"E":{"Addr":"abss@123.com","CreateTime":"2022-09-01T09:43:09.841066+08:00"}}`)
	var h Human
	var jsonMap map[string]interface{}
	json.Unmarshal(bs, &h)
	json.Unmarshal(bs, &jsonMap)
	fmt.Printf("%v\n", h)
	fmt.Printf("%v\n", jsonMap)

	for k, v := range jsonMap {
		//interface{}可以代表任意类型, 通过v.(type)进行类型转换, 也可以像下面那样switch匹配类型
		switch v.(type) {
		case map[string]interface{}:
			//乐行
			im := v.(map[string]interface{})
			for ik, iv := range im {
				fmt.Printf("%v\n", ik)
				fmt.Printf("%v\n", iv)
			}
		default:
			fmt.Printf("%v\n", k)
			fmt.Printf("%v\n", v)
		}
	}
}
