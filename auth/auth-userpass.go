// auth-userpass

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
    "encoding/json"
    "errors"
    "strings"

	"roscoe/session"
)


type UserPassV2 struct {
    OSAuth struct {
        PasswordCredentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        } `json:"passwordCredentials"`
        ProjectName string `json:"tenantName"`
    } `json:"auth"`
    AuthUrl string `json:"-"`
}

func (self *UserPassV2) GetAuthRef() (*Auth, error) {
	return getAuth(self.AuthUrl, self.JSON())
}

// Produce JSON output
func (self *UserPassV2) JSON() ([]byte) {
    reqAuth, err := json.Marshal(self)
    if err != nil {
        // Return an empty structure
        reqAuth = []byte{'{', '}'}
    }
    return reqAuth
}

func (self *UserPassV2) AuthUserPassV2(opts interface{}) (session.TokenInterface, error) {
	auth, err := self.GetAuthRef()
	return session.TokenInterface(auth), err
}


type UserPassV3 struct {
    OSAuth struct {
        PasswordCredentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        } `json:"passwordCredentials"`
        ProjectName string `json:"projectName"`
    } `json:"auth"`
    AuthUrl string `json:"-"`
}

func (self *UserPassV3) GetAuthRef() (*Auth, error) {
	return getAuth(self.AuthUrl, self.JSON())
}

// Produce JSON output
func (self *UserPassV3) JSON() ([]byte) {
    reqAuth, err := json.Marshal(self)
    if err != nil {
        // Return an empty structure
        reqAuth = []byte{'{', '}'}
    }
    return reqAuth
}

func (self *UserPassV3) AuthUserPassV3(opts interface{}) (session.TokenInterface, error) {
	auth, err := self.GetAuthRef()
	return session.TokenInterface(auth), err
}


// Basic auth call
// These args should be an interface??
func getAuth(url string, body []byte) (*Auth, error) {
	var auth = Auth{}

	path := url + "/tokens"
	resp, err := session.Post(path, nil, &body)
	if err != nil {
		return &Auth{}, err
	}

	contentType := strings.ToLower(resp.Resp.Header.Get("Content-Type"))
	if strings.Contains(contentType, "json") != true {
		return &Auth{}, errors.New("err: header Content-Type is not JSON")
	}

	if err = json.Unmarshal(resp.Body, &auth); err != nil {
		return &Auth{}, err
	}

	return &auth, nil
}
