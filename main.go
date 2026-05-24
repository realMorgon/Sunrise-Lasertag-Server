package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"lasertag-server/data"
)

var devices = make(map[int]data.Device)
var players = make(map[int]data.Player)
var startTime = time.Now()

func main() {
	router := gin.Default()

	// Statische Dateien servieren (CSS, JS, etc.)
	router.Static("/static", "./static")

	// Landing Page (index.html)
	router.StaticFile("/", "./static/index.html")

	router.GET("/api/test", func(context *gin.Context) {
		fmt.Println("Anfrage erhalten!")
		context.String(200, "Anfrage erhalten!")
	})

	router.POST("/api/sync", func(c *gin.Context) {
		clienttime, _ := strconv.Atoi(c.Query("time"))
		clientid := len(devices)
		clienttype := c.Query("type")
		c.String(200, strconv.Itoa(clientid))
		fmt.Println("ID: " + strconv.Itoa(clientid) + " beginnt bei " + strconv.Itoa(clienttime) + " mills")

		saveDevice(clientid, clienttime, clienttype)
		fmt.Println(devices)
		if clienttype == "Vest" {
			players[clientid] = data.NewPlayer(clientid)
			fmt.Println(players)
		}
	})

	router.POST("/api/hit/send", func(c *gin.Context) {
		clientid, _ := strconv.Atoi(c.Query("id"))
		clienttime, _ := strconv.Atoi(c.Query("time"))
		c.String(200, " Du hast ID: "+strconv.Itoa(clientid)+" und TIME: "+strconv.Itoa(clienttime)+" (Clienttime) als Schuss gesendet")

		gametime := calculateTime(clientid, clienttime)
		fmt.Println("ID: " + strconv.Itoa(clientid) + " schießt bei " + strconv.Itoa(gametime) + " mills")
	})

	router.POST("/api/hit/recieve", func(c *gin.Context) {
		client_id, _ := strconv.Atoi(c.Query("my_id"))
		hitert_id, _ := strconv.Atoi(c.Query("their_id"))
		clienttime, _ := strconv.Atoi(c.Query("time"))
		c.String(200, " Du hast EIGENE_ID: "+strconv.Itoa(client_id)+", FREMDE_ID: "+strconv.Itoa(hitert_id)+" und TIME: "+strconv.Itoa(clienttime)+" (Clienttime) als Treffer gesendet")

		gametime := calculateTime(client_id, clienttime)
		fmt.Println("ID: " + strconv.Itoa(client_id) + " wurde von ID: " + strconv.Itoa(hitert_id) + " bei " + strconv.Itoa(gametime) + " mills getroffen")

		players[client_id].DealDamage(1)
	})

	router.Run("0.0.0.0:5000")
}

func millis() int {
	return int(time.Since(startTime).Milliseconds())
}

func saveDevice(clientid int, clienttime int, clienttype string) {
	devices[clientid] = data.NewDevice(clienttime, clienttype)
}

func calculateTime(clientid int, clienttime int) int {
	return devices[clientid].GetTime() + clienttime
}
