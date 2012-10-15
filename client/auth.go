// client/auth.go - Client Authentication

package client

import (
    "encoding/json"
)

// Credentials

// The 'password flow' credentials information provided by the user
// and used to aquire a token from the Identity service
type Credentials struct {
    OSAuth struct {
        PasswordCredentials struct {
            Username string `json:"username"`
            Password string `json:"password"`
        } `json:"passwordCredentials"`
        TenantName string `json:"tenantName"`
    } `json:"auth"`
    AuthUrl string `json:"-"`
}

// Produce JSON output
func (c *Credentials) JSON() ([]byte) {
    reqAuth, err := json.Marshal(c)
    if err != nil {
        // Return an empty structure
        reqAuth = []byte{'{', '}'}
    }
    return reqAuth
}

