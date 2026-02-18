package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var timesyncs = make(map[int]int)
var startTime = time.Now()

func main() {
	router := gin.Default()

	router.GET("/api/test", func(context *gin.Context) {
		fmt.Println("Anfrage erhalten!")
		context.String(200, "Anfrage erhalten!")
	})

	router.POST("/api/sync", func(c *gin.Context) {
		clienttime, _ := strconv.Atoi(c.Query("time"))
		clientid := len(timesyncs)
		c.String(200, strconv.Itoa(clientid))
		fmt.Println("ID: " + strconv.Itoa(clientid) + " beginnt bei " + strconv.Itoa(clienttime) + " mills")

		saveTimesync(clientid, clienttime)
		fmt.Println(timesyncs)
	})

	router.POST("/api/hit/send", func(c *gin.Context) {
		clientid, _ := strconv.Atoi(c.Query("id"))
		clienttime, _ := strconv.Atoi(c.Query("time"))
		c.String(200, " Du hast ID: "+strconv.Itoa(clientid)+" und TIME: "+strconv.Itoa(clienttime)+" (Clienttime) als Schuss gesendet")

		gametime := calculateTime(clientid, clienttime)
		fmt.Println("ID: " + strconv.Itoa(clientid) + " schießt bei " + strconv.Itoa(gametime) + " mills")
	})

	router.POST("/api/hit/recieve", func(c *gin.Context) {
		clientid, _ := strconv.Atoi(c.Query("my_id"))
		hitertid, _ := strconv.Atoi(c.Query("their_id"))
		clienttime, _ := strconv.Atoi(c.Query("time"))
		c.String(200, " Du hats ID: "+strconv.Itoa(clientid)+" und TIME: "+strconv.Itoa(clienttime)+" (Clienttime) als Treffer gesendet")

		gametime := calculateTime(clientid, clienttime)
		fmt.Println("ID: " + strconv.Itoa(clientid) + " wurde von ID: " + strconv.Itoa(hitertid) + " bei " + strconv.Itoa(gametime) + " mills getroffen")
	})

	router.Run("0.0.0.0:5000")
}

func millis() int {
	return int(time.Since(startTime).Milliseconds())
}

func saveTimesync(clientid int, clienttime int) {
	timesyncs[clientid] = millis() - clienttime
}

func calculateTime(clientid int, clienttime int) int {
	return timesyncs[clientid] + clienttime
}
