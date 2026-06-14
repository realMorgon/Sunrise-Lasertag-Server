package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/realMorgon/Sunrise-Lasertag-Server/data"
)

var devices = make(map[int]data.Device)
var players = make(map[int]data.Player)
var startTime = time.Now().UnixMilli()

func main() {

	fmt.Println("Beep boop beep beep boop...")
	fmt.Println("Server started at " + time.Now().String() + " which is equal to " + strconv.FormatInt(startTime, 10) + " mills")

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
		time, _ := strconv.Atoi(c.Query("time"))
		clientTime := int64(time)
		clientId := len(devices)
		clientType := c.Query("type")
		c.String(200, strconv.Itoa(clientId))
		fmt.Println("ID: " + strconv.Itoa(clientId) + " beginnt bei " + strconv.FormatInt(clientTime, 10) + " mills")

		saveDevice(clientId, clientTime, clientType)
		fmt.Println(devices)
		if clientType == "Vest" {
			players[clientId] = data.NewPlayer(clientId)
			fmt.Println(players)
		}
	})

	router.POST("/api/hit/send", func(c *gin.Context) {
		clientId, _ := strconv.Atoi(c.Query("id"))
		time, _ := strconv.Atoi(c.Query("time"))
		clientTime := int64(time)
		c.String(200, "Du hast ID: "+strconv.Itoa(clientId)+" und TIME: "+strconv.Itoa(int(clientTime))+" (Clienttime) als Schuss gesendet")

		servertime := calculateTime(clientId, clientTime)
		fmt.Println("ID: " + strconv.Itoa(clientId) + " schießt bei " + strconv.FormatInt(servertime, 10) + " mills")
	})

	router.POST("/api/hit/recieve", func(c *gin.Context) {
		clientId, _ := strconv.Atoi(c.Query("my_id"))
		hitertId, _ := strconv.Atoi(c.Query("their_id"))
		time, _ := strconv.Atoi(c.Query("time"))
		clientTime := int64(time)
		c.String(200, "Du hast EIGENE_ID: "+strconv.Itoa(clientId)+", FREMDE_ID: "+strconv.Itoa(hitertId)+" und TIME: "+strconv.Itoa(int(clientTime))+" (Clienttime) als Treffer gesendet")

		servertime := calculateTime(clientId, clientTime)
		fmt.Println("ID: " + strconv.Itoa(clientId) + " wurde von ID: " + strconv.Itoa(hitertId) + " bei " + strconv.FormatInt(servertime, 10) + " mills getroffen")

		players[clientId].DealDamage(1)
	})

	router.Run("0.0.0.0:5000")
}

func saveDevice(clientid int, client_time int64, clienttype string) {
	devices[clientid] = data.NewDevice(client_time, clienttype)
}

func calculateTime(clientid int, clienttime int64) int64 {
	return startTime + clienttime - devices[clientid].GetTimeDifference()
}
