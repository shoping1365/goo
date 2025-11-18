package repository

import (
	"fmt"
	"my-go-backend/internal/models"

	"gorm.io/gorm"
)

// SMSGatewayRepository عملیات دیتابیس مربوط به درگاه پیامک را مدیریت می‌کند
// شامل ثبت، ویرایش، حذف، لیست و دریافت پترن‌ها

type SMSGatewayRepository struct {
	DB *gorm.DB
}

// NewSMSGatewayRepository سازنده ریپازیتوری
func NewSMSGatewayRepository(db *gorm.DB) *SMSGatewayRepository {
	return &SMSGatewayRepository{DB: db}
}

// ایجاد درگاه جدید
func (r *SMSGatewayRepository) Create(gateway *models.SMSGateway) error {
	return r.DB.Create(gateway).Error
}

// ویرایش درگاه
func (r *SMSGatewayRepository) Update(gateway *models.SMSGateway) error {
	fmt.Printf("Repository Update - به‌روزرسانی درگاه %s (ID: %d, Priority: %d)\n", gateway.Name, gateway.ID, gateway.Priority)
	err := r.DB.Save(gateway).Error
	if err != nil {
		fmt.Printf("Repository Update - خطا: %v\n", err)
	} else {
		fmt.Printf("Repository Update - موفق\n")
	}
	return err
}

// حذف درگاه (حذف سخت)
func (r *SMSGatewayRepository) Delete(id uint) error {
	return r.DB.Delete(&models.SMSGateway{}, id).Error
}

// دریافت لیست همه درگاه‌ها
func (r *SMSGatewayRepository) List() ([]models.SMSGateway, error) {
	var gateways []models.SMSGateway
	err := r.DB.Preload("Patterns").Order("priority ASC, id ASC").Find(&gateways).Error

	fmt.Printf("Repository List - تعداد درگاه‌ها: %d\n", len(gateways))
	for _, gateway := range gateways {
		fmt.Printf("Repository List - %s (ID: %d, Priority: %d, Active: %t)\n", gateway.Name, gateway.ID, gateway.Priority, gateway.IsActive)
	}

	return gateways, err
}

// دریافت یک درگاه با پترن‌ها
func (r *SMSGatewayRepository) GetByID(id uint) (*models.SMSGateway, error) {
	var gateway models.SMSGateway
	err := r.DB.Preload("Patterns").First(&gateway, id).Error
	if err != nil {
		return nil, err
	}
	return &gateway, nil
}

// --- عملیات روی پترن‌ها ---

// افزودن پترن جدید به یک درگاه
func (r *SMSGatewayRepository) AddPattern(pattern *models.SMSGatewayPattern) error {
	return r.DB.Create(pattern).Error
}

// حذف یک پترن
func (r *SMSGatewayRepository) DeletePattern(id uint) error {
	return r.DB.Delete(&models.SMSGatewayPattern{}, id).Error
}

// لیست پترن‌های یک درگاه
func (r *SMSGatewayRepository) ListPatterns(gatewayID uint) ([]models.SMSGatewayPattern, error) {
	var patterns []models.SMSGatewayPattern
	err := r.DB.Where("gateway_id = ?", gatewayID).Find(&patterns).Error
	return patterns, err
}

// دریافت درگاه‌های فعال
func (r *SMSGatewayRepository) GetActiveGateways() ([]models.SMSGateway, error) {
	var gateways []models.SMSGateway
	err := r.DB.Where("is_active = ?", true).Preload("Patterns").Order("priority ASC, id ASC").Find(&gateways).Error
	return gateways, err
}
