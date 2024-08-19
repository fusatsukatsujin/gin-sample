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

func GetMembers(db *sql.DB) ([]Member, error) {
	rows, err := db.Query("SELECT name, age, sex FROM members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		err := rows.Scan(&member.Name, &member.Age, &member.Sex)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func GetMemberById(db *sql.DB, id int) (*Member, error) {
	row := db.QueryRow("SELECT name, age, sex FROM members WHERE id = $1", id)
	var member Member
	err := row.Scan(&member.Name, &member.Age, &member.Sex)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *Member) AddMember(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO members (name, age, sex) VALUES ($1, $2, $3)", m.Name, m.Age, m.Sex)
	return err
}
