package simplehttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Test() {
	fmt.Println("Test : ok")
}

func Get(httpUrl string, httpParams string) (respHttpcode int, respBody string, err error) {

	finalUrl := httpUrl
	if httpParams != "" {
		finalUrl = httpUrl + "?" + httpParams
	}

	req, err := http.NewRequest("GET", finalUrl, nil)

	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	respHttpcode = resp.StatusCode
	respBody = string(body)

	return

}

func Post(httpUrl string, httpData string) (respHttpcode int, respBody string, err error) {

	finalUrl := httpUrl

	req, err := http.NewRequest("POST", finalUrl, strings.NewReader(httpData))

	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return
	}

	respHttpcode = resp.StatusCode
	respBody = string(body)

	return

}
