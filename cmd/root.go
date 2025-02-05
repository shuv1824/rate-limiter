package cmd

import (
	"log"
	"time"

	"github.com/shuv1824/rate-limiter/internal/domain"
	"github.com/shuv1824/rate-limiter/internal/infra/http"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the app",
	Run:   rootRun,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func rootRun(cmd *cobra.Command, args []string) {
	limiter := domain.NewFixedWindowLimiter(5, time.Second)

	r := http.SetupRouter(http.RateLimitMiddleware(limiter))

	r.Run(":8080")
}
