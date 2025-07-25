package main

import (
	"fmt"
	"log"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/host/v3/rpi"
)

import "periph.io/x/host/v3"

func main() {
	//初始化引脚
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
		fmt.Println(err)
		return
	}
	//第12物理引脚
	pin := rpi.P1_12
	//输出高电平
	if err := pin.Out(gpio.High); err != nil {
		fmt.Println("设置GPIO引脚失败", err)
		return
	}

	fmt.Println("GPIO引脚设置成功")

}
