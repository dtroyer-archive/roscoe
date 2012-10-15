// json/json.go

/*

JSON library to handle dynamic structures

*/

package json

import (
    "encoding/json"
    "errors"
)

// Set up the error returns
var (
    ErrArrayFailed = errors.New("Error converting to array type")
    ErrBoolFailed = errors.New("Error converting to bool type")
    ErrFloat64Failed = errors.New("Error converting to float64 type")
    ErrInt64Failed = errors.New("Error converting to int64 type")
    ErrMapFailed = errors.New("Error converting to map type")
    ErrStringFailed = errors.New("Error converting to string type")
)

// A generic struct to hold the JSON
type JsonStruct struct {
    raw interface{}
}

// Create a new JsonStruct object from the raw JSON input
func NewJsonStruct(jtext []byte) (js *JsonStruct, err error) {
    js = new(JsonStruct)
    err = json.Unmarshal(jtext, &js.raw)
    if err != nil {
        return nil, err
    }
    return js, nil
}

// Return an ``array`` type
func (js *JsonStruct) Array() ([]interface{}, error) {
    if a, ok := (js.raw).([]interface{}); ok {
        return a, nil
    }
    return nil, ErrArrayFailed
}

// Return a ``bool`` type
func (js *JsonStruct) Bool() (bool, error) {
    if b, ok := (js.raw).(string); ok {
        if b == "true" {
            return true, nil
        }
        if b == "false" {
            return false, nil
        }
    }
    return false, ErrBoolFailed
}

// Return a ``float64`` type
func (js *JsonStruct) Float64() (float64, error) {
    if i, ok := (js.raw).(float64); ok {
        return i, nil
    }
    return -1, ErrFloat64Failed
}

// Return a pointer to the JsonStruct indexed by key ``k0 [, k1 ...]``
// Can be chained: js.GetKey("a").GetKey("b")
// or given a list of keys: js.GetKey("a", "b")
func (js *JsonStruct) GetKey(k ...string) (v *JsonStruct) {
    var r *JsonStruct = js
    // Loop through key list
    for _, y := range k {
        // Make a map for each key
        m, err := r.Map()
        if err != nil {
            // Map conversion failed...bad JSON?
            return &JsonStruct{nil}
        }
        if v, ok := m[y]; ok {
            r = &JsonStruct{v}
        }
    }
    return r
}

// Return an ``int64`` type
func (js *JsonStruct) Int64() (int64, error) {
    if f, ok := (js.raw).(float64); ok {
        return int64(f), nil
    }
    return -1, ErrInt64Failed
}

// Return a ``map`` type
func (js *JsonStruct) Map() (map[string]interface{}, error) {
    if m, ok := (js.raw).(map[string]interface{}); ok {
        return m, nil
    }
    return nil, ErrMapFailed
}

// Marshal the JsonStruct back into []byte
func (js *JsonStruct) Marshal() ([]byte, error) {
        return json.Marshal(&js.raw)
}

// Return a ``string`` typw
func (js *JsonStruct) String() (string, error) {
    if s, ok := (js.raw).(string); ok {
        return s, nil
    }
    return "", ErrStringFailed
}
