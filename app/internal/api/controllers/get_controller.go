package controllers

import (
    "net/http"
    "html/template"
)

func GamePageHandler(w http.ResponseWriter, r *http.Request){

   template, err := template.ParseFiles("internal/views/game/game.html") 

   if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
   }
    
   template.Execute(w, nil)
}
