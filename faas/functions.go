package faas

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/johnmccabe/openfaas-bitbar/config"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// GetFunctions TODO
func GetFunctions(auth config.Auth, target interface{}) error {
	fnURL := strings.TrimRight(auth.Gateway, "/") + "/system/functions"

	req, err := http.NewRequest("GET", fnURL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth.Token))

	resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

// DeleteFunction TODO
func DeleteFunction(auth config.Auth, name string) error {
	fnURL := strings.TrimRight(auth.Gateway, "/") + "/system/functions"

	deleteBody := fmt.Sprintf("{\"functionName\": \"%s\"}", name)

	req, err := http.NewRequest("DELETE", fnURL, strings.NewReader(deleteBody))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth.Token))

	req.Header.Add("Content-Type", "application/json")

	resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
