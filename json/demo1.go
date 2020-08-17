package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	resp := `{"code": "00",
              "message": "SUCCESS",
              "describe": "成功",
              "resultInfo": { "uniqueNumber": "201808161133401673324075025000035" } 
             }`

	type JsonResp struct {
		Code       int               `json:"code"`
		Message    string            `json:"message"`
		Describe   string            `json:"describe"`
		ResultInfo map[string]string `json:"resultInfo"`
	}
	var smsresp JsonResp
	temp := []byte(resp)
	errs := json.Unmarshal(temp, &smsresp)
	if errs != nil {
		return
	}
	fmt.Println(smsresp.Code)
	fmt.Println(smsresp.Describe)
	fmt.Println(smsresp.Message)
	fmt.Println(smsresp.ResultInfo["uniqueNumber"])
}
