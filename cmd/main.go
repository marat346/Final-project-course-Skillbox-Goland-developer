package main

import (
	"flag"
	"log"

	"github.com/marat346/Final-project-course-Skillbox-Goland-developer/tree/master/include/server"
	"github.com/marat346/Final-project-course-Skillbox-Goland-developer/tree/master/pkg/simulator"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Println("Configuration file not loaded")
	}

	mode := flag.String("mode", "server", "a string")
	addr := flag.String("addr", viper.GetString("server.addr"), "a string")
	port := flag.String("port", viper.GetString("server.port"), "a string")
	flag.Parse()

	simUrl := viper.GetString("simulator.addr") + ":" + viper.GetString("simulator.port")
	url := *addr + ":" + *port

	switch {
	case *mode == "server":
		server.StartServer(url)
	case *mode == "simulator":
		simulator.Start(simUrl)
	}
}

// initConfig() - инициализирует файл конфигурации в формате yaml
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
