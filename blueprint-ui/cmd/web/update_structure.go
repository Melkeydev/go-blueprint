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

	// Get all advanced options
	advancedOptions, ok := r.Form["advancedOptions"]
	if !ok {
		advancedOptions = []string{}
	}

	// Handle mutual exclusivity and dependencies
	hasReact := false
	hasTailwind := false
	hasHtmx := false
	filteredOptions := []string{}

	// First pass to check what we have
	for _, opt := range advancedOptions {
		switch opt {
		case "react":
			hasReact = true
		case "htmx":
			hasHtmx = true
		case "tailwind":
			hasTailwind = true
		}
	}

	for _, opt := range advancedOptions {
		// Skip React if HTMX is selected or will be auto-selected
		if opt == "react" && (hasHtmx || (hasTailwind && !hasReact)) {
			continue
		}
		// Skip HTMX if React is selected
		if opt == "htmx" && hasReact {
			continue
		}
		filteredOptions = append(filteredOptions, opt)
	}

	// Auto-add HTMX if Tailwind is selected and React is not
	if hasTailwind && !hasReact && !hasHtmx {
		filteredOptions = append(filteredOptions, "htmx")
	}

	options := components.OptionsStruct{
		ProjectName:     r.FormValue("projectName"),
		SelectedBackend: r.FormValue("backend"),
		SelectedDB:      r.FormValue("database"),
		SelectGit:       r.FormValue("git"),
		AdvancedOptions: filteredOptions,
		ShortFlags:      r.FormValue("shortFlags") == "on",
	}

	commandStr := components.GetCommandString(options)
	err = components.FolderStructure(options, commandStr).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
