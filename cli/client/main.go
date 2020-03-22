package main

import (
	"flag"

	"time"

	pbgameengine "github.com/M-APIS/m-game-engine/v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var addressPtr = flag.String("address", "localhost:60051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to dail m-game-engine gRPC service")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr).Msg("Failed to close connection")
		}
	}()

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c := pbgameengine.NewGameEngineClient(conn)

	if c == nil {
		log.Info().Msg("Client nil")
	}

	r, err := c.GetSize(timeoutCtx, &pbgameengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("Size", r.GetSize()).Msg("GetSize from m-game-engine microservice")
	} else {
		log.Error().Msg("Couldnt receive Size")
	}
}
