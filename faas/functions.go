package faas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// GetFunctions TODO
func GetFunctions(baseURL string, target interface{}) error {
	fnURL := strings.TrimRight(baseURL, "/") + "/system/functions"
	resp, err := myClient.Get(fnURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// DeleteFunction TODO
func DeleteFunction(baseURL string, name string) error {
	fnURL := strings.TrimRight(baseURL, "/") + "/system/functions"

	deleteBody := fmt.Sprintf("{\"functionName\": \"%s\"}", name)

	req, err := http.NewRequest("DELETE", fnURL, strings.NewReader(deleteBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
