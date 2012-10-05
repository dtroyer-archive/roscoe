// osc.go - OpenStackClient

package main

import (
    "fmt"
    "log"

    "roscoe/client"
    "roscoe/osclib"
)

func main() {
    // Get auth values from the environment
    var creds osclib.Creds
    c, err := client.NewClient(creds)
    if err != nil {
        log.Fatal(err)
    }

    osclib.GetVersions(c.Auth)

    c.Authenticate()
    fmt.Printf("token: %s\n", c.Token)
    fmt.Printf("servcat: %s\n", c.ServCat)
}
