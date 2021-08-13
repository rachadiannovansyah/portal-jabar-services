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
		Reference("categories").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.String("title", 500).Nullable()
	table.String("content", 1000).Nullable()
	table.String("slug", 500).Nullable()
	table.String("imagePath", 500).Nullable()
	table.String("videoUrl", 500).Nullable()
	table.String("newsSource", 500).Nullable()
	table.String("showDate", 500).Nullable()
	table.String("endDate", 500).Nullable()
	table.String("published", 500).Default("unpublished")
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateNewsTable) Down(con *sqlx.DB) {
	builder.DropTable("news", con).MustExec()
}
