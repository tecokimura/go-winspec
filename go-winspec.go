package main

import (
    "fmt"
    "os/exec"
)

func main() {
    fmt.Printf("Hello go-winspec\n")
    fmt.Printf("----------------\n")
    
    out, err := exec.Command("ls", "-la").Output()
    if out != nil {
        fmt.Printf("out=%s", out)
    }
    
    if out != nil {
        fmt.Printf("err=%s", err)
    }
}

