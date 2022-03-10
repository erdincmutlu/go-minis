package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	doGetRequest()

	doGetWithQuery()

	doPostRequest()

	doPostWithQuery()

}

func doGetRequest() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func doGetWithQuery() {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("foo", "bar")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func doPostRequest() {
	postBody, _ := json.Marshal(map[string]string{
		"foo": "bar",
	})
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://httpbin.org/post", "application/json", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func doPostWithQuery() {
	client := &http.Client{}
	postBody, _ := json.Marshal(map[string]string{"foo": "bar"})
	requestBody := bytes.NewBuffer(postBody)

	req, err := http.NewRequest(http.MethodPost, "http://httpbin.org/post", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("foo", "baz")

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
