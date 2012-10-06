// client.go

package client

import (
    "bytes"
    "encoding/json"
//    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "net/http/httputil"

    "roscoe/osclib"
)


// TODO(dtroyer): Is there a better way to do this to pick up
//                or distribute globals?
var Debug = new(bool)

// Identity Types

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

type ServiceEndpoint struct {
    Type string
    Region string
    URL string
    VersionId string
}


// Client

type Client struct {
    httpClient *http.Client
    Auth osclib.Creds
    Token Token
    ServCat map[string]ServiceEndpoint
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
    oscc.Authenticate()
    return oscc, nil
}

func (c *Client) NewRequest(method, url string, body io.Reader) (req *http.Request, err error) {
    req, err = http.NewRequest(method, url, body)
    if err != nil {
		return
	}
	// add token, get one if needed
	if c.Token.Id == "" {
        err = c.Authenticate()
        if err != nil {
            return nil, err
        }
    }
    req.Header.Add("X-Auth-Token", c.Token.Id)
	return
}

func (c *Client) Do(req *http.Request) (res *Result, err error) {
    if *Debug {
        d, _ := httputil.DumpRequestOut(req, true)
        print("\nREQUEST:\n----------\n", string(d), "\n----------\n")
    }

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
    if *Debug {
        dr, _ := httputil.DumpResponse(resp, true)
        print("\nRESULT:\n----------\n", string(dr), "\n----------\n")
    }

	res = NewResult(resp)
	return
}

// Authenticate to Identity service
func (c *Client) Authenticate() (err error) {
    // TODO(dtroyer): Inspect the token to see if it has expired if possible
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

// Perform a simple get to an endpoint selected
// by 'api' from the service catalog
func (c *Client) Get(api string, url string) (resp *Result, err error) {
    req, err := c.NewRequest("GET", c.ServCat[api].URL + url, nil)
    if err != nil {
		return nil, err
    }

    resp, err = c.Do(req)
    if err != nil {
		return nil, err
    }
    // do we need to parse this in this func? yes...
    defer resp.HResponse.Body.Close()

    resp.Body, err = ioutil.ReadAll(resp.HResponse.Body)
    if err != nil {
		return nil, err
    }

    return resp, nil
}

// Retrieve a token from the Identity service
func (c *Client) getToken() (err error) {
    // Build the request body
    // Call the http method directly to avoid a recursive call
    // back here if no current token is held
    req, err := http.NewRequest(
        "POST",
        c.Auth.AuthUrl + "/tokens",
        bytes.NewBuffer(c.Auth.JSON()),
    )
    if err != nil {
        return err
    }

    req.Header.Add("content-type", "application/json")

    resp, err := c.Do(req)
    if err != nil {
        // TODO(dtroyer): Handle specific errors
        return err
    }
    defer resp.HResponse.Body.Close()

    contents, err := ioutil.ReadAll(resp.HResponse.Body)
    if err != nil {
        return err
    }

    var access IdentTokens
    err = json.Unmarshal(contents, &access)
    if err != nil {
        return err
    }

    c.Token = access.Access.Token

    // Parse service catalog
    c.ServCat = map[string]ServiceEndpoint{}
    for _, v := range access.Access.ServiceCatalog {
        // we only look at the first endpoint for each type
        c.ServCat[v.Type] = ServiceEndpoint{
            Type: v.Type,
            Region: v.Endpoints[0]["region"],
            URL: v.Endpoints[0]["publicURL"],
            VersionId: v.Endpoints[0]["versionId"],
        }
    }

    return nil
}


// Result

type Result struct {
    HResponse *http.Response
    Body []byte
}

func NewResult(resp *http.Response) *Result {
	result := new(Result)
	result.HResponse = resp

	//result.parseHeader()
	//result.parseBody()

	return result
}
