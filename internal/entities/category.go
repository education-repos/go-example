package entities

type Category struct {
	ID            int    `db:"id"`
	Title         string `db:"title"`
	ProductAmount int    `db:"productAmount"`
	ParentId      int    `db:"parentId"`
	CreatedAt     string `db:"createdAt"`
	UpdatedAt     string `db:"updatedAt"`
}
