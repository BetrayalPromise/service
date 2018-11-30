package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"utils"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		str, err := utils.Unsuport()
		if err != nil {
			fmt.Print(err)
			return
		}
		fmt.Fprint(w, str)
	} else {
		data, err := ioutil.ReadFile("src/controller/result.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		var jsonInstance map[string]interface{}
		var jsonError = json.Unmarshal(data, &jsonInstance)
		if jsonError != nil {
			fmt.Println(err)
			return
		}
		json, err := json.Marshal(jsonInstance)
		if err != nil {
			fmt.Println("MapToJsonDemo err: ", err)
			return
		}
		// 模拟网络情况不好
		time.Sleep(1 * time.Second)
		fmt.Fprint(w, string(json))
	}

}