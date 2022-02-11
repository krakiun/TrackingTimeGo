package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    "time"
_   "log"

    "github.com/lextoumbourou/idle"
)

var(
    TimeWork = 0
    TotalTime = 0
)

const LOGO = `
___________                     __   ___________.__     2022-02   
\__    ___/___________    ____ |  | _\__    ___/|__| _____   ____  
  |    |  \_  __ \__  \ _/ ___\|  |/ / |    |   |  |/     \_/ __ \ 
  |    |   |  | \// __ \\  \___|    <  |    |   |  |  Y Y  \  ___/ 
  |____|   |__|  (____  /\___  >__|_ \ |____|   |__|__|_|  /\___  >
       WORK SMART     \/     \/     \/                   \/     \/ 
`

func cleanup() {
    TotalTime += TimeWork
    pingBackend(TimeWork)
    fmt.Println("----- STOPED -----")
}

func pingBackend(timeWork int) {
    fmt.Printf("+ %d secconds work (total : %d)\n", timeWork, TotalTime)
}

func main() {
    var err error
    var idleTime time.Duration

    fmt.Println(LOGO)

    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        cleanup()
        os.Exit(1)
    }()

    fmt.Println("----- STARTED -----")


    for {
        idleTime, err = idle.Get()

        if idleTime.Seconds() < 30.0 {
            TimeWork = TimeWork + 1 
        }

        if TimeWork >= 300 {
            go pingBackend(TimeWork)
            TotalTime += TimeWork
            TimeWork = 0
        }

        time.Sleep(1 * time.Second)
    }

    _ = err
}