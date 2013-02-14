// token.go - roscoe token example

package main

import (
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "log"
    "os"

    "github.com/voxelbrain/goptions"

    "roscoe/client"
//    "roscoe/osclib"
)


func main() {
    options := struct {
        Debug bool      `goptions:"-x, --debug, description='Enable debugging'"`
        Verbose bool    `goptions:"-v, --verbose, description='Be not quiet with output'"`
    }{
        Debug: false,
        Verbose: false,
    }
    goptions.Parse(&options)

    // Propagate debug setting to packages
    client.Debug = &options.Debug

    // Get auth values from the environment
    var creds client.Credentials
    creds.AuthUrl = os.Getenv("OS_AUTH_URL")
    creds.OSAuth.TenantName = os.Getenv("OS_TENANT_NAME")
    creds.OSAuth.PasswordCredentials.Username = os.Getenv("OS_USERNAME")
    creds.OSAuth.PasswordCredentials.Password = os.Getenv("OS_PASSWORD")
    cacert = os.Getenv("OS_CACERT")

    // read cacert

    pool = NewCertPool()
    for j, root := range test.roots {
        ok := opts.Roots.AppendCertsFromPEM([]byte(root))
        if !ok {
            t.Errorf("#%d: failed to parse root #%d", i, j)
            return
        }
    }

    // Set up TLS config
    conf := &tls.Config{InsecureSkipVerify: true}

    c, err := client.NewClient(creds, conf)
    if err != nil {
        log.Fatal(err)
    }

    //osclib.GetVersions(c.Auth)

    if options.Verbose == true {
        fmt.Printf("Token.Id=%s\n", c.Token.Id)
        fmt.Printf("Token.Expires=%s\n", c.Token.Expires)
        fmt.Printf("Tenant.Id=%s\n", c.Token.Tenant.Id)
        fmt.Printf("Tenant.Name=%s\n", c.Token.Tenant.Name)
    } else {
        fmt.Printf("%s %s\n", c.Token.Id, c.Token.Expires)
    }
}
