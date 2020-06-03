package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	//	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/alanyeung95/GoProjectDemo/pkg/config"
	"github.com/alanyeung95/GoProjectDemo/pkg/demo"
	"github.com/alanyeung95/GoProjectDemo/pkg/items"
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

			demoSrv, err := newDemoSrv()
			if err != nil {
				return err
			}

			itemSrv, err := newItemSrv()
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

func newItemSrv() (items.Service, error) {
	//func newItemSrv( client *mongodriver.Client,) ( items.Service, error) {
	//itemRepository, err := mongo.NewPageRepository(client, cfg.MongoDB.Database, cfg.MongoDB.NewsItemCollection, cfg.MongoDB.PubHistoryCollection, cfg.MongoDB.EnableSharding)
	//if err != nil {
	//		return nil, nil, err
	//}

	itemSrv, err := items.NewService()
	if err != nil {
		return nil, err
	}

	return itemSrv, nil
}
