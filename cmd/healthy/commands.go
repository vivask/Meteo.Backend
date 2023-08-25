package main

import (
	"fmt"
	"meteo/internal/kit"
	"os"

	"github.com/spf13/cobra"
)

const (
	//logFile  = "healthy"
	//logLevel = "info"
	response = "de030a3d4e727eb99842"
)

func cluster(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/cluster/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func esp32(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/esp32/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func media(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/media/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func messanger(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/messanger/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func nut(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/nut/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func proxy(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/proxy/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func radius(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/radius/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func schedule(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/schedule/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func sshclient(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/sshclient/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func web(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/web/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}

func radio(cmd *cobra.Command, agrs []string) {
	//log.SetLogger(logFile, logLevel)
	kit.InitClient()
	if kit.IsHealthyInt("/radio/health") {
		fmt.Println(response)
	} else {
		os.Exit(1)
	}
}
