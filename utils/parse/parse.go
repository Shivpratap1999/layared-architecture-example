package parse

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseReqBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	dataByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(dataByte), v); err != nil {
		return err
	}
	return err
}
