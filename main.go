package main

import (
	"fmt"
	"os/exec"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"os"
    "os/signal"
    "path/filepath"
    "time"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}


func main() {

	var userName, password, databaseName string
	var port int32

	fmt.Print("Enter username : ")
	_, err := fmt.Scan(&userName)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter password : ")
	_, err1 := fmt.Scan(&password)
	if err1 != nil {
		panic(err1)
	}

	fmt.Print("Enter database name : ")
	_, err2 := fmt.Scan(&databaseName)
	if err2 != nil {
		panic(err2)
	}

	fmt.Print("Enter port number : ")
	_, err3 := fmt.Scan(&port)
	if err3 != nil {
		panic(err2)
	}

	fmt.Println("==================================[ Thanks For Information ]==================================")
 
	log.Info("Create new cron")
	c := cron.New()
	c.AddFunc("@daily", func() { 
		log.Info("[Job 1] Every day job started\n") 
		getDunmp(userName, password, databaseName, port)
		log.Info("[Job 1] Every day job completed\n") 
	})

	// Start cron with one scheduled job
	log.Info("Start cron")

	go c.Start()
	printCronEntries(c.Entries())
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, os.Kill)
    <-sig
}

func printCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}

func getDunmp(userName, password, databaseName string, port int32) {
	path, err := os.Getwd()
	if err != nil {log.Println(err)}
	fileName := fmt.Sprintf("%v%v.sql", databaseName, time.Now().Unix())
	mysqlPath, _ := exec.LookPath("mysqldump")
	cmd := exec.Command(mysqlPath, "-h127.0.0.1", fmt.Sprintf("-P%v", port),
		fmt.Sprintf("-u%v", userName), fmt.Sprintf("-p%v", password), databaseName,
		"-r", filepath.Join(path, fileName))
	if err := cmd.Run(); err != nil {
		log.Info("Error ", err)
	}
}