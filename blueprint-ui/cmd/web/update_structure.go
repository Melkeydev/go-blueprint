package web

import (
	"blueprint-ui/cmd/web/components"
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
	// Set HTMX response headers
	w.Header().Set("HX-Trigger", "updateComplete")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	advancedOptions, ok := r.Form["advancedOptions"]
	if !ok {
		advancedOptions = []string{}
	}

	advancedFrontend, ok := r.Form["advancedFrontend"]
	if !ok {
		advancedFrontend = []string{}
	}

	options := components.OptionsStruct{
		ProjectName:      r.FormValue("projectName"),
		SelectedBackend:  r.FormValue("backend"),
		SelectedDB:       r.FormValue("database"),
		SelectGit:        r.FormValue("git"),
		SelectFrontend:   r.FormValue("frontend"),
		AdvancedFrontend: advancedFrontend,
		AdvancedOptions:  advancedOptions,
		LongFlags:        r.FormValue("longFlags") == "on",
	}

	commandStr := components.GetCommandString(options)
	err = components.FolderStructure(options, commandStr).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
