package main

import (
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/infiniteloopsco/guartz/endpoint"
	"github.com/infiniteloopsco/guartz/models"
	"gopkg.in/robfig/cron.v2"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	argsCad1 := "-X POST --data payload={\"channel\":\"#general\",\"text\":\"EOOO\"} https://hooks.slack.com/services/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"
	args1 := strings.Split(argsCad1, " ") //[]string{"-i", "tcp:3000"}

	argsCad2 := "-X POST --data payload={\"channel\":\"#general\",\"text\":\"OTOOOOO\"} https://hooks.slack.com/services/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"
	args2 := strings.Split(argsCad2, " ") //[]string{"-i", "tcp:3000"}

	c := cron.New()
	id, _ := c.AddFunc("@every 10s", func() {
		exec.Command("curl", args1...).Output()
	})
	c.Start()
	c.Remove(154)
	time.Sleep(1 * time.Minute)
	c.Remove(id)
	c.AddFunc("@every 10s", func() {
		exec.Command("curl", args2...).Output()
	})
	time.Sleep(1 * time.Minute)
	c.Stop()
}

func mai1() {
	initEnv()
	models.InitDB()
	models.InitCron()

	router := endpoint.GetMainEngine()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}

func initEnv() {
	if err := godotenv.Load(".env_dev"); err != nil {
		log.Fatal("Error loading .env_dev file")
	}
}
