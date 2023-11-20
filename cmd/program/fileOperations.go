package program

type FileCreationOperation struct {
	Path            string
	FileName        string
	TemplateContent []byte
}

func (p *Project) GenerateFileCreationOperations() (map[string]FileCreationOperation, error) {
	operations := make(map[string]FileCreationOperation)

	mainTemplateContent := p.FrameworkMap[p.ProjectType].templater.Main()
	operations["main"] = FileCreationOperation{
		Path:            cmdApiPath,
		FileName:        "main.go",
		TemplateContent: mainTemplateContent,
	}

	return operations, nil
}
