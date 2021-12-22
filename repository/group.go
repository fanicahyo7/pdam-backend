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
	FindGroupByID(id int) (model.Group, error)
	FindGroupByName(name string) (model.Group, error)
}

func (r *grouprepository) FindGroupByID(id int) (model.Group, error) {
	var group model.Group

	err := r.db.Where("KodeGroup = ?", id).Find(&group).Error

	if err != nil {
		return group, err
	}

	return group, nil
}

func (r *grouprepository) FindGroupByName(name string) (model.Group, error) {
	var group model.Group

	err := r.db.Where("Name = ?", name).Find(&group).Error

	if err != nil {
		return group, err
	}

	return group, nil
}
