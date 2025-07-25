package main

import (
	"log"
	"time"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

func main() {
	// 初始化外设（必须执行）
	if _, err := host.Init(); err != nil {
		log.Fatalf("初始化失败: %v", err)
	}

	// 获取 GPIO 引脚（物理引脚7 = BCM GPIO4）
	pin := gpioreg.ByName("GPIO4")
	if pin == nil {
		log.Fatal("未找到 GPIO4 引脚")
	}
	defer pin.Halt() // 程序结束时释放资源

	// 配置引脚为输出模式，初始值设为低电平
	if err := pin.Out(gpio.Low); err != nil {
		log.Fatalf("配置引脚失败: %v", err)
	}

	// 让 LED 闪烁
	log.Println("开始闪烁 LED（物理引脚7），按 Ctrl+C 退出")
	for {
		pin.Out(gpio.High)
		time.Sleep(500 * time.Millisecond)
		pin.Out(gpio.Low)
		time.Sleep(500 * time.Millisecond)
	}
}
