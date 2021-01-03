package log

import (
	"os"

	"github.com/rs/zerolog"
)

// We do not want to define a Logger interface. Stick to concrete impl,
// because for some reason Go ;o/
func NewZerolog(target *os.File, level string) (*zerolog.Logger, error) {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	zerolog.SetGlobalLevel(logLevel)
	l := zerolog.New(target).With().Timestamp().Logger()

	return &l, nil
}
