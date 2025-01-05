# gofs

A lightweight and fast CLI utility for searching files, implemented in Go.

## Features

- **Search by Pattern**: Locate files matching a specific name or pattern.
- **Custom Directory**: Specify the directory to search in.
- **User-Friendly CLI**:
  - `--help` to display usage information.
  - `--version` to display the current version of the tool.
- **Cross-Platform**: Works on Linux, macOS, and Windows.

## Usage

### Build the Tool

To build the CLI:

```bash
go build -o gofs cmd/gofs.go
```

### Run Commands

Display Help:

```bash
./gofs -h
./gofs --help
```

Output:

```bash
Usage: gofs [options] <directory> <filename>

Options:
  -h, --help       Show help message
  -v, --version    Show version of the utility
  --dir            Directory to search (default: current directory)

Positional Arguments:
  <directory>      Directory to search (overrides --dir)
  <filename>       Pattern to search for (required)
```

Display Version:

```bash
./gofs -v
./gofs --version
```

Output:

```bash
gofs version 1.0.0
```

Search for files:

- Search for a file in the current directory:

```bash
./gofs example.txt
```

- Search for a file in a specific directory:

```bash
./gofs testdata example.txt
```

## Development

### Run Tests

To run the test suite:

```bash
go test ./... -v
```

### Directory Structure

```
gofs/
├── cmd/                     # CLI entry point
│   └── gofs.go
├── internal/                # Internal modules
│   └── search/              # File search logic
│       └── search.go
├── tests/                   # Unit tests
│   └── search_test.go
├── testdata/                # Sample test data
│   ├── example.txt
│   ├── dir1/
│   └── emptydir/
├── README.md                # Project documentation
└── go.mod                   # Go module file
```

## License

This project is licensed under the MIT License. See the [LICENSE](#License "Goto License") file for details.

## Future Plans

- Add support for glob patterns (`--glob`).
- Introduce regex-based searches (`--regex`).
- Implement parallel execution for faster searches.
