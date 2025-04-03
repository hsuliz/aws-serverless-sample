package types

type (
	Book struct {
		ID        string `dynamodbav:"id" json:"id"`
		Title     string `dynamodbav:"title" json:"title"`
		Pages     int    `dynamodbav:"pages" json:"pages"`
		PagesDone int    `dynamodbav:"pages_done" json:"pages_done"`
		BookDone  bool   `json:"book_done"`
	}

	BookRange struct {
		Books []Book `json:"books"`
	}
)
