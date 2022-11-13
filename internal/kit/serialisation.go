package kit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GenericResponse is the format of our response
type GenericResponse struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func FromJSON(r *http.Response) (body []byte, err error) {
	var gr GenericResponse
	body, err = io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &gr)
	if err != nil {
		return body, nil
	}
	//log.Infof("GenericResponse: %v", gr)
	if gr.Code != 0 && gr.Code != 200 {
		return nil, fmt.Errorf("%s, %s", gr.Error, gr.Message)
	}
	return body, nil
}
