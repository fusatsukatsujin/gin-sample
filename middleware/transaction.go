package middleware

import (
	"database/sql"
	"log"
	"net/http"

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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "トランザクションの開始に失敗しました"})
			c.Abort()
			return
		}

		// トランザクションをコンテキストに保存
		c.Set("tx", tx)

		// パニックをキャッチしてロールバック
		defer func() {
			if r := recover(); r != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Printf("Rollback error during panic: %v", rollbackErr)
				}
				panic(r) // パニックを再スロー
			}
		}()

		c.Next()

		// レスポンスのステータスコードを確認
		if c.Writer.Status() >= 400 {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("Rollback error on status >= 400: %v", rollbackErr)
			}
		} else {
			if err := tx.Commit(); err != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					log.Printf("Rollback error after commit failure: %v", rollbackErr)
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": "トランザクションのコミットに失敗しました"})
			}
		}
	}
}
