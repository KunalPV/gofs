# gofs

A lightweight and fast CLI utility for searching files, implemented in Go.

## Features

- **Flexible Search**:
  - Search using patterns, glob, or regex.
  - Filter results by file type, extension, and case sensitivity.
  - Limit the depth of directory traversal.
- **Exclusion Support**:
  - Exclude files or directories using glob patterns.
- **Absolute Paths**:
  - Convert results to absolute paths with the `--abs-path` option.
- **User-Friendly CLI**:
  - `--help` to display usage information.
  - `--version` to display the current version of the tool.
- **Cross-Platform**:
  - Works on Linux, macOS, and Windows.

## Usage

### Run Commands

Display Help:

```bash
./gofs -h
./gofs --help
```

Output:

```yaml
Usage:
  gofs <pattern> [pathname] [flags]

Flags:
  -A, --absolute-path      Display resuults as absolute paths
  -S, --case-sensitive     Perform case-sensitive searches
  -x, --exclude string     Exclude files/directories matching a glob pattern
  -e, --extension string   Filter results by file extensions
  -t, --file-type string   Filter results by file type (file, dir, symlink)
  -g, --glob string        Search using a glob pattern (default: empty string)
  -h, --help               Display help for gofs
  -H, --hidden             Include hidden files in the search
  -L, --hyper-link         Display results as hyperlinks
  -I, --ignore             Include .*ignore files like .gitignore
  -l, --long-list          Display results in long list format
  -d, --max-depth int      Limit search to a specific directory depth (-1 for no limit) (default -1)
  -T, --max-threads int    Set the maximum number of parallel threads for traversal (default 8)
  -v, --version            Display the version of gofs
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

```yaml
root/
├── dir1/
│   └── config.txt
├── dir2/
│   └── nested-dir/
│       └── data.json
├── unit-test.go
├── testdata/
├── example.txt
└── emptydir/
```

### Search

Search for files:

List all files and directories in the current directory:

```bash
gofs .
```

Output

```yaml
dir1/config.txt
dir2/nested-dir/data.json
testdata/example.txt
testdata/empty-dir/
unit-test.go
```

Search for a file

```bash
gofs config.txt
```

Output

```yaml
dir1/config.txt
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

### Advanced Features

Search using regex pattern

```bash
gofs '.*\.json$'
```

Output

```yaml
dir2/nested-dir/data.json
```

Search using glob pattern

```bash
gofs -g '*.txt'
```

Output

```yaml
dir1/config.txt
testdata/example.txt
```

Search with exclusion

```bash
gofs -g '*.txt' -x testdata
```

Output

```yaml
dir1/config.txt
```

Limit the depth of directory traversal

```bash
gofs . -d 1
```

Output

```yaml
dir1
dir2
unit-tests.go
testdata
```

Search for specific file type (file, dir, symlink)

```bash
gofs -t dir
```

Output

```yaml
dir1
dir2
testdata
testdata/empty-dir
```

Search for case-sensitive file name

```bash
gofs Config.txt -S
```

Output

```yaml
No files found.
```

## License

This project is licensed under the MIT License. See the [LICENSE](#License "Goto License") file for details.
