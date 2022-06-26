package wraps

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WrapErrorBadRequest(w http.ResponseWriter, err error) {
	WrapError(w, err, http.StatusBadRequest)
}

func WrapErrorServer(w http.ResponseWriter, err error) {
	WrapError(w, err, http.StatusInternalServerError)
}

func WrapError(w http.ResponseWriter, err error, httpStatus int) {
	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, err := json.Marshal(m)
	fmt.Println(err)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	_, err = fmt.Fprintln(w, string(res))
	if err != nil {
		fmt.Println(err)
	}

}

func WrapOK(w http.ResponseWriter, message map[string]interface{}) {
	res, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintln(w, string(res))
	if err != nil {
		fmt.Println(err)
	}
}
