# Structure CLI

A command-line tool to create a project directory structure based on a YAML file.

## Overview

This CLI allows users to generate a predefined directory structure and files for a project using a simple YAML configuration. Users can either create the structure in the current directory or specify a root directory name to create the structure within.

## Features

- Create a directory structure and files based on a YAML configuration file.
- Option to create the project structure in the current directory.
- Specify a project root name when creating the structure.

## Prerequisites

- Go (version 1.16 or later) installed on your machine.

## Installation

1. Clone the repository:
   ```bash
   git clone git@github.com:johnretsas/structure.git
   cd structure
   ```

2. Build the binary:
   ```bash
   go build -o structure
   ```

## Usage

Run the CLI with the following command:

```bash
./structure [flags]
```

### Flags

- `--here`  
  Create the project structure directly in the current directory, ignoring the project root.

- `--project-root <name>`  
  Specify the name for the project root directory. This flag is required if `--here` is not specified.

- `--structure <filename>`  
  Specify the YAML file containing the project structure (default is `structure.yaml`).

### Example

1. Create a project structure in the current directory:
   ```bash
   ./structure --here
   ```

2. Create a project structure with a specified project root:
   ```bash
   ./structure --project-root my_project
   ```

### YAML Structure

The YAML file should define the directory and file structure. Here is an example:

```yaml
project:
  main.go: file
  pkg:
    utils.go: file
    helpers.go: file
  cmd:
    myapp.go: file
  internal:
    db:
      connection.go: file
    api:
      handler.go: file
  README.md: file
  LICENSE: file
```

- Each item can either be a directory (denoted by a colon at the end of the name) or a file (indicated without a colon).
  
## Contributing

Contributions are welcome! Please submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.