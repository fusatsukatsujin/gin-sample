package member

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
