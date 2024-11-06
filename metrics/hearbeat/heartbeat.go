package heartbeat

import (
	"net/http"

	"github.com/squid-labs/gLibrary/types"
	"github.com/rs/zerolog/log"
)

func Send(conf types.HeartbeatConfig) {
	for _, u := range conf.URLs {
		log.Info().Str("url", u).Msg("sending heartbeat")
		_, err := http.Get(u)
		if err != nil {
			log.Error().Str("monitoring", "heartbeat").Msg(err.Error())
		}
	}
}
