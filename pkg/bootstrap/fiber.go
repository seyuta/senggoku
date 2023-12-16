package bootstrap

import (
	"fmt"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

type Fiber struct {
	App *fiber.App
}

func NewFiber(config ...fiber.Config) *Fiber {
	var cfg fiber.Config
	if len(config) > 0 {
		cfg = config[0]
	}

	return &Fiber{
		App: fiber.New(cfg),
	}
}

func (f *Fiber) Middleware() {
	// Initialize zerolog
	f.App.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
	}))

	// Initialize recover
	f.App.Use(recover.New(recover.Config{EnableStackTrace: true}))
}

func (f *Fiber) Start(host, port string) {
	if err := f.App.Listen(fmt.Sprintf("%s:%s", host, port)); err != nil {
		log.Fatal().Err(err).Msg("Fiber app error")
	}
}
