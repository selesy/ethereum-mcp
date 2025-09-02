# ethereum-mcp

![https://pkg.go.dev/github.com/selesy/ethereum-mcp](https://pkg.go.dev/badge/github.com/selesy/ethereum-mcp.svg)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/selesy/ethereum-mcp/pre-commit.yaml)
![https://goreportcard.com/report/github.com/selesy/ethereum-mcp](https://goreportcard.com/badge/github.com/selesy/ethereum-mcp)
![https://github.com/RichardLitt/standard-readme](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)
![GitHub License](https://img.shields.io/github/license/selesy/ethereum-mcp)
![GitHub Release](https://img.shields.io/github/v/release/selesy/ethereum-mcp)

The `ethereum-mcp` library provides Go packages to make creating Model Context Protocol servers for Ethereum block-chains easier.

## Install

Include this library in your project using the following command:

```sh
go get github.com/selesy/ethereum-mcp
```

## Usage

The `/pkg/scheam` package provides "augmented" JSONSchema files and their associated Go embeddings in a format that's suitable for use as a "Raw Schema" with the `mcp-go` library's [`NewToolWithRawSchema`](https://pkg.go.dev/github.com/mark3labs/mcp-go/mcp#NewToolWithRawSchema) constructor.

Full documentation for this library is available as [Go docs](https://pkg.go.dev/github.com/selesy/ethereum-mcp).

## Contributing

This project happily accepts both code and non-code contributions.  For code and documentation submissions, please open a PR.  Otherwise, see the links in the [Community](#Community) section below to submit bug reports or to discuss a proposed feature or use-case.

### Community

* Please report issues using https://github.com/selesy/ethereum-mcp/issues[GitHub Issues].
* PRs are happily considered when submitted to https://github.com/selesy/ethereum-mcp/pulls[GitHub Pull requests].
* Other questions or discussions can be submitted to https://github.com/selesy/ethereum-mcp/discussions[GitHub Discussions].

### Development

This project strives to follow the tenets below:

* Maintain minimal external dependencies:  If you have a feature that requires specific libraries, let's discuss whether a new Go module should be created in a sub-directory.
* [Conventional commits](https://www.conventionalcommits.org/en/v1.0.0/):  Releases numbers and release notes are generated from the commit message "title".
* Local checks: Install the pre-commit hooks as described below.  If a commit passes on your workstation, it should also pass when checked by GitHub actions.
* Code style: Generally follows the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md).  The `.golangci-lint` configuration and `pre-commit` hooks help support the ideas in this document.

The tools required to develop this project and to run the `pre-commit` checks are defined in the `.tool-versions` file.  If you're using `asdf`, simply run `asdf install`.  Otherwise, install the listed tools in the manner required by your operating system.  Once the required tools are installed, install the `pre-commit` hooks by running `pre-commit install --install-hooks`.  Test your environment by running `pre-commit run --all-files`.

## License

This project is distributed under the https://github.com/selesy/ethereum-mcp/blob/main/LICENSE[Apache 2.0 License].
