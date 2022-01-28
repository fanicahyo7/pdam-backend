package service

import (
	helper "pdam/Helper"
	"pdam/model"
	"pdam/repository"
)

type GroupService interface {
	GetGroups() ([]model.Group, error)
	GetGroupByID(input helper.InputKodeGetGroup) (model.Group, error)
	GetGroupByNama(input helper.InputNamaGetGroup) ([]model.Group, error)
	SaveGroup(input helper.Inputgroup) (model.Group, error)
	EditGroup(input helper.Inputgroup) (model.Group, error)
	DeleteGroup(input helper.InputKodeGetGroup) error
}

type groupService struct {
	repository repository.GroupRepository
}

func NewService(repository repository.GroupRepository) *groupService {
	return &groupService{repository}
}

func (s *groupService) GetGroups() ([]model.Group, error) {
	groups, err := s.repository.FindGroupByAll()
	if err != nil {
		return groups, err
	}

	return groups, nil
}

func (s *groupService) GetGroupByID(input helper.InputKodeGetGroup) (model.Group, error) {

	groups, err := s.repository.FindGroupByID(input.Kode)

	if err != nil {
		return groups, err
	}

	return groups, nil
}

func (s *groupService) GetGroupByNama(input helper.InputNamaGetGroup) ([]model.Group, error) {

	groups, err := s.repository.FindGroupByName(input.Nama)

	if err != nil {
		return groups, err
	}

	return groups, nil
}

func (s *groupService) SaveGroup(input helper.Inputgroup) (model.Group, error) {
	group := model.Group{}
	group.Kode = input.Kode
	group.Nama = input.Nama
	group.Tarif1 = input.Tarif1
	group.Tarif2 = input.Tarif2
	group.Abonemen = input.Abonemen
	group.Kompensasi = input.Kompensasi

	newGroup, err := s.repository.CreateGroup(group)
	if err != nil {
		return newGroup, err
	}

	return newGroup, nil
}

func (s *groupService) EditGroup(input helper.Inputgroup) (model.Group, error) {
	group, err := s.repository.FindGroupByID(input.Kode)
	if err != nil {
		return group, err
	}

	group.Nama = input.Nama
	group.Tarif1 = input.Tarif1
	group.Tarif2 = input.Tarif2
	group.Abonemen = input.Abonemen
	group.Kompensasi = input.Kompensasi

	updateGroup, err := s.repository.UpdateGroup(group)
	if err != nil {
		return updateGroup, err
	}

	return updateGroup, nil
}

func (s *groupService) DeleteGroup(input helper.InputKodeGetGroup) error {
	group, err := s.repository.FindGroupByID(input.Kode)
	if err != nil {
		return err
	}

	err = s.repository.DeleteGroup(group)
	if err != nil {
		return err
	}

	return nil
}
