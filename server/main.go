package main

import (
    "github.com/YuyaYoshioka/todo-app-with-go-and-react/server/db"
    "github.com/YuyaYoshioka/todo-app-with-go-and-react/server/server"
)

func main() {
    db.Init()
    server.Init()
    defer db.Close()
}
