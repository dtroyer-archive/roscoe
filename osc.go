// osc.go - OpenStackClient

package main

import (
    "fmt"
    "log"

    "./osclib"
)

func main() {
    // Get auth values from the environment
    var auth osclib.Creds
    err := auth.GetEnv()
    if err != nil {
        log.Fatal(err)
    }

    osclib.GetVersions(auth)

    token, sc, err := osclib.GetToken(auth)
    _ = token
    fmt.Printf("token: %+v\n\n", token)
    fmt.Printf("sc: %+v\n\n", sc)
//    fmt.Printf("volume admin: %+v\n", sc[0].Endpoints[0]["adminURL"])
}
