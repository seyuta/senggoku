package main

import (
	"github.com/seyuta/senggoku/pkg/bootstrap"
	"github.com/seyuta/senggoku/pkg/constant"
)

func init() {
	// zerolog configuration
	bootstrap.InitZerolog()

	// reading application configuration
	cfg := bootstrap.NewKoanf(".")
	cfg.UseYamlFile("configs/local.yaml", true)
}

func main() {
	f := bootstrap.NewFiber()
	f.Middleware()
	f.Start(
		bootstrap.KoanfYamlFile.String(constant.AppHttpHost),
		bootstrap.KoanfYamlFile.String(constant.AppHttpPort),
	)
}
