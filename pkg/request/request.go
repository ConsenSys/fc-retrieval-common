/*
Package request - contains common methods for JSON over HTTP communications
*/
package request

/*
 * Copyright 2020 ConsenSys Software Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

// do not remove
//go:generate mockgen -destination=../mocks/mock_request.go -package=mocks github.com/ConsenSys/fc-retrieval-common/pkg/request HttpCommunications

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
)

type HttpCommunicator struct {
	httpClient *http.Client
	sync.Mutex
}

// HttpCommunications - facilitates communications between nodes using HTTP
type HttpCommunications interface {
	GetJSON(url string) ([]byte, error)
	SendJSON(url string, data interface{}) error
	SendMessage(url string, message *fcrmessages.FCRMessage) (*fcrmessages.FCRMessage, error)
	Delete(url string) error
}

func NewHttpCommunicator() HttpCommunications {
	return &HttpCommunicator{
		httpClient: &http.Client{Timeout: 180 * time.Second},
		Mutex:      sync.Mutex{},
	}
}

// GetJSON request Get JSON
func (c *HttpCommunicator) GetJSON(url string) ([]byte, error) {
	c.Lock()
	defer c.Unlock()
	r, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	result, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		return nil, fmt.Errorf("can't read from JSON response, error: %s", readErr.Error())
	}

	if closeErr := r.Body.Close(); closeErr != nil {
		return nil, closeErr
	}
	return result, nil
}

// SendJSON request Send JSON
func (c *HttpCommunicator) SendJSON(url string, data interface{}) error {
	c.Lock()
	defer c.Unlock()
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	if req == nil {
		return errors.New("SendJSON error, can't create request")
	}
	req.Header.Set("Content-Type", "application/json")

	r, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	return nil
}

// SendMessage request Send JSON
func (c *HttpCommunicator) SendMessage(url string, message *fcrmessages.FCRMessage) (*fcrmessages.FCRMessage, error) {
	c.Lock()
	defer c.Unlock()
	var data fcrmessages.FCRMessage
	jsonData, _ := json.Marshal(message)
	logging.Info("Sending JSON to url: %v", url)
	contentReader := bytes.NewReader(jsonData)
	req, err := http.NewRequest("POST", "http://"+url+"/v1", contentReader)
	if req == nil {
		return nil, errors.New("SendMessage error, can't create request")
	}
	req.Header.Set("Content-Type", "application/json")

	r, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &data, err
	}

	if r.StatusCode != 200 {
		err := errors.New("SendMessage receive error code: " + r.Status)
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		return nil, errors.New("SendMessage error, can't unmarshal request body")
	}
	if err := r.Body.Close(); err != nil {
		return &data, errors.New("SendMessage error, can't close request body")
	}
	return &data, nil
}

// Delete request to delete a resource
func (c *HttpCommunicator) Delete(url string) error {
	c.Lock()
	defer c.Unlock()
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	r, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("can't delete, got http status: %d", r.StatusCode)
	}
	return nil
}
