package model

import (
	"github.com/DevtronLabs/CatPicHub/internal/providers/database"
)

type CatPics struct {
	ID             string `json:"id" gorm:"primaryKey"`
	CatPicFileName string `json:"cat_pic_file_name" gorm:"unique;not null"`
	CatPicData     []byte `json:"cat_pic_data" gorm:"type:BLOB"`
	CatPicFileType string `json:"cat_pic_file_type"`
}

// Creates a new entry in cat_pics table
func (cps *CatPics) Create() error {
	err := database.Client().Create(&cps).Error
	return err
}

// Fetched a entry in cat_pics table by its ID
func (cps *CatPics) GetByID(id string) error {
	err := database.Client().Where("id = ?", id).First(&cps).Error
	return err
}

// Fetched a list of entries in cat_pics table using pagination concept
func GetPaginatedData(page, pageSize int) ([]CatPics, error) {
	var catPics []CatPics
	result := database.Client().Select("id, cat_pic_file_name, cat_pic_file_type").Offset((page - 1) * pageSize).Limit(pageSize).Find(&catPics)
	if result.Error != nil {
		return nil, result.Error
	}

	return catPics, nil
}

// Updates a already existing entry in cat_pics table by its ID
func (cps *CatPics) Update(catPicUpdates map[string]interface{}) error {
	err := database.Client().Model(&cps).Updates(catPicUpdates).Error
	return err
}

// Deletes a entry in cat_pics table by its ID
func (cps *CatPics) DeleteById() error {
	dbClient := database.Client()
	err := dbClient.Delete(&cps).Error
	return err
}
