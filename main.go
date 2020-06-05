package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	//	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	mongodriver "go.mongodb.org/mongo-driver/mongo"

	"github.com/alanyeung95/GoProjectDemo/pkg/config"
	"github.com/alanyeung95/GoProjectDemo/pkg/demo"
	"github.com/alanyeung95/GoProjectDemo/pkg/items"
	"github.com/alanyeung95/GoProjectDemo/pkg/mongo"
)

func main() {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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

			// Route - Middlewares
			r := chi.NewRouter()

			// Route - API
			r.Route("/", func(r chi.Router) {
				r.Mount("/", items.NewHandler(demoSrv))
				r.Mount("/items", items.NewHandler(itemSrv))
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
	demoSrv, err := demo.NewService()
	if err != nil {
		return nil, err
	}

	return demoSrv, nil
}

//func newItemSrv() (items.Service, error) {
func newItemSrv(cfg config.AppConfig, client *mongodriver.Client) (items.Service, error) {
	itemRepository, err := mongo.NewItemRepository(client, cfg.MongoDB.Database, cfg.MongoDB.ItemCollection, cfg.MongoDB.EnableSharding)

	if err != nil {
		return nil, err
	}

	itemSrv, err := items.NewService(itemRepository)
	if err != nil {
		return nil, err
	}

	return itemSrv, nil
}
