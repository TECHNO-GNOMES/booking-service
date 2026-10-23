package user

import (
	"context"
	"errors"
	"time"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

type UpdateUserRequest struct {
	Avatar    *string
	FirstName *string
	LastName  *string
	Gender    *string
	City      *string
	Birthday  *time.Time
	Address   *string
	Phone     *string
	Email     *string
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(ctx context.Context, user *models.User) error {
	query := u.db.WithContext(ctx).Create(user).Error
	return query
}

func (u *UserRepo) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) FindByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := u.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []models.User{}, nil
		}
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) Update(ctx context.Context, id uint, req *UpdateUserRequest) error {
	updates := make(map[string]interface{})

	if req.FirstName != nil {
		updates["first_name"] = *req.FirstName
	}

	if req.LastName != nil {
		updates["last_name"] = *req.LastName
	}

	if req.Gender != nil {
		updates["gender"] = *req.Gender
	}

	if req.City != nil {
		updates["city"] = *req.City
	}

	if req.Address != nil {
		updates["address"] = *req.Address
	}

	if req.Phone != nil {
		updates["phone"] = *req.Phone
	}

	if req.Email != nil {
		updates["email"] = *req.Email
	}

	if req.Birthday != nil {
		updates["birthday"] = *req.Birthday
	}

	if req.Avatar != nil {
		updates["avatar"] = *req.Avatar
	}

	return u.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
}

func (u *UserRepo) Delete(ctx context.Context, id uint) error {
	result := u.db.WithContext(ctx).Delete(&models.User{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("property not found")
	}

	return nil
}
