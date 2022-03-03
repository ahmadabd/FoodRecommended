package mysql_test

import (
	"database/sql"
	"os"
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository/mysql"
	"github.com/ahmadabd/FoodRecommended.git/mocks"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (repository.DB, func()) {
	// DB Connection
	db, err := sql.Open("mysql", "root:a142332144113@tcp(localhost:3306)/food_recommended_test")
	assert.Nil(t, err)

	// Migrations
	err = Migrations(db)

	assert.Nil(t, err)

	// Mock Logger
	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	mockLogger := mocks.NewMockLogger(mockctl)

	return &mysql.Mysql{Db: db, Logger: mockLogger}, func() {
		_, err := db.Exec(`
			DROP TABLE IF EXISTS foods;`)

		assert.Nil(t, err)

		// Close DB
		err = db.Close()
		assert.Nil(t, err)
	}
}

func Migrations(db *sql.DB) error {
	// Migrations
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS foods (
			id int auto_increment, 
			name nvarchar(20) not null, 
			country nvarchar(20), 
			city nvarchar(20), 
			vegetarian int default 0,
			primary key (id)
		);`)

	return err
}
