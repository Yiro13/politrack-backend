package main

import "github.com/Yiro13/politrack-backend/database"

func main() {

	database.DBConnection()
	database.Migrate(database.DB)
}
