package main

import (
	// "database/sql"
	// "github.com/mrcruz117/charm-ref/internal/database"
	// "log"

	"github.com/mrcruz117/charm-ref/cmd"
)

func main() {

	// db, err := sql.Open("sqlite3", "sqlite.db")
	// if err != nil {
	// 	log.Fatalf("could not open db: %v", err)
	// }
	// dbQueries := database.New(db)

	// cfg := config{
	// 	db: dbQueries,
	// }
	cmd.Cfg = cmd.InitConfig()
	cmd.Execute()

}
