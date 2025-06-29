package app

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/fahrilhadi/blog-portal-news/config"
	"github.com/fahrilhadi/blog-portal-news/lib/auth"
	"github.com/rs/zerolog/log"
)

func RunServer()  {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgres()
	if err != nil {
		log.Fatal().Msgf("Error connecting to database: %v", err)
		return
	}

	// Cloudflare R2
	cdfR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(cdfR2)

	_ = auth.NewJwt(cfg)
}