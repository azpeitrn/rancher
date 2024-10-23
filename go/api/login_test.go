package api

import (
    "bytes"
    "crypto/tls"
    "net/http"
    "testing"
)

func TestLogin(t *testing.T) {
    url := "https://44.214.93.44/v3-public/localProviders/local?action=login"
    method := "POST"

    payload := []byte(`{"description":"UI session","responseType":"cookie","username":"rancher-standard","password":"WaczB8tSX0BppPyyYAtU"}`)

    // Create a custom HTTP client with insecure transport
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}

    req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
    if err != nil {
        t.Fatalf("Error creating request: %v", err)
    }

    req.Header.Add("Accept", "application/json")
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Authorization", "Bearer token-j79bk:j97wmb4nm54v88ff79lqlb2rhl9n97xvhl2zlpp2dw6vrpfznjjbxf")
    req.Header.Add("x-api-key", "token-j79bk")

    resp, err := client.Do(req)
    if err != nil {
        t.Fatalf("Error making request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
    }
}
