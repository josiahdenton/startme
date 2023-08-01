package starter

type Starter struct {
	Title     string
	Content   string
}

func New(title, extension, content string) Starter {
	return Starter{
		Title: title,
        Content: content,
	}
}
