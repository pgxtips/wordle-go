package api 

import (
    "log"
    "net/http"
    "app/internal/api/controllers"
)

func RegisterGetServices(router *http.ServeMux){
    router.HandleFunc("/", controllers.GamePageHandler) 
}

func RegisterPostServices(router *http.ServeMux){
}

func ServerStart(){

    router := http.NewServeMux()
    RegisterGetServices(router)

    fs := http.FileServer(http.Dir("static/"))
    router.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Starting Server on port 8080")

    if err := http.ListenAndServe(":8080",router); err != nil {
        log.Println(err)
    }
}

