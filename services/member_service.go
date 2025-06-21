package services

import (
	"database/sql"
	"errors"
	member "gin-sample/models"
)

type MemberService struct {
	db *sql.DB
}

func NewMemberService(db *sql.DB) *MemberService {
	return &MemberService{db: db}
}

func (s *MemberService) GetMembers() ([]member.Member, error) {
	return member.GetMembers(s.db)
}

func (s *MemberService) GetMemberById(id int) (*member.Member, error) {
	return member.GetMemberById(s.db, id)
}

func (s *MemberService) AddMember(name string, age int, sex member.Sex) error {
	m := member.NewMember(name, age, sex)
	return m.AddMember(s.db)
}

func (s *MemberService) AddMemberWithTransaction(tx *sql.Tx, name string, age int, sex member.Sex) error {
	if tx == nil {
		return errors.New("トランザクションが見つかりません")
	}
	
	m := member.NewMember(name, age, sex)
	return m.AddMember(tx)
}