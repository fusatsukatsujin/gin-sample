package middleware

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type TransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) HandleTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := tm.db.Begin()
		if err != nil {
			c.JSON(500, gin.H{"error": "トランザクションの開始に失敗しました"})
			c.Abort()
			return
		}

		// トランザクションをコンテキストに保存
		c.Set("tx", tx)

		// パニックをキャッチしてロールバック
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r) // パニックを再スロー
			}
		}()

		c.Next()

		// レスポンスのステータスコードを確認
		if c.Writer.Status() >= 400 {
			tx.Rollback()
		} else {
			if err := tx.Commit(); err != nil {
				tx.Rollback()
				c.JSON(500, gin.H{"error": "トランザクションのコミットに失敗しました"})
			}
		}
	}
}
