Release process for Go projects, providing extensive customization options through its configuration file, `.goreleaser.yml`. By default, it ensures dependency cleanliness, builds binaries for various platforms and architectures, facilitates pre-release creation, and organizes binary packaging into archives with naming schemes.

For comprehensive insights into customization possibilities, refer to the [GoReleaser documentation](https://goreleaser.com/customization/).

## Usage with Tags

To initiate release builds with GoReleaser, you need to follow these steps:

- **Tag Creation:**
  When your project is ready for a release, create a new tag in your Git repository. For example:
```bash
git tag v1.0.0
```

- **Tag Pushing:**
  Push the tag to the repository to trigger GoReleaser:
```bash
git push origin v1.0.0
```

Following these steps ensures proper tagging of your project, prompting GoReleaser to execute configured releases. This approach simplifies release management and automates artifact distribution.

## Go Test - Continuous Integration for Go Projects

The `go-test.yml` file defines a GitHub Actions workflow for continuous integration (CI) of Go projects within a GitHub repository.

## Workflow Steps

The job outlined in this workflow includes the following steps:

1. **Checkout:**
   Fetches the project's codebase from the repository.

2. **Go Setup:**
   Configures the Go environment with version 1.21.x.

3. **Build and Test:**
   Builds the project using `go build` and runs tests across all packages (`./...`) using `go test`. 

This workflow serves to automate the testing process of a Go project within a GitHub repository, ensuring code quality and reliability with each commit and pull request.
