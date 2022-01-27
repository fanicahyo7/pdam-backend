package repository

import (
	"pdam/model"

	"gorm.io/gorm"
)

type grouprepository struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) *grouprepository {
	return &grouprepository{db}
}

type GroupRepository interface {
	FindGroupByAll() ([]model.Group, error)
	FindGroupByID(kode string) (model.Group, error)
	FindGroupByName(name string) ([]model.Group, error)
	CreateGroup(group model.Group) (model.Group, error)
	UpdateGroup(group model.Group) (model.Group, error)
	DeleteGroup(kode string) error
}

func (r *grouprepository) FindGroupByAll() ([]model.Group, error) {
	var group []model.Group

	err := r.db.Find(&group).Error
	if err != nil {
		return group, err
	}
	return group, nil
}

func (r *grouprepository) FindGroupByID(kode string) (model.Group, error) {
	var group model.Group

	err := r.db.Where("Kode = ?", kode).Find(&group).Error

	if err != nil {
		return group, err
	}

	return group, nil
}

func (r *grouprepository) FindGroupByName(name string) ([]model.Group, error) {
	var group []model.Group

	err := r.db.Where("Nama like ?", "%"+name+"%").Find(&group).Error

	if err != nil {
		return group, err
	}

	return group, nil
}

func (r *grouprepository) CreateGroup(group model.Group) (model.Group, error) {
	err := r.db.Create(&group).Error
	if err != nil {
		return group, err
	}

	return group, nil
}

func (r *grouprepository) UpdateGroup(group model.Group) (model.Group, error) {
	err := r.db.Save(&group).Error
	if err != nil {
		return group, err
	}

	return group, nil
}

func (r *grouprepository) DeleteGroup(kode string) error {
	var group model.Group
	err := r.db.Where("Kode = ?", kode).First(&group).Error
	if err != nil {
		return err
	}
	r.db.Delete(&group)
	return nil
}
