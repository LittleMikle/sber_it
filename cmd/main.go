package main

import (
	"context"
	todo "github.com/LittleMikle/sber_it"
	"github.com/LittleMikle/sber_it/pkg/controller"
	"github.com/LittleMikle/sber_it/pkg/repository"
	"github.com/LittleMikle/sber_it/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal().Msg("error with viper")
	} else {
		log.Info().Msg("Config initiation successful")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal().Msgf("error with .env file %s", err)
	} else {
		log.Info().Msg("Config load successful")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal().Msg("failed with Postgres connection")
	} else {
		log.Info().Msg("Connection to Postgres successful")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	controllers := controller.NewController(services)

	srv := new(todo.Server)

	go func() {
		err = srv.Run(viper.GetString("port"), controllers.InitRoutes())
		if err != nil {
			log.Fatal().Msg("")
		}
	}()

	log.Info().Msg("Starting server successful")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info().Msg("Shutting down server successful")

	err = srv.Shutdown(context.Background())
	if err != nil {
		log.Error().Msg("failed with shutting down")
	}
	err = db.Close()
	if err != nil {
		log.Error().Msg("failed with closing DB connection")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
