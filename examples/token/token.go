// token.go - roscoe token example

package main

import (
    // "crypto/tls"
    // "crypto/x509"
    "fmt"
    "log"
    "os"

    "github.com/voxelbrain/goptions"

    "roscoe/auth"
    "roscoe/session"
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
    session.Debug = &options.Debug

    // // read cacert
    // cacert := os.Getenv("OS_CACERT")

    // pool := NewCertPool()
    // for j, root := range test.roots {
    //     ok := opts.Roots.AppendCertsFromPEM([]byte(root))
    //     if !ok {
    //         t.Errorf("#%d: failed to parse root #%d", i, j)
    //         return
    //     }
    // }

    // // Set up TLS config
    // conf := &tls.Config{InsecureSkipVerify: true}

    // c, err := client.NewClient(creds, conf)
    // if err != nil {
    //     log.Fatal(err)
    // }


    // Get auth values from the environment
    var creds auth.UserPassV2
    creds.AuthUrl = os.Getenv("OS_AUTH_URL")
    creds.OSAuth.ProjectName = os.Getenv("OS_TENANT_NAME")
    creds.OSAuth.PasswordCredentials.Username = os.Getenv("OS_USERNAME")
    creds.OSAuth.PasswordCredentials.Password = os.Getenv("OS_PASSWORD")
    auth_ref, err := creds.GetAuthRef()

    // Make a new client with these creds
    _, err = session.NewSession(creds.AuthUserPassV2, nil)
    if err != nil {
        log.Fatal(err)
    }

    if options.Verbose == true {
        fmt.Printf("Token_Id=%s\n", auth_ref.GetTokenId())
        fmt.Printf("Token_Expires=%s\n", auth_ref.Access.Token.Expires)
        fmt.Printf("Project_Id=%s\n", auth_ref.Access.Token.Project.Id)
        fmt.Printf("Project_Name=%s\n", auth_ref.Access.Token.Project.Name)
    } else {
        fmt.Println(auth_ref.GetTokenId())
        // fmt.Printf("%s %s\n", c.Token.Id, c.Token.Expires)
    }
}
