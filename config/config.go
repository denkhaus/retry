package config

import (
	"flag"
	"time"

	"github.com/itzg/go-flagsfiller"
)

type Config struct {
	Version  bool   `flag:"v" usage:"show version and exit" env:""`
	Command  string `flag:"c" default:"" usage:"command to execute"`
	LogLevel string `flag:"l" default:"info" usage:"log level - debug, info, warn, error"`

	Interval   time.Duration `flag:"i" default:"1s" usage:"constant backoff: interval between retries"`
	RetriesMax int64         `flag:"m" default:"-1" usage:"constant/exponential backoff: max retries | -1 for no limit"`

	Exponential struct {
		Enabled             bool          `flag:"e" default:"false" usage:"exponential backoff: enable"`
		InitialInterval     time.Duration `flag:"ei" default:"1s" usage:"exponential backoff: initial interval between retries" `
		MaxInterval         time.Duration `flag:"ex" default:"10s" usage:"exponential backoff: maximum interval between retries"`
		MaxElapsedTime      time.Duration `flag:"et" default:"1h" usage:"exponential backoff: maximum total time for retries"`
		Multiplier          float64       `flag:"em" default:"2.0" usage:"exponential backoff: retry duration multiplier"`
		RandomizationFactor float64       `flag:"ej" default:"0.0" usage:"exponential backoff: add jitter to the retry interval like Interval = RetryInterval * (1 ± RandomizationFactor)"`
	}
}

func Parse(from interface{}) error {
	filler := flagsfiller.New(flagsfiller.WithEnv("retry"))
	err := filler.Fill(flag.CommandLine, from)
	if err != nil {
		return err
	}

	flag.Parse()
	return nil
}
