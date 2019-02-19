package main

import (
	"crypto/tls"
	"fmt"
	"github.com/theillego/apihelpers"
	"github.com/theillego/netfuncs"
	"github.com/theillego/tplgen"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
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

	// disable secure cert
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

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
	functionMap := template.FuncMap{
		"RemoveCidr":               tplgen.RemoveCidr,
		"ConvertIpCidrToIpNetmask": netfuncs.ConvertIpCidrToIpNetmask,
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

	url := "http://www.mocky.io/v2/5c6b1064330000a5387f4e89"

	reqBody := []byte(``)
	caw, _, _ := apihelpers.GetJson(url, reqBody)
	//fmt.Println(caw)

	for _, v := range caw {
		//fmt.Printf("%s %s\n", k, v)

		switch v := v.(type) {
		default:
			fmt.Printf("unexpected type %T", v)
		case uint64:
			fmt.Println("unit64")
		case string:
			fmt.Println("string")
		case float64, float32:
			fmt.Printf("%f\n", v) // don't print scientifically
		case bool:
			fmt.Println("bool")
		}


	}

	baw := caw["a_number"].(float64) // type assertion here to convert interface{} to float64
	baw = baw + 900000000000
	fmt.Println(baw)


	//pp, _ := PrettyPrintJson(respBody)
	//fmt.Print(pp)


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


