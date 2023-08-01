package starter

type Starter struct {
	title     string
	extension string
	content   string
}

func New(title, extension, content string) Starter {
	return Starter{
		title: title,
        extension: extension,
        content: content,
	}
}
