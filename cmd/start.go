package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/lendlord/lendlord-server/configs"
	"github.com/lendlord/lendlord-server/internal/app"
)

var (
	localConfig string
	startCmd    = &cobra.Command{
		Use:   "start",
		Short: "start lendlord-server",
		Run: func(cmd *cobra.Command, args []string) {
			online()
		},
	}
)

func init() {
	startCmd.Flags().StringVarP(&localConfig, "config", "c", "", "config path: /opt/local.toml")
	rootCmd.AddCommand(startCmd)
}

func online() {
	v := viper.New()
	// 配置文件名称
	v.AddConfigPath(localConfig)
	v.SetConfigName("config")
	v.SetConfigType("toml")
	// 读取配置文件内容
	if err := v.ReadInConfig(); err != nil {
		log.Errorf("read config err: %s", err.Error())
		return
	}
	var config configs.Config
	// 反序列化
	if err := v.Unmarshal(&config); err != nil {
		log.Errorf("unmarshal config err: %s", err.Error())
		return
	}
	fmt.Println("++++++++++read config success", localConfig)
	app.Server(&config)
}
