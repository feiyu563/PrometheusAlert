package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Endpoints struct {
	GlobalEndpoints         map[string]string            `json:"global_endpoints"`
	LocationCodeMapping     map[string]string            `json:"location_code_mapping"`
	RegionalEndpointPattern map[string]string            `json:"regional_endpoint_pattern"`
	Regions                 []string                     `json:"regions"`
	RegionalEndpoints       map[string]map[string]string `json:"regional_endpoints"`
	DocumentID              map[string]string            `json:"document_id"`
}

type RealEndpoints struct {
	Products []Product `json:"products"`
}

type Product struct {
	Code                    string             `json:"code"`
	DocumentID              string             `json:"document_id"`
	LocationServiceCode     string             `json:"location_service_code"`
	RegionalEndpoints       []RegionalEndpoint `json:"regional_endpoints"`
	GlobalEndpoint          string             `json:"global_endpoint"`
	RegionalEndpointPattern string             `json:"regional_endpoint_pattern"`
}

type RegionalEndpoint struct {
	Region   string `json:"region"`
	Endpoint string `json:"endpoint"`
}

//EndpointHandle process with endpoint
func endpointHandle(args []string) error {
	if len(os.Args) < 2 {
		return nil
	}
	switch args[1] {
	case "parse":
		if len(args) != 4 {
			return errors.New("The parameter is incorrect")
		}
		data, err := endpointParse(args[2])
		if err != nil {
			return err
		}
		err = generatEndpointsConfigFile(data, args[3])
		if err != nil {
			return err
		}
	}
	return nil
}

//endpointParse parse endpoint from java json to golang json
func endpointParse(srcpath string) (string, error) {

	_, err := os.Stat(srcpath)
	if err != nil {
		return "", errors.New("Source file error")
	}
	data, err := ioutil.ReadFile(srcpath)
	if err != nil {
		return "", err
	}
	endponit := &Endpoints{}
	err = json.Unmarshal(data, endponit)
	if err != nil || len(endponit.DocumentID) == 0 {
		return "", err
	}
	realEndpoints := &RealEndpoints{}
	for key := range endponit.GlobalEndpoints {
		if endponit.DocumentID[key] == "" {
			endponit.DocumentID[key] = "sdk"
		}
	}

	for key := range endponit.RegionalEndpointPattern {
		if endponit.DocumentID[key] == "" {
			endponit.DocumentID[key] = "sdk"
		}
	}

	for key := range endponit.RegionalEndpoints {
		if endponit.DocumentID[key] == "" {
			endponit.DocumentID[key] = "sdk"
		}
	}
	for key, value := range endponit.DocumentID {
		realEndpoint := Product{
			Code:                    key,
			LocationServiceCode:     key,
			DocumentID:              value,
			GlobalEndpoint:          endponit.GlobalEndpoints[key],
			RegionalEndpointPattern: endponit.RegionalEndpointPattern[key],
		}
		if realEndpoint.DocumentID == "sdk" {
			realEndpoint.DocumentID = ""
		}
		for key1, value1 := range endponit.LocationCodeMapping {
			if value1 == key {
				realEndpoint.Code = key1
			}
		}
		for key2, value2 := range endponit.RegionalEndpoints[key] {
			regionalEndpoint := RegionalEndpoint{
				Region:   key2,
				Endpoint: value2,
			}

			realEndpoint.RegionalEndpoints = append(realEndpoint.RegionalEndpoints, regionalEndpoint)
		}
		realEndpoints.Products = append(realEndpoints.Products, realEndpoint)
	}
	byte, err := json.MarshalIndent(realEndpoints, "", "\t")
	if err != nil {
		return "", err
	}
	return string(byte), err

}

func generatEndpointsConfigFile(data string, path string) error {

	lastData := `
package endpoints

import (
	"encoding/json"
	"fmt"
	"sync"
)

const endpointsJson =` + "`" + data + "`" + `
var initOnce sync.Once
var data interface{}

func getEndpointConfigData() interface{} {
	initOnce.Do(func() {
		err := json.Unmarshal([]byte(endpointsJson), &data)
		if err != nil {
			panic(fmt.Sprintf("init endpoint config data failed. %s", err))
		}
	})
	return data
}
`
	desfile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer desfile.Close()
	_, err = desfile.WriteString(lastData)
	if err != nil {
		return err
	}
	return nil
}
