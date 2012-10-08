// token.go - roscoe token example

package main

import (
    "fmt"
    "log"

    "github.com/voxelbrain/goptions"

    "roscoe/client"
    "roscoe/osclib"
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
    var creds osclib.Creds
    c, err := client.NewClient(creds)
    if err != nil {
        log.Fatal(err)
    }

    osclib.GetVersions(c.Auth)

    if options.Verbose == true {
        fmt.Printf("Token.Id=%s\n", c.Token.Id)
        fmt.Printf("Token.Expires=%s\n", c.Token.Expires)
        fmt.Printf("Tenant.Id=%s\n", c.Token.Tenant.Id)
        fmt.Printf("Tenant.Name=%s\n", c.Token.Tenant.Name)
    } else {
        fmt.Printf("%s %s\n", c.Token.Id, c.Token.Expires)
    }
}
