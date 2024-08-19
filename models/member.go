package member

import "database/sql"

type Member struct {
	Name string
	Age  int
	Sex  Sex
}

type Sex string

const (
	Male   Sex = "male"
	Female Sex = "female"
)

func (m *Member) GetName() string {
	return m.Name
}

func (m *Member) GetAge() int {
	return m.Age
}

// Memberを初期化する関数
func NewMember(name string, age int, sex Sex) *Member {
	return &Member{
		Name: name,
		Age:  age,
		Sex:  sex,
	}
}

func (m *Member) AddMember(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO members (name, age, sex) VALUES ($1, $2, $3)", m.Name, m.Age, m.Sex)
	return err
}
