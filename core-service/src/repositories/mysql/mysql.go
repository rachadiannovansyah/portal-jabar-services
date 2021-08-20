package mysql

import (
	"context"
	"database/sql"
	"fmt"
)

type mysqlRepository struct {
	Conn *sql.DB
}

func (m *mysqlRepository) count(ctx context.Context, query string) (total int64, err error) {

	err = m.Conn.QueryRow(query).Scan(&total)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return total, nil
}
