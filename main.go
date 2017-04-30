package main

import (
    "os"
    "os/signal"
    "syscall"
)

func main() {
    New().Start()

    // keep bots running until force closed
    sigChan := make(chan os.Signal)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
}
