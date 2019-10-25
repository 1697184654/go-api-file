package responses

import (
	"go-api-file/requests"
	"go-api-file/services"
)

type IndexSignUp struct {
	Code    string
	Message string
	Data    []string
}

func (r *IndexSignUp) Process(request requests.UserInfo) IndexSignUp {
	record := "\"" + request.UserName + "\",\"" + request.Email + "\",\"" + request.Tel + "\"\n"
	file := services.Files{FileName: "./sign_up.txt"}
	file.Write(record, "a+")
	r.Code = "0"
	r.Message = "success"
	return *r
}
