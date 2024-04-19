package utils

import(
    "fmt"
    "log"
    "net/http"
    "crypto/rand"
)


func CreateUuid() string{
    uuid := make([]byte, 16)

    _, err := rand.Read(uuid)
    if err != nil {
        panic(err)
    }

    // Set the version (4) and variant (2) bits
    // sets the 7th byte to version 4 as per UUID v4 RFC
    // sets the 9th byte to the variants as per the UUID RFC
    // & and | are bitwise ops 
    uuid[6] = (uuid[6] & 0x0f) | 0x40
    uuid[8] = (uuid[8] & 0x3f) | 0x80

    //format bytes to return string
    str_uuid := fmt.Sprintf("%x", uuid)

    return str_uuid; 
}

func ValidateCookie(cookie *http.Cookie) bool{
    return true
}

func CreateDataCookie() http.Cookie {
    // generate uuid
    uuid := CreateUuid()
    log.Println(uuid)

    // create cookie with id
    cookie := http.Cookie{
        Name:     "SESSION",
        Value:    uuid,
        Path:     "/",
        MaxAge:   3600,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
    }

    return cookie
}
