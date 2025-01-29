package web

import (
	"bluepront-ui/cmd/web/components"
	"net/http"
)

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func UpdateStructureHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	advancedOptions, ok := r.Form["advancedOptions"]
	if !ok {
		// Handle the case where no checkbox was checked
		advancedOptions = []string{}
	}

	if contains(advancedOptions, "tailwind") && !contains(advancedOptions, "htmx") {
		advancedOptions = append(advancedOptions, "htmx")
	}

	options := components.OptionsStruct{
		ProjectName:     r.FormValue("projectName"),
		SelectedBackend: r.FormValue("backend"),
		SelectedDB:      r.FormValue("database"),
		SelectGit:       r.FormValue("git"),
		AdvancedOptions: advancedOptions,
	}
	commandStr := components.GetCommandString(options)

	err = components.FolderStructure(options, commandStr).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
