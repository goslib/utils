package utlHttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string) (error, int, []byte) {
	res, err := http.Get(url)
	return doHttpCall(res, err)
}

// HTTP Post Request
func Post(url string, body interface{}) (error, int, []byte) {
	bts, err := json.Marshal(body)
	if err != nil {
		return err, http.StatusBadRequest, nil
	}
	return Post(url, bts)
}

// The real function being called.
func DoPost(url string, bts []byte) (error, int, []byte) {
	res, err := http.Post(url, "application/json", bytes.NewReader(bts))
	return doHttpCall(res, err)
}

func Delete(url string) (error, int, []byte) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	res, err := http.DefaultClient.Do(req)
	return doHttpCall(res, err)
}

func DoHttpCall(req *http.Request) (error, int, []byte) {
	res, err := http.DefaultClient.Do(req)
	return doHttpCall(res, err)
}

// Error, StatusCode, Response Content
func doHttpCall(res *http.Response, err error) (error, int, []byte) {
	// Error with the network.
	if err != nil {
		return err, http.StatusInternalServerError, nil
	}
	// Check the logical status code.
	//if res.StatusCode != http.StatusOK {
	//	return errors.New("error status code"), res.StatusCode, nil
	//}
	body, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close() // defer close the body @see https://golang.org/pkg/net/http/#Get
	if err != nil {
		return err, res.StatusCode, body
	}
	return nil, res.StatusCode, body
}
