package domain

type Starter struct {
	title     string
	extension string
	content   string
}

func New(title, extension string) Starter {
	return Starter{
		title: title,
        extension: extension,
	}
}
