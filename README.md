<p align="center"><img src="https://raw.githubusercontent.com/gobuffalo/buffalo/main/logo.svg" width="360"></p>

<p align="center">
<a href="https://pkg.go.dev/github.com/thegodwinproject/cli"><img src="https://pkg.go.dev/badge/github.com/thegodwinproject/cli" alt="PkgGoDev"></a>
<a href="https://github.com/thegodwinproject/cli/actions/workflows/standard-go-test.yml"><img src="https://github.com/thegodwinproject/cli/actions/workflows/standard-go-test.yml/badge.svg" alt="Standard Test" /></a>
<a href="https://goreportcard.com/report/github.com/thegodwinproject/cli"><img src="https://goreportcard.com/badge/github.com/thegodwinproject/cli" alt="Go Report Card" /></a>
</p>

# Buffalo CLI

This is the repo for the Buffalo CLI. The Buffalo CLI is a tool to develop, test and deploy your Buffalo applications.

## Installation

To install the Buffalo CLI you can run the following command:

```bash
go install github.com/thegodwinproject/cli/cmd/buffalo@latest
```

<!-- Installing the Buffalo CLI requires Go 1.16 or newer as it depends heavily on the embed package. Once you have ensured you installed Go 1.16 or newer,  -->

## Usage

Once installed, the Buffalo CLI can be used by invoking the `buffalo` command. To know more about the available commands, run the `buffalo help` command. or you can also get details on a specific command by running `buffalo help <command>`.
