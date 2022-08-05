package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	// baca json, karena request.Body mrpkn io.Read maka kita bisa stream datanya memakai decoder
	decoder := json.NewDecoder(request.Body)
	// decode data
	err := decoder.Decode(result)
	PanicIfError(err)
}

func Write2ResponseBody(writer http.ResponseWriter, response interface{}) {
	// tambah header
	writer.Header().Add("Content-Type", "Application/json")
	// convert data jd json as response data ke user
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
