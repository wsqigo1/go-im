package wuid

import (
	"database/sql"
	"fmt"
	"github.com/edwingeng/wuid/mysql/wuid"
)

var w *wuid.WUID

func Init(dsn string) {
	// 也可以基于 redis
	newDB := func() (*sql.DB, bool, error) {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, false, err
		}
		return db, true, nil
	}

	w = wuid.NewWUID("default", nil)
	_ = w.LoadH28FromMysql(newDB, "wuid")
}

func GenUid(dsn string) string {
	if w == nil {
		Init(dsn)
	}
	// 16位宽度的十六进制字符串，并在前面添加0x前缀，不足16位的部分使用零填充
	return fmt.Sprintf("%#016x", w.Next())
}
