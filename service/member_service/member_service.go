package member_service

import (
	"daily-brew/models"
	"daily-brew/utils"
)

type Member struct {
	ID       uint
	Email    string
	Name     string
	Phone    string
	Password string
}

func (m *Member) GetMemberByEmail() (*Member, error) {
	member, err := models.GetMemberByEmail(m.Email)

	if err != nil {
		return nil, err
	}

	newMember := &Member{
		ID:       member.ID,
		Email:    member.Email,
		Name:     member.Name,
		Phone:    member.Phone,
		Password: member.Password,
	}
	return newMember, nil
}

func (m *Member) GetMemberByID() (*Member, error) {
	member, err := models.GetMemberByID(m.ID)

	if err != nil {
		return nil, err
	}

	newMember := &Member{
		ID:       member.ID,
		Email:    member.Email,
		Name:     member.Name,
		Phone:    member.Phone,
		Password: member.Password,
	}
	return newMember, nil
}

func (m *Member) GetMember() (*Member, error) {
	memberGet, err := models.GetMemberByEmail(m.Email)

	if err != nil {
		return nil, err
	}

	member := &Member{
		ID:    memberGet.ID,
		Email: memberGet.Email,
		Name:  memberGet.Name,
		Phone: memberGet.Phone,
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
		Name:     m.Name,
		Phone:    m.Phone,
		Password: hashPassword,
	}

	err = newMember.Save()

	if err != nil {
		return err
	}
	return nil
}
