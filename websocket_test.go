package beechatt_socket

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInit(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:5120/payload", nil)
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	var objmap map[string]interface{}
	if err := json.Unmarshal([]byte(bodyString), &objmap); err != nil {
		assert.Equal(t, true, false, "Should not be error")
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expecting 200 status code")
	assert.Equal(t, "Welcome to home", objmap["message"].(string))
}

func TestSendPayload(t *testing.T) {
	payload := []byte(`{
		"from": "FAJRUL",
		"to": "WIDYA",
		"message": "I Love You",
		"callback": false
	}`)
	req, err := http.NewRequest("POST", "http://localhost:5120/payload", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	var objmap map[string]interface{}
	if err := json.Unmarshal([]byte(bodyString), &objmap); err != nil {
		assert.Equal(t, true, false, "Should not be error")
	}
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expecting 200 status code")
}
