package services

import (
	"errors"
	m "golang_app/golangApp/app/models"
	"golang_app/golangApp/config/localize"
	"golang_app/golangApp/config/redis"

	"gorm.io/gorm"
)

// Ensure UserService implements the BaseService interface
var _ BaseService = (*UserService)(nil)

// UserService is the ActiveRecord-ish layer for the users table.
type UserService struct {
	db          *gorm.DB
	translation *localize.Localization
	redis       *redis.RedisClient
}

func NewUserService(database *gorm.DB, locale *localize.Localization, redis *redis.RedisClient) *UserService {
	return &UserService{db: database, translation: locale, redis: redis}
}

func (s *UserService) ConvertToMap(data interface{}) (map[string]interface{}, error) {
	return StructToMap(data)
}

func (s *UserService) CreateUser(user m.User, locale string) (string, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_CREATE_USER", locale))
	}
	return user.ID, nil
}

func (s *UserService) UpdateUserById(id string, updates m.User, locale string) (string, error) {
	if !ValidUUID(id) {
		return "", errors.New(s.translation.GetMessage("USER_NOT_FOUND", locale))
	}
	data, err := s.ConvertToMap(updates)
	if err != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_UPDATE_USER", locale))
	}
	// The primary key must never be part of the update payload.
	delete(data, "id")
	if len(data) == 0 {
		return "", errors.New(s.translation.GetMessage("NO_DATA_USER_UPDATED", locale))
	}

	result := s.db.Model(&m.User{}).Where("id = ?", id).Updates(data)
	if result.Error != nil {
		return "", errors.New(s.translation.GetMessage("FAILED_UPDATE_USER", locale))
	}
	if result.RowsAffected == 0 {
		return "", errors.New(s.translation.GetMessage("NO_DATA_USER_UPDATED", locale))
	}
	return s.translation.GetMessage("SUCCESS_UPDATE_USER", locale), nil
}

func (s *UserService) GetUserByEmail(email string) *m.User {
	result := m.User{}
	if err := s.db.Where("email = ?", email).First(&result).Error; err != nil {
		return nil
	}
	return &result
}

func (s *UserService) DeleteUserById(id string) error {
	if !ValidUUID(id) {
		return gorm.ErrRecordNotFound
	}
	return s.db.Where("id = ?", id).Delete(&m.User{}).Error
}

func (s *UserService) GetUserByEmailAndPassword(email string, password string) (*m.User, error) {
	var result m.User
	err := s.db.Where("email = ? AND password = ?", email, password).First(&result).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func (s *UserService) GetAllUserList(page, pageSize int) ([]m.User, error) {
	var results []m.User
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}
	err := s.db.
		Order("created_at ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&results).Error
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, nil
	}
	return results, nil
}

func (s *UserService) SearchUser(searchType string, query string) *[]m.User {
	results := []m.User{}
	search := s.db.Model(&m.User{})

	switch searchType {
	case "name":
		pattern := "%" + query + "%"
		search = search.Where("firstname ILIKE ? OR lastname ILIKE ?", pattern, pattern)
	case "email":
		search = search.Where("email = ?", query)
	default:
		return nil
	}

	if err := search.Find(&results).Error; err != nil {
		return nil
	}
	if len(results) == 0 {
		return nil // No users found, return nil
	}
	return &results
}
