package main 

import (
    "app/internal/api"
    "app/internal/api/globals"
)

func main(){
    globals.Init()
    api.ServerStart()
}

