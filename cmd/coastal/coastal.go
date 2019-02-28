package main

import (
    "coastal/internal/env"
    "coastal/internal/runtime"
    "flag"
)

func main() {
    env.New()
    runtime.New()
    migrate()

    server := flag.String("server", "http", "server types:\"http\",\"rpc\";")
    flag.Parse()

    if *server == "rpc" {
        grpcServer()
    } else {
        httpServer()
    }
}
