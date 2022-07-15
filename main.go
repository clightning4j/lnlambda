package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/clightning4j/lnsocket-function/lnsocket"
	"github.com/gin-gonic/gin"
)

type Request struct {
	NodeID  string         `json:"node_id"`
	Address string         `json:"host"`
	Rune    string         `json:"rune"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
}

func PostMethod(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
		return
	}
	var req Request
	if err := json.Unmarshal(jsonData, &req); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
		return
	}
	client := lnsocket.New(req.NodeID, req.Address)
	if err := client.Connect(); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
		return
	}
	response, err := client.Call(req.Method, req.Params, req.Rune)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err))
		return
	}
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.POST("/lnsocket", PostMethod)

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	listenPort := "9002"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)
}
