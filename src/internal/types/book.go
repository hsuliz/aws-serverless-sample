package types

type (
	Book struct {
		Id        string `dynamodbav:"id" json:"id"`
		Title     string `dynamodbav:"title" json:"title"`
		Pages     int    `dynamodbav:"pages" json:"pages"`
		PagesDone int    `dynamodbav:"pages_done" json:"pages_done"`
	}

	BookRange struct {
		Books []Book `json:"books"`
	}
)
