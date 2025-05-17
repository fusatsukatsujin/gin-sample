package controllers

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testServer struct {
	db     *sql.DB
	mock   sqlmock.Sqlmock
	router *gin.Engine
}

func setupTestServer(t *testing.T) *testServer {
	// データベースのモックを設定
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}

	// ルーターのセットアップ
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/api/members", func(c *gin.Context) {
		GetMembers(c, db)
	})

	return &testServer{
		db:     db,
		mock:   mock,
		router: router,
	}
}

func TestGetMembers(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.db.Close()

	// モックの期待値を設定
	rows := sqlmock.NewRows([]string{"name", "age", "sex"}).
		AddRow("テスト太郎", 20, "male").
		AddRow("テスト花子", 25, "female")

	ts.mock.ExpectQuery("SELECT name, age, sex FROM members").
		WillReturnRows(rows)

	// リクエストの実行
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/members", nil)
	ts.router.ServeHTTP(w, req)

	// レスポンスの検証
	assert.Equal(t, http.StatusOK, w.Code)

	// モックの期待値が満たされたか確認
	assert.NoError(t, ts.mock.ExpectationsWereMet(), "未実行のモックが存在します")

	expectedResponse := `{"members":[{"Name":"テスト太郎","Age":20,"Sex":"male"},{"Name":"テスト花子","Age":25,"Sex":"female"}]}`
	assert.Equal(t, expectedResponse, w.Body.String())
}

func TestGetMembersEmpty(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.db.Close()

	ts.mock.ExpectQuery("SELECT name, age, sex FROM members").
		WillReturnRows(sqlmock.NewRows([]string{"name", "age", "sex"}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/members", nil)
	ts.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"members\":[]}", w.Body.String())
}

func TestGetMembersDBError(t *testing.T) {
	ts := setupTestServer(t)
	defer ts.db.Close()

	// データベースエラーをシミュレート
	ts.mock.ExpectQuery("SELECT name, age, sex FROM members").
		WillReturnError(sql.ErrConnDone)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/members", nil)
	ts.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	// エラーレスポンスの内容も検証
	assert.Contains(t, w.Body.String(), "メンバーの取得に失敗しました")
}
