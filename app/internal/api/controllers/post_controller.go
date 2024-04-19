package controllers

import (
    "net/http"
    "log"
    "encoding/json"
)

type Payload struct{
    Command string `json:"command"`
    Data []string `json:"data"`
}

type Response struct{
    Status int `json:"status"`
    Message string `json:"message"`
    Data []int `json:"data"`
}

func decodePayload(req *http.Request) Payload{
    var data Payload
    decoder := json.NewDecoder(req.Body)
    err := decoder.Decode(&data)

    if err != nil {

        log.Println("Data parsed badly")
        log.Println(err)
        p := Payload{
            Command: "",
            Data: make([]string, 1),
        }
        return p
    }

    log.Println("Data parsed successfully")

    return data
}

func PostHandler(w http.ResponseWriter, r *http.Request){
    payload := decodePayload(r)
    if (payload.Command == "submit_word"){
        SubmitWord(&payload.Data)
    }
}

func SubmitWord(data *[]string){
    
}






