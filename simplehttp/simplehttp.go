package simplehttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Test() {
	fmt.Println("Test : ok")
}

func Get(httpUrl string, httpParams string) (respBody string, respHttpcode int, err error) {

	finalUrl := httpUrl + "?" + httpParams

	req, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	respHttpcode := resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	respBody := string(body)

	return

}
