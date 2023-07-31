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

func (s *Starter) AddContent(editor Editor) error {
    editor.Open()
}


func (s *Starter) SaveToDb(db *Db) error {
    return nil
}
