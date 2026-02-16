package sqlite

import (
	"time"

	"github.com/Amirhos-esm/portfolio/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MessageGorm struct {
	DB *gorm.DB
}

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto create table
	err = db.AutoMigrate(&models.Message{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewMessageGorm() (*MessageGorm, error) {
	db, err := InitDB()
	if err != nil {
		return nil, err
	}
	return &MessageGorm{DB: db}, nil
}

func (r *MessageGorm) Create(m *models.Message) error {

	if m.Createat.IsZero() {
		m.Createat = time.Now()
	}

	return r.DB.Create(m).Error
}
func (r *MessageGorm) GetbyId(id uint) (*models.Message, error) {

	var m models.Message

	err := r.DB.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *MessageGorm) Get(offset uint, page uint) ([]*models.Message, error) {

	var list []*models.Message

	sqlOffset := int(offset * page)

	err := r.DB.
		Order("id DESC").
		Limit(int(page)).
		Offset(sqlOffset).
		Find(&list).Error

	return list, err
}
func (r *MessageGorm) Delete(id uint) error {
	return r.DB.Delete(&models.Message{}, id).Error
}
