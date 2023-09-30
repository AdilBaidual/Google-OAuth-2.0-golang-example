package main

import (
	"Kokos"
	"Kokos/internal/handler"
	"Kokos/internal/repo"
	"Kokos/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env error: %s", err.Error())
	}

	db, err := repo.NewPostgresDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("db initialize error: %s", err.Error())
	}

	conf := &oauth2.Config{
		ClientID:     os.Getenv("G_CLIENT_ID"),
		ClientSecret: os.Getenv("G_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("G_REDIRECT"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	repository := repo.NewRepo(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services, conf)
	srv := new(Kokos.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoute()); err != nil {
		logrus.Fatalf("Error starting server")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
