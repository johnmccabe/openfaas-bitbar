package faas

import (
	"encoding/json"
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
