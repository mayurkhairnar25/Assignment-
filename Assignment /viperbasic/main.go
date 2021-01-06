package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("in main")
	viper.SetConfigName("config") // name of config file (without extension)
}
