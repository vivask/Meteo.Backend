package main

import (
	"meteo/internal/config"
	"meteo/internal/log"
	"strings"

	"github.com/spf13/cobra"
)

var (
	configPath string

	commandCmd = &cobra.Command{
		Use:   "commbad",
		Short: "commbad upssched-cmd",
		Long:  `commbad upssched-cmd`,
		Run:   commbad,
	}

	commokCmd = &cobra.Command{
		Use:   "commok",
		Short: "commok upssched-cmd",
		Long:  `commok upssched-cmd`,
		Run:   commok,
	}

	nocommCmd = &cobra.Command{
		Use:   "nocomm",
		Short: "nocomm upssched-cmd",
		Long:  `nocomm upssched-cmd`,
		Run:   nocomm,
	}

	poweroutCmd = &cobra.Command{
		Use:   "powerout",
		Short: "powerout upssched-cmd",
		Long:  `powerout upssched-cmd`,
		Run:   powerout,
	}

	shutdownnowCmd = &cobra.Command{
		Use:   "shutdownnow",
		Short: "shutdownnow upssched-cmd",
		Long:  `shutdownnow upssched-cmd`,
		Run:   shutdownnow,
	}

	shutdowncriticalCmd = &cobra.Command{
		Use:   "shutdowncritical",
		Short: "shutdowncritical upssched-cmd",
		Long:  `shutdowncritical upssched-cmd`,
		Run:   shutdowncritical,
	}

	poweruplCmd = &cobra.Command{
		Use:   "powerup",
		Short: "powerup upssched-cmd",
		Long:  `powerup upssched-cmd`,
		Run:   powerup,
	}

	replbattCmd = &cobra.Command{
		Use:   "replbatt",
		Short: "replbatt upssched-cmd",
		Long:  `replbatt upssched-cmd`,
		Run:   replbatt,
	}

	noparentCmd = &cobra.Command{
		Use:   "noparent",
		Short: "noparent upssched-cmd",
		Long:  `noparent upssched-cmd`,
		Run:   noparent,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(commandCmd)
	rootCmd.AddCommand(commokCmd)
	rootCmd.AddCommand(nocommCmd)
	rootCmd.AddCommand(poweroutCmd)
	rootCmd.AddCommand(shutdownnowCmd)
	rootCmd.AddCommand(shutdowncriticalCmd)
	rootCmd.AddCommand(poweruplCmd)
	rootCmd.AddCommand(replbattCmd)
	rootCmd.AddCommand(noparentCmd)
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
