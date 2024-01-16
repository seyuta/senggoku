package bootstrap

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

var (
	KoanfYamlFile *koanf.Koanf
)

type Koanf struct {
	New *koanf.Koanf
}

func NewKoanf(delimiter string) *Koanf {
	return &Koanf{New: koanf.New(delimiter)}
}

func (k *Koanf) UseYamlFile(path string, isWatch ...bool) {
	f := file.Provider(path)
	if err := k.New.Load(f, yaml.Parser()); err != nil {
		log.Fatal().Err(err).Msg("error loading config")
	}
	KoanfYamlFile = k.New

	if len(isWatch) > 0 {
		if isWatch[0] {
			f.Watch(func(event interface{}, err error) {
				if err != nil {
					log.Fatal().Err(err).Msg("koanf watch yaml file error")
					return
				}

				// Throw away the old config and load a fresh copy.
				log.Info().Msg("config changed. Reloading ...")
				KoanfYamlFile = NewKoanf(".").New
				if err := KoanfYamlFile.Load(f, yaml.Parser()); err != nil {
					log.Fatal().Err(err).Msg("error loading config")
				}
			})
		}
	}
}
