package member_service

import (
	"daily-brew/models"
	"daily-brew/utils"
)

type Member struct {
	ID       uint
	Email    string
	FullName string
	Phone    string
	Password string
}

func mapModelToService(m *models.Member) *Member {
	return &Member{
		ID:       m.ID,
		FullName: m.FullName,
		Email:    m.Email,
		Phone:    m.Phone,
		Password: m.Password,
	}
}

func (m *Member) GetMemberByEmail() (*Member, error) {
	member := &models.Member{
		Email: m.Email,
	}
	err := member.GetMemberByEmail()

	if err != nil {
		return nil, err
	}

	newMember := mapModelToService(member)
	return newMember, nil
}

func (m *Member) GetMemberByID() (*Member, error) {
	member := &models.Member{}
	err := member.GetMemberByID(m.ID)

	if err != nil {
		return nil, err
	}

	newMember := mapModelToService(member)
	return newMember, nil
}

func (m *Member) GetMember() (*Member, error) {
	memberGet := &models.Member{
		Email: m.Email,
	}
	err := memberGet.GetMemberByEmail()

	if err != nil {
		return nil, err
	}

	member := &Member{
		ID:       memberGet.ID,
		Email:    memberGet.Email,
		FullName: memberGet.FullName,
		Phone:    memberGet.Phone,
	}

	return member, nil
}

func (m *Member) Register() error {

	hashPassword, err := utils.BcryptHash(m.Password)
	if err != nil {
		return err
	}

	var newMember models.Member
	newMember = models.Member{
		Email:    m.Email,
		FullName: m.FullName,
		Phone:    m.Phone,
		Password: hashPassword,
	}

	err = newMember.Create()

	if err != nil {
		return err
	}
	return nil
}
