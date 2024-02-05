package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main () {
    var path string
    args := os.Args

    if len(args) > 1 {
        path = args[1]
    } else {
        fmt.Fprintf(os.Stderr, "No path passed")
        os.Exit(1)
    }

    fileData, err := os.ReadFile(path)
    if err != nil {
        fmt.Fprintf(os.Stderr, "No valid file path")
        os.Exit(1)
    }

    var parsedJson map[string]interface{}
    if err := json.Unmarshal(fileData, &parsedJson); err != nil {
        fmt.Fprintf(os.Stderr, err.Error())
        os.Exit(1)
    }
}

