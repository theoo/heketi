//
// Copyright (c) 2014 The heketi Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func jsonFromBody(r io.Reader, n int64, v interface{}) error {

	// Check body
	body, err := ioutil.ReadAll(io.LimitReader(r, n))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

func GetJsonFromRequest(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return jsonFromBody(r.Body, r.ContentLength, v)
}

func GetJsonFromResponse(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	return jsonFromBody(r.Body, r.ContentLength, v)
}
