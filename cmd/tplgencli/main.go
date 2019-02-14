package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/theillego/tplgen"
	"log"
	"net/http"
	"os"
	"text/template"
)

type TestStruct struct {
	Hello          string `json:"hello"`
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

	url := fmt.Sprintf("https://www.mocky.io/v2/5185415ba171ea3a00704eed")
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

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

	// Fill the record with the data from the JSON
	var record TestStruct

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println(record.Hello)
	//fmt.Println(resp)
	//fmt.Println(json.MarshalIndent(resp, "", "    "))
}


