package services

import (
	"errors"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

	"gorm.io/gorm"
)

// Ensure WeddingService implements the BaseService interface
var _ BaseService = (*WeddingService)(nil)

// WeddingService is the ActiveRecord-ish layer for the wedding_data table.
type WeddingService struct {
	db          *gorm.DB
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewWeddingService(database *gorm.DB, locale *localize.Localization, redis *redis.RedisClient) *WeddingService {
	return &WeddingService{db: database, translation: locale, redis: redis}
}

func (s *WeddingService) ConvertToMap(data interface{}) (map[string]interface{}, error) {
	return StructToMap(data)
}

func (s *WeddingService) CreateWeddingData(data m.WeddingData, locale string) (string, error) {
	if err := s.db.Create(&data).Error; err != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_CREATE_WEDDING", locale))
	}
	return data.ID, nil
}

func (s *WeddingService) UpdateWeddingDataById(id string, updates m.WeddingData, locale string) (string, error) {
	if !ValidUUID(id) {
		return "", gorm.ErrRecordNotFound
	}
	if updates.UserID != nil {
		updates.UserID = NormalizeUUID(*updates.UserID)
	}

	// Select forces every payload column to be written, including empty ones,
	// which is what a full form submit is expected to do.
	result := s.db.Model(&m.WeddingData{}).
		Where("id = ?", id).
		Select("user_id", "bride", "groom", "wedding", "love_story", "envelop", "gallery", "guest", "streaming", "gifts").
		Updates(updates)
	if result.Error != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_UPDATE_WEDDING", locale))
	}
	if result.RowsAffected == 0 {
		return "", errors.New(s.translation.GetMessage("NO_DATA_WEDDING_UPDATED", locale))
	}
	return s.translation.GetMessage("SUCCESS_UPDATE_WEDDING", locale), nil
}

func (s *WeddingService) GetWeddingDataById(id string) (*m.WeddingData, error) {
	if !ValidUUID(id) {
		return nil, gorm.ErrRecordNotFound
	}
	result := m.WeddingData{}
	if err := s.db.Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WeddingService) GetWeddingDataByUserId(userId string) (*m.WeddingData, error) {
	normalized := NormalizeUUID(userId)
	if normalized == nil {
		return nil, gorm.ErrRecordNotFound
	}
	result := m.WeddingData{}
	if err := s.db.Where("user_id = ?", *normalized).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WeddingService) DeleteWeddingDataById(id string) error {
	if !ValidUUID(id) {
		return gorm.ErrRecordNotFound
	}
	return s.db.Where("id = ?", id).Delete(&m.WeddingData{}).Error
}
