package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/danmory/messaging/consumer/models"
)

type Repository struct {
	db *sql.DB
}

func (pdb *Repository) Init(cfg *models.Config) (error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	pdb.db = db
	return nil
}

func (pdb *Repository) Save(message *models.Message) error {
	if message.Table == 0 {
		return pdb.saveToTable1( message)
	}
	return pdb.saveToTable2(message)
}

func (pdb *Repository) saveToTable1(message *models.Message) error {
	_, err := pdb.db.Exec("INSERT INTO table1 (message) VALUES ($1)", message.Text)
	return err
}

func (pdb *Repository) saveToTable2(message *models.Message) error {
	_, err := pdb.db.Exec("INSERT INTO table2 (message) VALUES ($1)", message.Text)
	return err
}

func (pdb *Repository) Close() error {
	return pdb.db.Close()
}