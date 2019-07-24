package jsontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BasicInfo struct {
	Name string `jsontest:"name"`
	Age  int    `jsontest:"age"`
}
type JobInfo struct {
	Skills []string `jsontest:"skills"`
}
type Employee struct {
	BasicInfo BasicInfo `jsontest:"basic_info"`
	JobInfo   JobInfo   `jsontest:"job_info"`
}

/*
	"detail" : {
		"zzk_age": 30,
		"zzk_name": "zzk"
	}
*/

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

func TestEmbeddedJson(t *testing.T) {
	// 序列化
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(e)
	// 反序列化
	if v, err := json.Marshal(e); err == nil {
		fmt.Println("---")
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}


