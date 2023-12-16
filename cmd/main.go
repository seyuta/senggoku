package main

import (
	"github.com/seyuta/senggoku/pkg/bootstrap"
)

func init() {
	// zerolog configuration
	bootstrap.InitZerolog()
}

func main() {
	f := bootstrap.NewFiber()
	f.Middleware()
	f.Start("", "3000")
}
