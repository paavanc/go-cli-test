package utils

import (
	"bytes"
	"encoding/json"
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func HomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err.Error())
	}
	return home
}

var src = rand.NewSource(time.Now().UnixNano())

func RestHelper(method string, url string, body interface{},headers map[string]string) map[string]interface{}{

	bytesRepresentation, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}


	req, err:= http.NewRequest(method, url, bytes.NewReader(bytesRepresentation))
	if err!=nil{
		panic(err)
	}


	for k, v := range headers{
		req.Header.Set(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		panic( err)
	}
	result, err:= ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}

	var dat map[string]interface{}

	if err := json.Unmarshal(result, &dat); err != nil {
		panic(err)
	}
	return dat
}

