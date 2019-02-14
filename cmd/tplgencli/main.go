package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/theillego/tplgen"
	"log"
	"net/http"
	"os"
	"text/template"
)

type TestStruct struct {
	Hello          string `json:"hello"`
}

// For control over HTTP client headers,
// redirect policy, and other settings,
// create a Client
// A Client is an HTTP client

var client = &http.Client{
	Timeout: time.Second * 10, // always configure timeout
}

func main() {

	router1 := tplgen.Router{
		"ROUTER-1",
		"2ZNTEF29104F",
		"192.0.2.1/25",
	}

	// create new file
	file, err := os.Create(router1.Hostname)
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	// pass functions to be used in template
	functionMap := template.FuncMap {
		"RemoveCidr": tplgen.RemoveCidr,
		"ConvertIpCidrToIpNetmask": tplgen.ConvertIpCidrToIpNetmask,
	}

	// parse template and create new config file
	t := template.Must(template.New("base_config.tmpl").Funcs(functionMap).ParseFiles("base_config.tmpl"))
	err = t.Execute(file, tplgen.CreateRouterConfiguration(router1))
	if err != nil {
		panic(err)
	}

	// close new config file
	err = file.Close()
	if err != nil {
		panic(err)
	}

	//fmt.Println(tplgen.RemoveCidr("192.168.1.1/24"))

	url := "https://www.mocky.io/v2/5185415ba171ea3a00704eed"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	//fmt.Println(resp.Body)
	//fmt.Println(record.Hello)

	// getting json data without a struct
	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		panic(err)
	}

	//fmt.Println(data["hello"])

	// pretty print json string
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "\t")
	fmt.Println(string(prettyJSON.Bytes()))

	// Fill the record with the data from the JSON
	var record TestStruct

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
}



//func getJson(url, target interface{}) error {
//	r, err := client.Get(url)
//	if err != nil {
//		return err
//	}
//	defer r.Body.Close()
//
//	return json.NewDecoder(r.Body).Decode(target)
//}


