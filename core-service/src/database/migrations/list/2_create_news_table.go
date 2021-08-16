package list

import (
	"github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateNewsTable struct{}

func (m *CreateNewsTable) GetName() string {
	return "CreateNewsTable"
}

func (m *CreateNewsTable) Up(con *sqlx.DB) {
	table := builder.NewTable("news", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.Column("newsCategoryId").Type("int unsigned")
	table.ForeignKey("newsCategoryId").
		Reference("news_categories").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.String("title", 80)
	table.Column("content").Type("text")
	table.String("slug", 100).Nullable()
	table.String("image", 255).Nullable()
	table.String("videoUrl", 80).Nullable()
	table.String("source", 80).Nullable()
	table.Column("showDate").Type("datetime")
	table.Column("endDate").Type("datetime")
	table.String("status", 12).Default("published")
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateNewsTable) Down(con *sqlx.DB) {
	builder.DropTable("news", con).MustExec()
}
