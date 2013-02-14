// osclib.go - roscoe library package

package osclib

import (
    "bytes"
    "encoding/json"
    "fmt"
//    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "reflect"
    "strings"
)


// API Versions

type APIVersion struct {
    Id string
    Status string
    Updated string
}

type IdentityVersion struct {
    Versions struct {
        Values []APIVersion
    }
}

func GetVersions(auth Creds) {
    verurl, err := url.Parse(auth.AuthUrl)
    verurl.Path = ""

    resp, err := http.Get(verurl.String())
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var apiver IdentityVersion
    err = json.Unmarshal(body, &apiver)
    if err != nil {
        log.Fatal(err)
    }
    for k, v := range apiver.Versions.Values {
        print("Version(", k, "): ", v.Id, "\n")
    }
}


// Service Catalog and Token

type IdentTokens struct {
    Access struct {
        Token Token
        User interface{}
        ServiceCatalog []ServiceCatalogEntry
    }
}

type ServiceCatalogEntry struct {
    Name string
    Type string
    Endpoints []map[string]string
}

type Endpoint struct {
    AdminUrl string
}

type Token struct {
    Expires string
    Id string
    Tenant struct {
        Id string
        Name string
    }
}

type Auth struct {
    TenantName string
    PasswordCredentials struct {
        Username string
        Password string
    }
}

func GetToken(auth Creds) (token Token, sc []ServiceCatalogEntry, err error) {
    ic := &http.Client{}

    // Build the request body
    req, err := http.NewRequest(
        "POST",
        auth.AuthUrl + "/tokens",
        bytes.NewBuffer(auth.JSON()),
    )
    if err != nil {
        log.Fatal(err)
    }
    req.Header.Add("content-type", "application/json")
    d, err := httputil.DumpRequestOut(req, true)
    _ = d
    print("req: ", string(d), "\n\n")

    resp, err := ic.Do(req)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    dr, err := httputil.DumpResponse(resp, true)
    _ = dr
    print("resp: ", string(dr), "\n\n")

    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    var access IdentTokens
    err = json.Unmarshal(contents, &access)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", access)
    return access.Access.Token, access.Access.ServiceCatalog, err
}


// Output routines

func OutputData(headers []string, data []interface{}) (err error) {
    fmt.Print(strings.Join(headers, "\t"), "\n")
    for _, v := range data {
        d := reflect.ValueOf(v)
        fmt.Printf("%v+\n", d)
//        fmt.Print(strings.Join(d.String{}, "\t"), "\n")
    }
    return
}
