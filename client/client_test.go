// client_test.go - roscoe client package tests

package client

import (
    "os"
    "testing"

    "roscoe/osclib"
)

func TestNewClient(t *testing.T) {
    // Set env
    os.Setenv("OS_USERNAME", "ford")
    os.Setenv("OS_PASSWORD", "prefect")
    os.Setenv("OS_TENANT_NAME", "ccc-guide")
    os.Setenv("OS_AUTH_URL", "http://ccc.com:42")
    var creds osclib.Creds
    oscc, _ := NewClient(creds)
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
