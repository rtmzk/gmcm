package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"gmcm/pkg/log"
	"gmcm/pkg/routes"
	"gmcm/pkg/utils"
	"gmcm/server"
	"math/rand"
	"os"
	"runtime"
	"time"
)

var run = &cobra.Command{
	Use:   "run",
	Short: "start app",
	RunE: func(cmd *cobra.Command, args []string) error {
		rand.Seed(time.Now().UTC().UnixNano())
		if len(os.Getenv("GOMAXPROCS")) == 0 {
			runtime.GOMAXPROCS(runtime.NumCPU())
		}
		err := godotenv.Load(DefaultConfigPath)
		if err != nil {
			fmt.Println(".env file not found.")
			os.Exit(1)
		}

		log.Init()
		defer log.Flush()

		LISTEN := utils.GetEnv("app.listen", ":9527")

		r := gin.New()
		routes.SetRoutes(r)

		srv := server.NewServer(r)

		return srv.Start(LISTEN)
	},
}
