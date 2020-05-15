package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8000"
	url := "http://" + "localhost" + port + "/welcome"
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/welcome", func(c *gin.Context) {
		name := c.Query("name")
		time := c.Query("time")
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"name": name,
			"time": time,
		})
	})

	router.GET("/v1/welcome", func(c *gin.Context) {
		name := c.Query("name")
		time := c.Query("time")
		message := name + " & " + time
		c.String(http.StatusOK, message)
		openbrowser(url + "?name=" + name + "&time=" + time)
	})

	router.Run(port)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
