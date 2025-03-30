package types

type (
	Book struct {
		Id   string `dynamodbav:"id" json:"id"`
		Name string `dynamodbav:"name" json:"name"`
	}

	BookRange struct {
		Books []Book `json:"books"`
	}
)
