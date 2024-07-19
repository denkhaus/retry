package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/cenkalti/backoff/v4"
	"github.com/denkhaus/retry/config"
	"github.com/denkhaus/retry/logging"
	"github.com/pkg/errors"
)

var (
	BuildVersion = "0.0.0"
	BuildDate    = "n/a"
	BuildCommit  = "n/a"
)

var (
	logger = logging.Logger()
	cnf    config.Config
)

// createBackOff creates a backoff strategy based on the configuration settings.
//
// This function reads the configuration settings to determine whether to use
// exponential backoff or constant backoff. It constructs the backoff strategy
// accordingly and returns the configured backoff object.
//
// Returns a backoff.BackOff object.
func createBackOff() backoff.BackOff {
	if cnf.Exponential.Enabled {
		logger.Debug("using exponential backoff")

		opts := []backoff.ExponentialBackOffOpts{
			backoff.WithMaxInterval(cnf.Exponential.MaxInterval),
			backoff.WithInitialInterval(cnf.Exponential.InitialInterval),
			backoff.WithMultiplier(cnf.Exponential.Multiplier),
			backoff.WithMaxElapsedTime(cnf.Exponential.MaxElapsedTime),
			backoff.WithRandomizationFactor(cnf.Exponential.RandomizationFactor),
		}

		if cnf.RetriesMax > 0 {
			return backoff.WithMaxRetries(
				backoff.NewExponentialBackOff(opts...),
				uint64(cnf.RetriesMax),
			)
		}

		return backoff.NewExponentialBackOff(opts...)
	}

	logger.Debug("using constant backoff")

	if cnf.RetriesMax > 0 {
		return backoff.WithMaxRetries(
			backoff.NewConstantBackOff(cnf.Interval),
			uint64(cnf.RetriesMax),
		)

	}

	return backoff.NewConstantBackOff(cnf.Interval)
}

// main is the entry point function that parses configuration, executes a command, and handles backoff strategies based on the configuration settings.
//
// No parameters.
// No return value.
func main() {

	if err := config.Parse(&cnf); err != nil {
		logger.Fatalf("can't create input flags: %v", err)
	}

	if cnf.Version {
		fmt.Printf("%s %s (%s)-(%s)\n", os.Args[0], BuildVersion, BuildCommit, BuildDate)
		os.Exit(0)
	}

	logging.SwitchLogLevel(cnf.LogLevel)

	command := cnf.Command
	if command == "" {
		logger.Fatalf("no command specified")
	}

	execute := func() error {
		logging.Logger().Debugf("executing: %s", command)
		cmd := exec.Command("sh", "-c", command)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		defer func() {
			if r := recover(); r != nil {
				logger.Errorf("recovered from panic: %v", r)
			}
		}()

		if err := cmd.Run(); err != nil {
			return errors.Wrapf(err, "command %s failed", command)
		}

		return nil
	}

	if err := backoff.Retry(execute, createBackOff()); err != nil {
		logger.Fatalf("error: %v", err)
	}
}
