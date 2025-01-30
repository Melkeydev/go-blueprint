# Advanced Flag in Blueprint

The `--advanced` or `-a` flag in Blueprint serves as a switch to enable additional features during project creation. It is applied with the `create` command and unlocks the following features:

- **CI/CD Workflow Setup using GitHub Actions:**
Automates the setup of a CI/CD workflow using GitHub Actions.

- **Websocket Support:**
WebSocket endpoint that sends continuous data streams through the WS protocol.

- **Docker:**
Docker configuration for go project.

To utilize the `--advanced` flag, use the following command:

```bash
go-blueprint create -n <project_name> -b <selected_backend> -driver <selected_driver> -a
```

By including the `--advanced` flag, users can choose one or all of the advanced features. The flag enhances the simplicity of Blueprint while offering flexibility for users who require additional functionality. 

Non-Interactive Setup is also possible:

```bash
go-blueprint create -n my_project -b standard-library -d redis -a --feature docker --feature githubaction --feature websocket -g commit
```
