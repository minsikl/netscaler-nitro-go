package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/minsikl/netscaler-nitro-go/datatypes"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

var resourceMap = map[string]string{
	"Service":                 "service",
	"Lbvserver":               "lbvserver",
	"LbvserverServiceBinding": "lbvserver_service_binding",
	"Systemfile":              "systemfile",
	"Nsfeature":               "nsfeature",
}

// Nitro Client
type NitroClient struct {
	Protocol  string
	IpAddress string
	Mode      string
	User      string
	Password  string
	Debug     bool
}

func (n *NitroClient) Add(req interface{}) error {
	resource, err := getResourceStringByObject(req)
	if err != nil {
		return err
	}
	reqJson, err := json.Marshal(req)

	if err != nil {
		return err
	}
	responseBody, _, err := HTTPRequest(n, resource, "POST", reqJson)
	if err != nil {
		return fmt.Errorf("Error in POST %s", err.Error())
	}
	if len(responseBody) > 0 {
		res := datatypes.BaseRes{}
		err = json.Unmarshal(responseBody, res)
		if err != nil {
			return fmt.Errorf("Error in Unmarshal %s", err.Error())
		}
		if res.Severity == "ERROR" {
			return fmt.Errorf("Error in POST : %+v", res)
		}
	}
	return nil
}

func (n *NitroClient) Get(res interface{}, resourceName string, filter string, attrs string) error {
	resource, err := getResourceStringByObject(res)
	if err != nil {
		return err
	}

	resourceQuery := resource
	if resourceName != "" {
		resourceQuery = resourceQuery + "/" + resourceName
		if attrs != "" {
			resourceQuery = resourceQuery + "?attrs=" + attrs
		}
	} else {
		if filter != "" {
			resourceQuery = resourceQuery + "?filter=" + filter
		}
	}

	responseBody, _, err := HTTPRequest(n, resourceQuery, "GET", nil)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, res)
	if err != nil {
		return fmt.Errorf("Error in Unmarshal %s", err.Error())
	}
	resMessage := datatypes.BaseRes{}
	err = json.Unmarshal(responseBody, &resMessage)
	if err != nil {
		return fmt.Errorf("Error in Unmarshal %s", err.Error())
	}
	if resMessage.Severity == "ERROR" {
		return fmt.Errorf("Error in POST : %+v", resMessage)
	}

	return nil
}

func (n *NitroClient) Delete(req interface{}, resourceName string, args string) error {
	resource, err := getResourceStringByObject(req)
	if err != nil {
		return err
	}

	resourceQuery := resource + "/" + resourceName
	if len(args) > 0 {
		resourceQuery = resourceQuery + "?args=" + args
	}

	responseBody, _, err := HTTPRequest(n, resourceQuery, "DELETE", nil)
	if err != nil {
		return err
	}
	resMessage := datatypes.BaseRes{}
	err = json.Unmarshal(responseBody, &resMessage)
	if resMessage.Severity == "ERROR" {
		return fmt.Errorf("Error in DELETE : %+v", resMessage)
	}

	return nil
}

func (n *NitroClient) Enable(req interface{}) error {
	resource, err := getResourceStringByObject(req)
	if err != nil {
		return err
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return err
	}
	responseBody, _, err := HTTPRequest(n, resource+"/?action=enable", "POST", reqJson)
	if err != nil {
		return fmt.Errorf("Error in POST %s for Enable", err.Error())
	}
	if len(responseBody) > 0 {
		res := datatypes.BaseRes{}
		err = json.Unmarshal(responseBody, res)
		if err != nil {
			return fmt.Errorf("Error in Unmarshal %s", err.Error())
		}

		if res.Severity == "ERROR" {
			return fmt.Errorf("Error in POST for Enable: %+v", res)
		}
	}
	return nil
}

func (n *NitroClient) Disable(req interface{}) error {
	resource, err := getResourceStringByObject(req)
	if err != nil {
		return err
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return err
	}
	responseBody, _, err := HTTPRequest(n, resource+"/?action=disable", "POST", reqJson)
	if err != nil {
		return fmt.Errorf("Error in POST %s for Disable", err.Error())
	}
	if len(responseBody) > 0 {
		res := datatypes.BaseRes{}
		err = json.Unmarshal(responseBody, res)
		if err != nil {
			return fmt.Errorf("Error in Unmarshal %s", err.Error())
		}

		if res.Severity == "ERROR" {
			return fmt.Errorf("Error in POST for Disable: %+v", res)
		}
	}
	return nil
}

func GetNitroClient(protocol string, ipAddress string, mode string, user string, password string, debug bool) *NitroClient {
	nClient := NitroClient{
		Protocol:  protocol,
		IpAddress: ipAddress,
		Mode:      mode,
		User:      user,
		Password:  password,
		Debug:     debug,
	}
	return &nClient
}

func HTTPRequest(nClient *NitroClient, resourceQuery string, requestType string, requestBody []byte) ([]byte, int, error) {

	// Create a request
	Url := nClient.Protocol + "://" + nClient.IpAddress + "/nitro/v1/" + nClient.Mode + "/" + resourceQuery
	requestBodyBuffer := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest(requestType, Url, requestBodyBuffer)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-NITRO-USER", nClient.User)
	req.Header.Set("X-NITRO-PASS", nClient.Password)

	if nClient.Debug {
		log.Println("[DEBUG] Nitro Request Path: ", requestType, req.URL)
		log.Println("[DEBUG] Nitro Request Parameters: ", requestBodyBuffer.String())
	}

	// Execute http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	if nClient.Debug {
		log.Println("[DEBUG] Nitro Response: ", string(responseBody))
	}
	return responseBody, resp.StatusCode, nil
}

func getResourceStringByObject(obj interface{}) (string, error) {
	resourceType := reflect.TypeOf(obj).Elem().Name()
	resource := resourceMap[resourceType[:len(resourceType)-3]]

	if len(resource) == 0 {
		return "", fmt.Errorf("Cannot find a resource name for struct '%s' \r\n", resourceType)
	}
	return resource, nil
}
