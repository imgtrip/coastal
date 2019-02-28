package main

import (
    "flag"
    "fmt"
)

func main() {
    server := flag.String("server", "http", "server types:\"http\",\"rpc\";")
    flag.Parse()

    fmt.Println(*server)
}
