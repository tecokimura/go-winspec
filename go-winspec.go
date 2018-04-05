package main

import (
    "fmt"
    "os/exec"
    "strings"
)

type Spec struct {
    pcname string
    system string
}


func main() {
    fmt.Printf("Hello go-winspec\n")
    fmt.Printf("----------------\n")
    
    pcname, err := exec.Command("cmd", "/C", "echo", "%COMPUTERNAME%").Output()
    if pcname != nil {
        fmt.Printf("PCName=%s", pcname)
    }
    
    if err != nil {
        fmt.Printf("err=%s", err)
    }

/*
    systeminfo, err := exec.Command("cmd", "/C", "systeminfo").Output()
    if systeminfo != nil {
        fmt.Printf("systeminfo=%s", systeminfo)
    }
    
    if err != nil {
        fmt.Printf("err=%s", err)
    }
*/
    applist, err := exec.Command("cmd", "/C", "reg", "query", "HKEY_LOCAL_MACHINE\\Software\\Microsoft\\Windows\\CurrentVersion\\Uninstall", "/s" ).Output()
    if applist != nil {
        // fmt.Printf("application=%s", applist)
        
        split := strings.Split(string(applist), "\n")
        for key, value := range split {
            fmt.Printf("%s=%s\n", key, value)
        }

    }
    
    if err != nil {
        fmt.Printf("err=%s", err)
    }

}

