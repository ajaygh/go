package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func ReadConfig(name string) {
	viper.SetConfigName(name)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error in reading")
		return
	}
	dev_name := viper.GetString("dev.name")
	dev_age := viper.GetInt("dev.age")
	dev_pin := viper.GetInt("dev.pincode")
	fmt.Printf("dev name = %s, age = %d, pin = %d\n", dev_name, dev_age, dev_pin)

}

func main() {
	ReadConfig("data")
}
