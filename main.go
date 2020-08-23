package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/go-chi/chi"

	//	"github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	mongodriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/alanyeung95/GoProjectDemo/pkg/config"
	"github.com/alanyeung95/GoProjectDemo/pkg/demo"
	"github.com/alanyeung95/GoProjectDemo/pkg/items"
	"github.com/alanyeung95/GoProjectDemo/pkg/mongo"
	"github.com/alanyeung95/GoProjectDemo/pkg/users"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	initSentry(cfg)

	fmt.Println("Starting server...")
	rootCmd := &cobra.Command{
		Use:   "demo",
		Short: "Service to demo go project",
	}
	rootCmd.AddCommand(startCmd(cfg))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func loadConfig() (config.AppConfig, error) {
	cfg := config.NewAppConfig()
	if err := cfg.LoadFromEnv(); err != nil {
		return cfg, err
	}
	if err := cfg.Validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func initSentry(cfg config.AppConfig) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: cfg.Sentry.DSN,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

func startCmd(cfg config.AppConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Server",
		RunE: func(cmd *cobra.Command, args []string) error {

			mongoClient, err := mongo.NewClient(
				cfg.MongoDB.Addresses,
				cfg.MongoDB.Username,
				cfg.MongoDB.Password,
				cfg.MongoDB.Database,
			)

			if err != nil {
				return err
			}

			demoSrv, err := newDemoSrv()
			if err != nil {
				return err
			}

			itemSrv, err := newItemSrv(cfg, mongoClient)
			if err != nil {
				return err
			}

			userSrv, err := newUserSrv(cfg, mongoClient)
			if err != nil {
				return err
			}

			// Route - Middlewares
			r := chi.NewRouter()

			// Route - API
			r.Route("/", func(r chi.Router) {
				r.Mount("/", demo.NewHandler(demoSrv))
				r.Mount("/items", items.NewHandler(itemSrv))
				r.Mount("/users", users.NewHandler(userSrv))
			})

			// Start server
			addr := fmt.Sprintf(":%d", cfg.API.Port)
			//logger.Infof("Start listening on %s", addr)
			fmt.Printf("Start listening on %d", cfg.API.Port)
			return http.ListenAndServe(addr, r)
		},
	}
}

func newDemoSrv() (demo.Service, error) {
	demoSrv := demo.NewService()

	return demoSrv, nil
}

func newItemSrv(cfg config.AppConfig, client *mongodriver.Client) (items.Service, error) {
	itemRepository, err := mongo.NewItemRepository(client, cfg.MongoDB.Database, cfg.MongoDB.ItemCollection, cfg.MongoDB.EnableSharding)

	if err != nil {
		return nil, err
	}

	itemSrv := items.NewService(itemRepository)

	return itemSrv, nil
}

func newUserSrv(cfg config.AppConfig, client *mongodriver.Client) (users.Service, error) {
	userRepository, err := mongo.NewUserRepository(client, cfg.MongoDB.Database, cfg.MongoDB.UserCollection, cfg.MongoDB.EnableSharding)

	if err != nil {
		return nil, err
	}

	userSrv := users.NewService(userRepository)

	return userSrv, nil
}
