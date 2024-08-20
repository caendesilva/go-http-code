package main

import (
    "fmt"
    "os"
    "strconv"
)

// Map of HTTP status codes to descriptions
var httpStatusDescriptions = map[int]string{
    200: "OK - The request has succeeded.",
    201: "Created - The request has been fulfilled and has resulted in a new resource being created.",
    400: "Bad Request - The server could not understand the request due to invalid syntax.",
    401: "Unauthorized - The client must authenticate itself to get the requested response.",
    403: "Forbidden - The client does not have access rights to the content.",
    404: "Not Found - The server can not find the requested resource.",
    500: "Internal Server Error - The server has encountered a situation it doesn't know how to handle.",
    502: "Bad Gateway - The server, while acting as a gateway or proxy, received an invalid response from the upstream server.",
    503: "Service Unavailable - The server is not ready to handle the request.",
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: http-code <status_code>")
        os.Exit(1)
    }

    // Convert the first argument to an integer
    code, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Printf("Invalid status code: %s\n", os.Args[1])
        os.Exit(1)
    }

    // Look up the status code in the map
    description, exists := httpStatusDescriptions[code]
    if !exists {
        fmt.Printf("Unknown status code: %d\n", code)
        os.Exit(1)
    }

    // Print the description
    fmt.Printf("%d %s\n", code, description)
}
