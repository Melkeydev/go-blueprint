# Advanced Flag in Blueprint

The `--advanced` flag in Blueprint serves as a switch to enable additional features during project creation. It is applied with the `create` command and unlocks the following features:

- **HTMX Support using Templ:**
Enables the integration of HTMX support for dynamic web pages using Templ.

- **CI/CD Workflow Setup using GitHub Actions:**
Automates the setup of a CI/CD workflow using GitHub Actions.

## Usage

To utilize the `--advanced` flag, use the following command:

```bash
go-blueprint create --name <project_name> --framework <selected_framework> --driver <selected_driver> --advanced
```

By including the `--advanced` flag, users can choose one or both of the advanced features, HTMX support and GitHub Actions for CI/CD, during the project creation process. The flag enhances the simplicity of Blueprint while offering flexibility for users who require additional functionality.

To recreate the project using the same configuration non-interactively, use the following command:
```bash
go-blueprint create --name my-project --framework chi --driver mysql --advanced true
```

## HTMX Testing and Templ Setup

After creating your project with HTMX support using the `--advanced` flag, you can test HTMX functionality on `localhost:port/web`.

- **Navigate to Project Directory:**
```bash
cd my-project
```

- **Install Templ CLI:**
```bash
go install github.com/a-h/templ/cmd/templ@latest
```
- **Generate Templ Function Files:**
```bash
templ generate
```
- **Start server:**
```bash
make run
```

## GoReleaser - Creating and Pushing Tags

To create and push tags for builds using GoReleaser, follow these steps:


- **Creating a Tag:**
When you're ready for a release, create a new tag in your Git repository. For example:
```bash
git tag v1.0.0
```

- **Pushing the Tag:**
Push the tag to the repository to trigger GoReleaser:
```bash
git push origin v1.0.0
```

By following these steps, you ensure that your project is properly tagged, triggering GoReleaser to create and publish releases as configured in the workflow. This approach simplifies the release process and automates the creation of distribution artifacts.
