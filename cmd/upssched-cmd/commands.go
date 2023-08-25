package main

import (
	"fmt"
	"meteo/internal/config"
	"meteo/internal/kit"
	"meteo/internal/log"
	"meteo/internal/utils"
	"time"

	"github.com/spf13/cobra"
)

/*const (
	logFile  = "upssched"
	logLevel = "info"
)*/

func commbad(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	nutLogging("UPS communications failure")
}

func commok(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	nutLogging("UPS communications restored")
}

func nocomm(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	nutLogging("UPS communications cannot be established")
}

func powerout(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()

	command := fmt.Sprintf("upscmd -u %s -p %s %s@localhost:%d shutdown.return",
		config.Default.Nut.ApiUser,
		config.Default.Nut.ApiPass,
		config.Default.Nut.Driver,
		config.Default.Nut.Port)

	shell := utils.NewShell(command)
	err, _, _ := shell.Run(1)
	if err != nil {
		log.Errorf("upscmd error: %w", err)
		return
	}

	_, err = kit.PostInt("/messanger/telegram", fmt.Sprintf("Обесточен %v", time.Now().Format("Jan 02, 2006 15:04:05")))
	if err != nil {
		log.Errorf("can't send telegram message: %w", err)
	}

	nutLogging("UPS on battery. Shutdown in 20 minuts....")
}

func shutdownnow(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()

	time.Sleep(10 * time.Second)
	command := "upsmon -c fsd"
	shell := utils.NewShell(command)
	err, _, _ := shell.Run(1)
	if err != nil {
		log.Errorf("upsmon error: %w", err)
		return
	}
	_, err = kit.PostInt("/messanger/telegram", fmt.Sprintf("Выключен %v", time.Now().Format("Jan 02, 2006 15:04:05")))
	if err != nil {
		log.Errorf("can't send telegram message: %w", err)
	}

	nutLogging("UPS has been on battery. Starting orderly shutdown")
}

func shutdowncritical(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()

	command := "upsmon -c fsd"
	shell := utils.NewShell(command)
	err, _, _ := shell.Run(1)
	if err != nil {
		log.Errorf("upsmon error: %w", err)
		return
	}

	nutLogging("UPS battery level CRITICAL. Shutting down NOW!!!!")
}

func powerup(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()

	nutLogging("UPS on line. Shutdown aborted.")

	_, err := kit.PostInt("/messanger/telegram", fmt.Sprintf("Напряжение восстановлено %v", time.Now().Format("Jan 02, 2006 15:04:05")))
	if err != nil {
		log.Errorf("can't send telegram message: %w", err)
	}
}

func replbatt(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()

	nutLogging(fmt.Sprintf("UPS %s battery needs to be replaced", config.Default.Nut.Driver))

	_, err := kit.PostInt("/messanger/telegram", fmt.Sprintf("Необходимо заменить батарею ИБП %v", time.Now().Format("Jan 02, 2006 15:04:05")))
	if err != nil {
		log.Errorf("can't send telegram message: %w", err)
	}
}

func noparent(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	nutLogging("upsmon parent process died - shutdown impossible")
}

func nutLogging(message string) {
	_, err := kit.PostInt("/nut/message", message)
	if err != nil {
		log.Errorf("can't send telegram message: %w", err)
	}
}
