package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
)

func main() {
	username := flag.String("username", "", "Username for the session")
	password := flag.String("password", "", "Passord for the session")
	flag.Parse()

	if *username == "" {
		fmt.Println("Error : Missing mandatory parameter username")
		flag.PrintDefaults()
		os.Exit(1)
	} else {
		if *password == "" {
			fmt.Println("Error : Missing mandatory parameter password")
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	osfciurl := "https://osfci.tech/user/"
	gettokenstring := "/getToken"
	gettokenurl := fmt.Sprintf("%s%s%s", osfciurl, *username, gettokenstring)

	fmt.Println(gettokenurl)

	client := &http.Client{}

	data := url.Values{}
	data.Add("password", *password)

	req, err := http.NewRequest("POST", gettokenurl, bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		//error handling
	}

	resp, err := client.Do(req)
	if err != nil {
                //error handling
        }
	fmt.Println("resp.Header=",resp.Header)
	//fmt.Println("resp.Status=",resp.Status)


	f, err := ioutil.ReadAll(resp.Body)
	if err != nil {
                //error handling
        }
	resp.Body.Close()

	fmt.Println(string(f))
}
