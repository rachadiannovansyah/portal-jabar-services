package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateNewsCategoriesTable struct{}

func (m *CreateNewsCategoriesTable) GetName() string {
	return "CreateNewsCategoriesTable"
}

func (m *CreateNewsCategoriesTable) Up(con *sqlx.DB) {
	table := builder.NewTable("news_categories", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("title", 500).Nullable()
	table.String("description", 1000).Nullable()
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateNewsCategoriesTable) Down(con *sqlx.DB) {
	builder.DropTable("news_categories", con).MustExec()
}
