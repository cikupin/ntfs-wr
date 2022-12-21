package main

import (
	"os"
	"time"

	"github.com/cikupin/ntfs-wr/cmd"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	// // configure logger
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.Kitchen,
	})

	err := cmd.RootCmd.Execute()
	if err != nil {
		zlog.Fatal().Err(err).Msg("error executing command")
	}
}
