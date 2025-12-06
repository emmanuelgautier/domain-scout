# Domain Scout CLI

<p align="left">
    <a href="https://github.com/emmanuelgautier/domain-scout/actions/workflows/ci.yml"><img src="https://github.com/emmanuelgautier/domain-scout/actions/workflows/ci.yml/badge.svg?branch=main&event=push" alt="CI Tasks for domain-scout"></a>
    <a href="https://goreportcard.com/report/github.com/emmanuelgautier/domain-scout"><img src="https://goreportcard.com/badge/github.com/emmanuelgautier/domain-scout" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/emmanuelgautier/domain-scout"><img src="https://pkg.go.dev/badge/www.github.com/emmanuelgautier/domain-scout" alt="PkgGoDev"></a>
</p>

This is a command-line tool for scanning domains and subdomains. It uses the [domain-scout](https://github.com/emmanuelgautier/domain-scout) library to check if subdomains are available.

## Usage

1. Clone or download this repository:

```bash
git clone https://github.com/emmanuelgautier/domain-scout.git
cd domain-scout
```

2. Build the CLI tool:

```bash
go build -o domain-scout
```

3. Run the CLI tool with the `--help` flag to see the available commands:

```bash
./domain-scout --help
```

You should see output similar to the following:

```
Scan domains and subdomains

Usage:
  domain-scout [command]

Available Commands:
  completion          Generate the autocompletion script for the specified shell
  help                Help about any command
  subdomain-available Check if subdomains are available

Flags:
  -h, --help   help for domain-scout

Use "domain-scout [command] --help" for more information about a command.
```

4. Run the `domain-scout` subcommand:

```bash
echo "example.com subdomain1.example.com subdomain2.example.com" | ./domain-scout subdomain-available
```

## License

This Domain Scout CLI is open-source and available under the MIT License. Feel free to use it as a starting point for your own CLI applications. Contributions and improvements are welcome!

## Author

[Emmanuel Gautier](https://www.emmanuelgautier.com/)
