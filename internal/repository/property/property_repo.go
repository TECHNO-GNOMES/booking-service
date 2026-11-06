package property

import (
	"context"
	"errors"

	"github.com/nurgal1ev/booking-service/internal/models"
	"gorm.io/gorm"
)

type PropertyRepo struct {
	db *gorm.DB
}

func NewPropertyRepo(db *gorm.DB) *PropertyRepo {
	return &PropertyRepo{db: db}
}

type PropertyFilter struct {
	Search         string
	CategoryID     uint
	LocationTypeID uint
	GuestCountMin  *int
	MinPrice       *int
	MaxPrice       *int
	City           string
}

type UpdatePropertyRequest struct {
	Name          *string
	Description   *string
	Address       *string
	City          *string
	Country       *string
	PricePerNight *int
	PropertyType  *string
}

func (p *PropertyRepo) Create(ctx context.Context, property *models.Property) error {
	query := p.db.WithContext(ctx).Create(property).Error
	return query
}

func (p *PropertyRepo) FindByID(ctx context.Context, id uint) (*models.Property, error) {
	var property models.Property
	err := p.db.WithContext(ctx).Preload("Reviews", nil).Where("id = ?", id).First(&property).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &property, nil
}

func (r *PropertyRepo) FindAll(ctx context.Context, filter PropertyFilter) ([]models.Property, error) {
	query := r.db.WithContext(ctx).Model(&models.Property{})

	if filter.Search != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	if filter.CategoryID > 0 {
		query = query.Where("category_id = ?", filter.CategoryID)
	}

	if filter.LocationTypeID > 0 {
		query = query.Where("location_type_id = ?", filter.LocationTypeID)
	}

	if filter.City != "" {
		query = query.Where("city = ?", filter.City)
	}

	if filter.GuestCountMin != nil {
		query = query.Where("guest_count_banquet >= ?", *filter.GuestCountMin)
	}

	if filter.MinPrice != nil {
		query = query.Where("rental_price >= ?", *filter.MinPrice)
	}
	if filter.MaxPrice != nil {
		query = query.Where("rental_price <= ?", *filter.MaxPrice)
	}

	var properties []models.Property
	err := query.Find(&properties).Error
	return properties, err
}

func (p *PropertyRepo) FindByOwnerID(ctx context.Context, id uint) ([]models.Property, error) {
	var properties []models.Property
	err := p.db.WithContext(ctx).Where("owner_id = ?", id).Find(&properties).Error
	if err != nil {
		return nil, err
	}

	return properties, nil
}

func (p *PropertyRepo) Update(ctx context.Context, id uint, req *UpdatePropertyRequest) error {
	updates := make(map[string]interface{})

	if req.Name != nil {
		updates["name"] = *req.Name
	}

	if req.Description != nil {
		updates["description"] = *req.Description
	}

	if req.Address != nil {
		updates["address"] = *req.Address
	}

	if req.City != nil {
		updates["city"] = *req.City
	}

	if req.Country != nil {
		updates["country"] = *req.Country
	}

	if req.PricePerNight != nil {
		updates["price_per_night"] = *req.PricePerNight
	}

	return p.db.WithContext(ctx).Model(&models.Property{}).Where("id = ?", id).Updates(updates).Error
}

func (p *PropertyRepo) Delete(ctx context.Context, id uint) error {
	result := p.db.WithContext(ctx).Delete(&models.Property{}, id)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("property not found")
	}

	return nil
}
