package member

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewMember(t *testing.T) {
	member := NewMember("田中太郎", 30, Male)
	assert.Equal(t, "田中太郎", member.Name)
	assert.Equal(t, 30, member.Age)
	assert.Equal(t, Male, member.Sex)
}

func TestGetName(t *testing.T) {
	member := Member{Name: "山田花子", Age: 25, Sex: Female}
	assert.Equal(t, "山田花子", member.GetName())
}

func TestGetAge(t *testing.T) {
	member := Member{Name: "佐藤次郎", Age: 40, Sex: Male}
	assert.Equal(t, 40, member.GetAge())
}

func TestGetMembers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "age", "sex"}).
		AddRow("田中太郎", 30, "male").
		AddRow("山田花子", 25, "female")

	mock.ExpectQuery("SELECT name, age, sex FROM members").WillReturnRows(rows)

	members, err := GetMembers(db)
	assert.NoError(t, err)
	assert.Len(t, members, 2)
	assert.Equal(t, "田中太郎", members[0].Name)
	assert.Equal(t, 25, members[1].Age)
}

func TestGetMemberById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"name", "age", "sex"}).
		AddRow("佐藤次郎", 40, "male")

	mock.ExpectQuery("SELECT name, age, sex FROM members WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	member, err := GetMemberById(db, 1)
	assert.NoError(t, err)
	assert.NotNil(t, member)
	assert.Equal(t, "佐藤次郎", member.Name)
	assert.Equal(t, 40, member.Age)
	assert.Equal(t, Male, member.Sex)
}

func TestAddMember(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}
	defer db.Close()

	member := &Member{Name: "鈴木一郎", Age: 35, Sex: Male}

	mock.ExpectExec("INSERT INTO members \\(name, age, sex\\) VALUES \\(\\$1, \\$2, \\$3\\)").
		WithArgs(member.Name, member.Age, member.Sex).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = member.AddMember(db)
	assert.NoError(t, err)
}
