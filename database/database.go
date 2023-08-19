package database

import (
	"database/sql"
	// "github.com/google/uuid"
)

type Database struct {
	Client *sql.DB
}

// type ConsumerSubsystemStateTableRow struct {
// 	ConsumerId uuid.UUID
// 	UpdatedAt  time.Time
// }

// type InventoriesTableRow struct {
// 	ItemId    uuid.UUID
// 	OwnerId   uuid.UUID
// 	UpdatedAt time.Time
// }

func NewDatabase(connStr string) (*Database, error) {
	//open database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// //check for liveness
	// err = db.Ping()
	// if err != nil {
	// 	return nil, err
	// }

	// //run migrations
	// path := "database/migrations.sql"
	// file, err := ioutil.ReadFile(path)
	// if err != nil {
	// 	return nil, err
	// }
	// sql := string(file)
	// res, err := db.Exec(sql)
	// if err != nil {
	// 	return nil, err
	// }
	// _ = res

	// fmt.Println("Database setup complete")

	return &Database{db}, nil
}
