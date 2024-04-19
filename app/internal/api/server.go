package api 

import (
    "log"
    "net/http"
    "app/internal/api/controllers"
    "app/internal/api/middleware"
)

func RegisterGetServices(router *http.ServeMux){
    finalHandler := http.HandlerFunc(controllers.GamePageHandler)
    router.Handle("/", middleware.PreProcess(finalHandler)) 
}

func RegisterPostServices(router *http.ServeMux){
    router.HandleFunc("/api/post", controllers.PostHandler) 
}

func ServerStart(){

    router := http.NewServeMux()

    fs := http.FileServer(http.Dir("static/"))
    router.Handle("/static/", http.StripPrefix("/static/", fs))

    RegisterGetServices(router)
    RegisterPostServices(router)

    log.Println("Starting Server on port 8080")

    if err := http.ListenAndServe(":8080",router); err != nil {
        log.Println(err)
    }
}

