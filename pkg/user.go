package pkg

import (
	"os/user"

	zlog "github.com/rs/zerolog/log"
)

// IsRoot detect whether this cli is run using root or not
func IsRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		zlog.Fatal().Err(err).Msg("cannot detect process user")
	}
	return currentUser.Username == "root"
}
