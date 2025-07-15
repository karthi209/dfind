# dfind - Under Development

A simple Go application to scan a folder recursively, compute SHA256 hashes for all files, and identify duplicate files based on their content.

## Usage

```go
go run main.go
```

## How it works

- Walks through all files in the specified folder.
- Computes the SHA256 hash for each file.
- Groups files with identical hashes (i.e., duplicate content).
- Prints out groups of duplicate files.

## Example output

```
Duplicate files found:
-----
/path/to/duplicate1.txt
/path/to/duplicate2.txt
```

> Errors encountered during file access or hashing are printed to the console.
