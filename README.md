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

### Run Commands

Display Help:

```bash
./gofs -h
./gofs --help
```

Output:

```bash
Usage: gofs [options] <pattern> [pathname]

Options:
  -h, --help       Show help message
  -v, --version    Show version of the utility

Positional Arguments:
  <pattern>       Pattern to search for (required)
  [pathname]      Directory to search (optional)
```

Display Version:

```bash
./gofs -v
./gofs --version
```

Output:

```bash
gofs version <version-number>
```

### Directory Structure

```bash
root/
├── dir1/
│   └── config.txt
├── dir2/
│   └── nested-dir/
│       └── data.json
├── unit-tests.go
├── testdata/
    ├── example.txt
    └── emptydir/
```

### Search

Search for files:

List all files and directories in the current directory:

```bash
gofs
```

Output

```yaml
dir1
dir2
testdata
unit-test.go
```

Search for a file in the current directory:

```bash
gofs config.txt
```

Output

```yaml
root/dir1/config.txt
```

Search for all files in a specified directory

```bash
gofs . root/dir1/
```

Output

```yaml
dir1/config.txt
```

Search for a file in a specified directory

```bash
gofs data.json dir2/
```

Output

```yaml
dir2/nested-dir/data.json
```

## License

This project is licensed under the MIT License. See the [LICENSE](#License "Goto License") file for details.

## Future Plans

- Add support for glob patterns (`--glob`).
- Introduce regex-based searches (`--regex`).
- Implement parallel execution for faster searches.
