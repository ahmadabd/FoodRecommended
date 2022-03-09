package gorm_test

import (
	"os"
	"testing"

	"github.com/ahmadabd/FoodRecommended.git/internal/entity/model"
	"github.com/ahmadabd/FoodRecommended.git/internal/repository"
	mygorm "github.com/ahmadabd/FoodRecommended.git/internal/repository/gorm"
	"github.com/ahmadabd/FoodRecommended.git/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"

	// "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (repository.DB, func()) {
	db, err := gorm.Open(mysql.Open("root:Ahmad@13741112@tcp(localhost:3306)/food_recommended_test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	assert.Nil(t, err)

	err = db.AutoMigrate(&model.Food{})
	assert.Nil(t, err)

	mockctl := gomock.NewController(t)
	defer mockctl.Finish()
	mockLogger := mocks.NewMockLogger(mockctl)

	return &mygorm.Gorm{Db: db, Logger: mockLogger}, func() {}
}
