package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

    port := os.Getenv("PORT")
 
	if port == "" {
		port = "8080"
	}
 
	// http.HandleFunc("/", HelloHandler)
 
	// log.Println("Listening on port", port)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
    
    
    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
    })
    
    e.Logger.Fatal(e.Start(":" + port))
}
