// client_test.go - roscoe client package tests

package client

import (
    "os"
    "testing"

    "roscoe/client"
)

var creds = client.Credentials {
    OSAuth: {
        PasswordCredentials {
            Username: "ford",
            Password: "prefect",
        },
        TenantName: "ccc-guide",
    },
    AuthUrl: "http://ccc.com:42",
}

func TestNewClient(t *testing.T) {
    // Set env
    os.Setenv("OS_USERNAME", "ford")
    os.Setenv("OS_PASSWORD", "prefect")
    os.Setenv("OS_TENANT_NAME", "ccc-guide")
    os.Setenv("OS_AUTH_URL", "http://ccc.com:42")
    oscc, _ := client.NewClient(creds, nil)
    if oscc.Auth.OSAuth.PasswordCredentials.Username != "ford" {
        t.Error("NewCLient didn't pick up OS_USERNAME from environment")
    }
    if oscc.Auth.OSAuth.PasswordCredentials.Password != "prefect" {
        t.Error("NewCLient didn't pick up OS_PASSWORD from environment")
    }
    if oscc.Auth.OSAuth.TenantName != "ccc-guide" {
        t.Error("NewCLient didn't pick up OS_TENANT_NAME from environment")
    }
    if oscc.Auth.AuthUrl != "http://ccc.com:42" {
        t.Error("NewCLient didn't pick up OS_AUTH_URL from environment")
    }
}
