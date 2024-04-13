package api 

import (
    "log"
    "net/http"
    "app/internal/api/controllers"
)

func RegisterServices(router *http.ServeMux){
    // GET SERVICES
    router.HandleFunc("/", controllers.GamePageHandler) 
}

func ServerStart(){

    router := http.NewServeMux()
    RegisterServices(router)

    fs := http.FileServer(http.Dir("static/"))
    router.Handle("/static/", http.StripPrefix("/static/", fs))

    log.Println("Starting Server on port 8080")

    if err := http.ListenAndServe(":8080",router); err != nil {
        log.Println(err)
    }
}

