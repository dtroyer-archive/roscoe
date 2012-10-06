// client/doc.go - roscoe server package documentation

/*

'client' package implements a low-level intoerface for OpenStack APIs

Example:

    resp, err := c.Get("compute", "/servers")
    var s ServerResponse
    err = json.Unmarshal(resp.Body, &s)

*/
package client
