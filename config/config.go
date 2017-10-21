package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func config(){
	a := viper.Viper{}
	fmt.Print(a)
}
