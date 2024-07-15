package validations

type CreateBookServiceInput struct {
	AuthorID    uint   `json:"author_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Genre       string `json:"genre" binding:"required"`
	NumPages    int    `json:"num_pages" binding:"required"`
	Languages   string `json:"languages" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

type UpdateBookServiceInput struct {
	AuthorID    uint   `json:"author_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	NumPages    int    `json:"num_pages"`
	Languages   string `json:"languages"`
	Stock       int    `json:"stock"`
	Price       int    `json:"price"`
}
