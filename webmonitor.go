package webmonitor

import (
	"net/http"
)

type NilResponseError struct {
}

func (e NilResponseError) Error() string {
	return "HTTP Response is nil"
}



func Get(url string) (statusCode int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	} else {
		if resp != nil {
			defer resp.Body.Close()
			return resp.StatusCode, nil
		} else {
			return 0, NilResponseError{}
		}
	}
}

func Check(handler func) {
}

// Use for passing into Check(handler func)
func CheckAlive(url string) bool {
	result, err := Get(url)
	if err != nil {
		return false
	}
	return true
}

func Alert(handler func) {
}

// Use for passing into Alert(handler func)
func SendText(cellphoneNum string) {
}

// Use for passing into Alert(handler func)
func SendEmail(emailAddress string) {
}
