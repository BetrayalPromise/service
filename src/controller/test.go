/// 测试接口是否通畅

package controller

import (
	"fmt"
	"net/http"
)

func API(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "AAAAAA")
}