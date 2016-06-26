package main

import (
    "time"
    "os/exec"
    "log"
)

func main(){
    for {
        time.Sleep(2* 60 * 60 * 1000 * time.Millisecond)
        cmd := exec.Command("durtucu")
        err := cmd.Start()
        if err != nil {
           log.Fatal(err)
        }
    }
}
