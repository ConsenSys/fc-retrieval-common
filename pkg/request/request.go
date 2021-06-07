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

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
)

var httpClient = &http.Client{Timeout: 180 * time.Second}

// GetJSON request Get JSON
func GetJSON(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// SendJSON request Send JSON
func SendJSON(url string, data interface{}) error {
	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")

	r, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// SendMessage request Send JSON
func SendMessage(url string, message *fcrmessages.FCRMessage) (*fcrmessages.FCRMessage, error) {
	var data fcrmessages.FCRMessage
	jsonData, _ := json.Marshal(message)
	logging.Info("Sending JSON to url: %v", url)
	contentReader := bytes.NewReader(jsonData)
	req, err := http.NewRequest("POST", "http://"+url+"/v1", contentReader)
	req.Header.Set("Content-Type", "application/json")

	r, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return &data, err
	}

	if r.StatusCode != 200 {
		err := errors.New("SendMessage receive error code: " + r.Status)
		return nil, err
	}

	json.Unmarshal(bytes, &data)
	return &data, nil
}
