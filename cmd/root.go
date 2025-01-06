/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrcruz117/charm-ref/internal/database"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "charm-ref",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
type config struct {
	db   *database.Queries
	conn *sql.DB
}

func InitConfig() *config {
	// Get the absolute path to the database file
	dbPath, err := filepath.Abs("./db/sqlite.db")
	if err != nil {
		log.Fatalf("could not get absolute path: %v", err)
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("could not open db: %v", err)
	}
	dbQueries := database.New(db)
	return &config{
		db:   dbQueries,
		conn: db,
	}
}

func (c *config) Close() {
	if err := c.conn.Close(); err != nil {
		log.Fatalf("could not close db: %v", err)
	}
}

var Cfg *config

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.charm-ref.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
