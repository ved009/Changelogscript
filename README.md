# Changelogscript

This repository provides a simple Go program (`changelog_monitor.go`) that monitors a list of changelog URLs for updates.

## Usage

The program reads URLs from command line arguments, a file, or standard input and periodically checks if today's date appears on each page. Only URLs with a matching entry are printed.

### Build

```bash
go build changelog_monitor.go
```

### Run with a file

```bash
./changelog_monitor -file urls.txt
```

### Run with arguments

```bash
./changelog_monitor https://example.com/changelog https://example.org/news
```

### Use with pipes

```bash
cat urls.txt | ./changelog_monitor | another-command
```

The default check interval is one hour and can be changed with the `-interval` flag.
