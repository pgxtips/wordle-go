package middleware

import(
    "fmt"
    "net/http"
    "app/internal/utils"
    "app/internal/api/globals"
)

func PreProcess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            
            cookie, err := r.Cookie("SESSION")

            if err != nil {
                fmt.Println("Could not find cookie, creating a new cookie")
                cookie := utils.CreateDataCookie()
                globals.AddServerData(cookie.Value) 
                http.SetCookie(w, &cookie)
            } else {
                fmt.Println("cookie found")
                if (!utils.ValidateCookie(cookie)) {
                    fmt.Println("Invalid Cookie, Removing Cookie from server data")
                    globals.RemoveServerData(cookie.Value)

                    fmt.Println("Due to invalid cookie, creating new cookie")
                    cookie := utils.CreateDataCookie()
                    globals.AddServerData(cookie.Value) 
                    http.SetCookie(w, &cookie)
                }
            } 

        } 
        next.ServeHTTP(w, r)
    })
}
