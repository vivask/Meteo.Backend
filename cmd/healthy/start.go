package main

import (
	"meteo/internal/config"
	"meteo/internal/log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	configPath string

	clusterCmd = &cobra.Command{
		Use:   "cluster",
		Short: "healthy cluster",
		Long:  `healthy cluster`,
		Run:   cluster,
	}

	esp32Cmd = &cobra.Command{
		Use:   "esp32",
		Short: "healthy esp32",
		Long:  `healthy esp32`,
		Run:   esp32,
	}

	mediaCmd = &cobra.Command{
		Use:   "media",
		Short: "healthy media",
		Long:  `healthy media`,
		Run:   media,
	}

	messangerCmd = &cobra.Command{
		Use:   "messanger",
		Short: "healthy messanger",
		Long:  `healthy messanger`,
		Run:   messanger,
	}

	nutCmd = &cobra.Command{
		Use:   "nut",
		Short: "healthy nut",
		Long:  `healthy nut`,
		Run:   nut,
	}

	proxyCmd = &cobra.Command{
		Use:   "proxy",
		Short: "healthy proxy",
		Long:  `healthy proxy`,
		Run:   proxy,
	}

	radiusCmd = &cobra.Command{
		Use:   "radius",
		Short: "healthy radius",
		Long:  `healthy radius`,
		Run:   radius,
	}

	scheduleCmd = &cobra.Command{
		Use:   "schedule",
		Short: "healthy schedule",
		Long:  `healthy schedule`,
		Run:   schedule,
	}

	sshclientCmd = &cobra.Command{
		Use:   "sshclient",
		Short: "healthy sshclient",
		Long:  `healthy sshclient`,
		Run:   sshclient,
	}

	webCmd = &cobra.Command{
		Use:   "web",
		Short: "healthy web",
		Long:  `healthy web`,
		Run:   web,
	}

	radioCmd = &cobra.Command{
		Use:   "radio",
		Short: "healthy radio",
		Long:  `healthy radio`,
		Run:   radio,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(clusterCmd)
	rootCmd.AddCommand(esp32Cmd)
	rootCmd.AddCommand(mediaCmd)
	rootCmd.AddCommand(messangerCmd)
	rootCmd.AddCommand(nutCmd)
	rootCmd.AddCommand(proxyCmd)
	rootCmd.AddCommand(radiusCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(sshclientCmd)
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(radioCmd)
	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "/run/secrets/config.yaml", "config file (default is $PWD/config/default.yaml)")
}

func initConfig() {
	if len(configPath) != 0 {
		config.Viper().SetConfigFile(configPath)
	} else {
		config.Viper().AddConfigPath("./config")
		config.Viper().SetConfigName("default")
	}
	config.Viper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.Viper().AutomaticEnv()
	if err := config.Viper().ReadInConfig(); err != nil {
		log.Fatalf("Load config from file [%s]: %v", config.Viper().ConfigFileUsed(), err)
	}
	config.Parse()
}
