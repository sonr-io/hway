package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/onsonr/hway/app"
	"github.com/onsonr/hway/config"
	hwayorm "github.com/onsonr/hway/internal/models"
	"github.com/onsonr/hway/pkg/common"
	"github.com/spf13/cobra"
)

// Command line flags
var (
	servePort      int    // Gateway http entry point (default 3000)
	chainID        string // Current chain ID (default sonr-testnet-1)
	ipfsGatewayURL string // IPFS gateway URL (default localhost:8080)
	sonrAPIURL     string // Sonr API URL (default localhost:1317)
	sonrGrpcURL    string // Sonr gRPC URL (default localhost:9090)
	sonrRPCURL     string // Sonr RPC URL (default localhost:26657)

	psqlHost string // PostgresSQL Host Flag
	psqlPort string // PostgresSQL Port Flag
	psqlUser string // PostgresSQL User Flag
	psqlPass string // PostgresSQL Password Flag
	psqlDB   string // PostgresSQL Database Flag
)

// main is the entry point for the application
func main() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hway",
		Short: "Sonr DID gateway",
		Run: func(cmd *cobra.Command, args []string) {
			env, err := loadEnvImplFromArgs(args)
			if err != nil {
				panic(err)
			}
			ipc, err := common.NewIPFS()
			if err != nil {
				panic(err)
			}
			dbq, err := setupSQLiteDB()
			if err != nil {
				panic(err)
			}
			e, err := app.New(env, ipc, dbq)
			if err != nil {
				panic(err)
			}
			if err := e.Start(fmt.Sprintf(":%d", env.GetServePort())); err != http.ErrServerClosed {
				log.Fatal(err)
				os.Exit(1)
				return
			}
		},
	}
	cmd.Flags().IntVar(&servePort, "serve-port", 3000, "Port to serve the gateway on")
	cmd.Flags().StringVar(&chainID, "chain-id", "sonr-testnet-1", "Chain ID")
	cmd.Flags().StringVar(&ipfsGatewayURL, "ipfs-gateway-url", "localhost:8080", "IPFS gateway URL")
	cmd.Flags().StringVar(&sonrAPIURL, "sonr-api-url", "localhost:1317", "Sonr API URL")
	cmd.Flags().StringVar(&sonrGrpcURL, "sonr-grpc-url", "localhost:9090", "Sonr gRPC URL")
	cmd.Flags().StringVar(&sonrRPCURL, "sonr-rpc-url", "localhost:26657", "Sonr RPC URL")
	cmd.Flags().StringVar(&psqlHost, "psql-host", "localhost", "PostgresSQL Host")
	cmd.Flags().StringVar(&psqlPort, "psql-port", "5432", "PostgresSQL Port")
	cmd.Flags().StringVar(&psqlUser, "psql-user", "highway_user", "PostgresSQL User")
	cmd.Flags().StringVar(&psqlPass, "psql-pass", "highway_password123", "PostgresSQL Password")
	cmd.Flags().StringVar(&psqlDB, "psql-db", "highway", "PostgresSQL Database")
	return cmd
}

func loadEnvImplFromArgs(args []string) (config.Hway, error) {
	cmd := rootCmd()
	if err := cmd.ParseFlags(args); err != nil {
		return nil, err
	}

	env := &config.HwayImpl{
		ServePort:      servePort,
		ChainId:        chainID,
		IpfsGatewayUrl: ipfsGatewayURL,
		SonrApiUrl:     sonrAPIURL,
		SonrGrpcUrl:    sonrGrpcURL,
		SonrRpcUrl:     sonrRPCURL,
	}
	return env, nil
}

func setupSQLiteDB() (*hwayorm.Queries, error) {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	return hwayorm.New(conn), nil
}
