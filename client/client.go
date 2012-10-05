// client.go

package client

import (
    "bytes"
    "encoding/json"
//    "fmt"
    "io/ioutil"
    "net/http"
    "net/http/httputil"

    "roscoe/osclib"
)

type Client struct {
    httpClient *http.Client
    Auth osclib.Creds
    Token Token
    ServCat []ServiceCatalogEntry
}

type IdentTokens struct {
    Access struct {
        Token Token
        User interface{}
        ServiceCatalog []ServiceCatalogEntry
    }
}

type Token struct {
    Expires string
    Id string
    Tenant struct {
        Id string
        Name string
    }
}

type ServiceCatalogEntry struct {
    Name string
    Type string
    Endpoints []map[string]string
}

func NewClient(creds osclib.Creds) (oscc *Client, err error) {
    // Get credentials
    if creds.OSAuth.PasswordCredentials.Username == "" {
        err := creds.GetEnv()
        if err != nil {
            return nil, err
        }
    }
    oscc = &Client{
        httpClient: &http.Client{},
        Auth: creds,
    }
    return oscc, nil
}

// Authenticate to Identity service
func (c *Client) Authenticate() (err error) {
    // Can we inspect the token to see if it has expired?
    // Check service catalog here too
    if c.Token.Id == "" {
        err := c.getToken()
        if err != nil {
            // auth failures will appear here
            return err
        }
    }
    return err
}

func (c *Client) getToken() (err error) {
    // Build the request body
    req, err := http.NewRequest(
        "POST",
        c.Auth.AuthUrl + "/tokens",
        bytes.NewBuffer(c.Auth.JSON()),
    )
    if err != nil {
        return err
    }

    req.Header.Add("content-type", "application/json")

    // TODO(dtroyer): work out debug semantics
    if 1 == 0 {
        d, _ := httputil.DumpRequestOut(req, true)
        print("req: ", string(d), "\n\n")
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        // TODO(dtroyer): Handle specific errors
        return err
    }
    defer resp.Body.Close()

    // TODO(dtroyer): work out debug semantics
    if 1 == 0 {
        dr, _ := httputil.DumpResponse(resp, true)
        print("resp: ", string(dr), "\n\n")
    }

    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return err
    }

    var access IdentTokens
    err = json.Unmarshal(contents, &access)
    if err != nil {
        return err
    }

    c.Token = access.Access.Token
    c.ServCat = access.Access.ServiceCatalog
    return nil
}

