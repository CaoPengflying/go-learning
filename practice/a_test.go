// @Description main
// @Author caopengfei 2022/3/24 10:26

package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"testing"
)

func TestFun(t *testing.T) {
	if strings.Contains(strings.ToLower("https://t.100tal.com/avatar/%E5%AE%B6%E5%AE%9D"), ".jpg") {
		t.Log("ok")
	}
	t.Log("fail")

}

type GetPlanInfosByIdsResponse map[string]GetPlanInfosByIdsItem

type GetPlanInfosByIdsItem struct {
	Pattern int    `json:"pattern,string"`
	Name    string `json:"name"`
}

func TestF(t *testing.T) {
	resp := GetPlanInfosByIdsResponse{}
	err := jsoniter.UnmarshalFromString("{\"1420694\":{\"pattern\":\"\",\"name\":\"null\"}}", &resp)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v",resp)
}
