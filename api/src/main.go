
package main

import (
    "flag"
    "app/config"
    "app/database"
    "app/server"
    "app/models"
)

func main() {

    env := flag.String("e", "development", "")
    flag.Parse()

    config.Init(*env)
    database.Init(true, models.GetModels()...)
    defer database.Close()
    if err := server.Init(); err != nil {
        panic(err)
    }
}
