package structs

type NotePayload struct {
	Title   string `dynamodbav:"title"`
	Content string `dynamodbav:"content"`
}

type Note struct {
	ID string `dynamodbav:"id"`
	NotePayload
}
