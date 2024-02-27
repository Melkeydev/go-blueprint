# Advanced Flag in Blueprint

The `--advanced` flag in Blueprint serves as a switch to enable additional features during project creation. It is applied with the `create` command and unlocks the following features:

- **HTMX Support using Templ:**
Enables the integration of HTMX support for dynamic web pages using Templ.

- **CI/CD Workflow Setup using GitHub Actions:**
Automates the setup of a CI/CD workflow using GitHub Actions.

- **Websocket Support:**
WebSocket endpoint that sends continuous data streams through the WS protocol.

To utilize the `--advanced` flag, use the following command:

```bash
go-blueprint create --name <project_name> --framework <selected_framework> --driver <selected_driver> --advanced
```

By including the `--advanced` flag, users can choose one or all of the advanced features. The flag enhances the simplicity of Blueprint while offering flexibility for users who require additional functionality.

To recreate the project using the same configuration non-interactively, use the following command:
```bash
go-blueprint create --name my-project --framework chi --driver mysql --advanced true
```