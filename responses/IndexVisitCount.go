package responses

import (
	"go-api-file/services"
	"strconv"
)

type IndexVisitCount struct {
	Code    string
	Message string
	Data    IndexVisitCountData
}

type IndexVisitCountData struct {
	Total int64
}

func (r *IndexVisitCount) Process() IndexVisitCount {
	file := services.Files{FileName: "visit_count.txt"}
	str := file.ReadAll()
	count, _ := strconv.ParseInt(str, 10, 64)
	count = count + 1
	r.Message = "success"
	r.Code = "0"
	r.Data.Total = count
	file.Write(strconv.FormatInt(count, 10), "w")
	return *r
}
