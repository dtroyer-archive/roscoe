// auth

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
	"errors"
)

// Identity Types

type Auth struct {
    Access struct {
        Token Token
        User interface{}
        ServiceCatalog []ServiceCatalogEntry
    }
}


type Token struct {
	Id      string
	Expires string
	Project struct {
		Id   string
		Name string
	}
}

type ServiceCatalogEntry struct {
    Name string
    Type string
    Endpoints []map[string]string
}


type ServiceEndpoint struct {
    Type string
    Region string
    URL string
    VersionId string
}


func (self Auth) GetTokenId() string {
	return self.Access.Token.Id
}

func (self Auth) GetEndpoint(serviceType string, regionName string) (string, error) {

    // Parse service catalog
    for _, v := range self.Access.ServiceCatalog {
    	if v.Type == serviceType {
    		for _, r := range v.Endpoints {
    			if r["region"] == regionName {
    				return r["publicURL"], nil
    			}
    		}
    	}
    }
    return "", errors.New("err: endpoint not found")
}
