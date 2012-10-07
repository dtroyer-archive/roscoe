// token.go - roscoe token example

package main

import (
    "fmt"
    "log"

    goopt "github.com/droundy/goopt"

    "roscoe/client"
    "roscoe/osclib"
)


// Declare options

var debug = goopt.Flag(
    []string{"-x", "--debug"},
    []string{"--nodebug"},
    "Enable debug mode",
    "Disable debug mode",
)

var verbose = goopt.Flag(
    []string{"-v", "--verbose"},
    []string{"-q", "--quiet"},
    "output verbosely",
    "be quiet, instead",
)


func main() {
    goopt.Description = func() string {
        return "OpenStack client example"
    }
    goopt.Version = "1.0"
    goopt.Parse(nil)

    // Propagate debug setting to packages
    client.Debug = debug

    // Get auth values from the environment
    var creds osclib.Creds
    c, err := client.NewClient(creds)
    if err != nil {
        log.Fatal(err)
    }

    osclib.GetVersions(c.Auth)

    if *verbose == true {
        fmt.Printf("Token.Id=%s\n", c.Token.Id)
        fmt.Printf("Token.Expires=%s\n", c.Token.Expires)
        fmt.Printf("Tenant.Id=%s\n", c.Token.Tenant.Id)
        fmt.Printf("Tenant.Name=%s\n", c.Token.Tenant.Name)
    } else {
        fmt.Printf("%s %s\n", c.Token.Id, c.Token.Expires)
    }
}
