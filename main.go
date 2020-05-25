package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	//	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/alanyeung95/GoProjectDemo/pkg/demo"
)

func main() {
	fmt.Println("Starting server...")
	rootCmd := &cobra.Command{
		Use:   "demo",
		Short: "Service to demo go project",
	}
	rootCmd.AddCommand(startCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func startCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start Server",
		RunE: func(cmd *cobra.Command, args []string) error {

			demoSrv, err := newDemoSrv()
			if err != nil {
				return err
			}

			// Route - Middlewares
			r := chi.NewRouter()

			// Route - API
			r.Route("/", func(r chi.Router) {
				r.Mount("/", demo.NewHandler(demoSrv))
			})

			// Start server
			addr := fmt.Sprintf(":%d", 8080)
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
