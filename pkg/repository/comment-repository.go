package repository

import (
	"ecommerce/pkg/models"

	"github.com/google/uuid"
)

func (p *Postgres) AddComment(comment *models.Comments) error {
	if err := p.db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}
func (p *Postgres) GetAllComments() (*[]models.Comments, error) {
	var comments []models.Comments

	if err := p.db.Find(&comments).Error; err != nil {
		return nil, err
	}
	return &comments, nil
}
func (p *Postgres) GetCommentByID(id string) (*models.Comments, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	var comment models.Comments

	if err := p.db.Where("id = ?", uuid).First(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}
func (p *Postgres) DeleteComment(id string) error {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	if err := p.db.Where("id = ?", uuid).Delete(&models.Comments{}).Error; err != nil {
		return err
	}
	return nil
}
func (p *Postgres) FindCommentByProductID(id string) (*[]models.Comments, error) {

	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var comments []models.Comments

	if err := p.db.Where("product_id = ?", uuidID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return &comments, nil
}
func (p *Postgres) FindCommentByUserID(id string) (*[]models.Comments, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var comments []models.Comments

	if err := p.db.Where("customer_id = ?", uuidID).Find(&comments).Error; err != nil {
		return nil, err
	}

	return &comments, nil
}
