// json/json_test.go

package json

import (
    "log"
    "testing"
    
    "roscoe/json"
)

var jt []byte = []byte(`
{
    "auth": {
        "passwordCredentials": {
            "username": "trillian",
            "password": "tea"
        },
        "tenantName": "ZZPluralZAlpha",
        "enabledFalse": "false",
        "enabledTrue": "true",
        "theAnswer": 42,
        "floater": 1.23
    }
}
`)

func TestJsonGetKey(t *testing.T) {
    log.Print("Test JSON GetKey()")

    js, err := json.NewJsonStruct(jt)
    if err != nil {
        log.Print(err)
    }

    auth := js.GetKey("")
    if auth != js {
        t.Fatalf("Expected '' key to return source pointer")
    }
    auth = js.GetKey("auth")
    if auth == js {
        t.Fatalf("Expected 'auth' key, not found")
    }
    authA := js.GetKey("auth").GetKey("tenantName")
    if authA == js {
        t.Fatalf("Expected 'auth'.'tenantName' key, not found")
    }
    authB := js.GetKey("auth", "tenantName")
    if authB == js {
        t.Fatalf("Expected 'auth'.'tenantName' key, not found")
    }
    if *authA != *authB {
        t.Fatalf("Expected chained and string list keys to match")
    }
}

func TestJsonBool(t *testing.T) {
    log.Print("Test JSON GetKey().Bool()")

    js, err := json.NewJsonStruct(jt)
    if err != nil {
        log.Print(err)
    }

    log.Print("Test JSON GetKey().Bool() == true")
    enabledT, err := js.GetKey("auth", "enabledTrue").Bool()
    if err != nil {
        t.Fatalf("json.GetKey.Bool(true) failed: %s", err)
    }
    if !enabledT {
        t.Fatalf("Expected enabled == true, got false")
    }

    log.Print("Test JSON GetKey().Bool() == false")
    enabledF, err := js.GetKey("auth", "enabledFalse").Bool()
    if err != nil {
        t.Fatalf("json.GetKey.Bool(false) failed: %s", err)
    }
    if enabledF {
        t.Fatalf("Expected enabled == false, got true")
    }
}

func TestJsonFloat64(t *testing.T) {
    log.Print("Test JSON GetKey().Float64()")

    js, err := json.NewJsonStruct(jt)
    if err != nil {
        log.Print(err)
    }

    f, err := js.GetKey("auth", "floater").Float64()
    if err != nil {
        t.Fatalf("json.GetKey.Float64() failed: %s", err)
    }
    if f != 1.23 {
        t.Fatalf("Expected 1.23 for float, got %f", f)
    }

}

func TestJsonInt64(t *testing.T) {
    log.Print("Test JSON GetKey().Int64()")

    js, err := json.NewJsonStruct(jt)
    if err != nil {
        log.Print(err)
    }

    i, err := js.GetKey("auth", "theAnswer").Int64()
    if err != nil {
        t.Fatalf("json.GetKey.Int64() failed: %s", err)
    }
    if i != 42 {
        t.Fatalf("Expected 42 for int, got %d", i)
    }

}

func TestJsonString(t *testing.T) {
    log.Print("Test JSON GetKey().String()")

    js, err := json.NewJsonStruct(jt)
    if err != nil {
        log.Print(err)
    }

    tenant, err := js.GetKey("auth", "tenantName").String()
    if err != nil {
        t.Fatalf("json.GetKey failed: %s", err)
    }
    if tenant != "ZZPluralZAlpha" {
        t.Fatalf("Expected 'ZZPluralZAlpha' for tenantName, got %s", tenant)
    }

}
