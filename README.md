# A utility to analyse linkerd check's output

[![Go Report Card](https://goreportcard.com/badge/github.com/nikhilsbhat/linkerd-checker)](https://goreportcard.com/report/github.com/nikhilsbhat/linkerd-checker)
[![shields](https://img.shields.io/badge/license-MIT-blue)](https://github.com/nikhilsbhat/linkerd-checker/blob/main/LICENSE)
[![shields](https://godoc.org/github.com/nikhilsbhat/linkerd-checker?status.svg)](https://godoc.org/github.com/nikhilsbhat/linkerd-checker)
[![shields](https://img.shields.io/github/v/tag/nikhilsbhat/linkerd-checker.svg)](https://github.com/nikhilsbhat/linkerd-checker/tags)
[![shields](https://img.shields.io/github/downloads/nikhilsbhat/linkerd-checker/total.svg)](https://github.com/nikhilsbhat/linkerd-checker/releases)

command-line utility for analyse linkerd check's [output](https://linkerd.io/2.14/reference/cli/check/).

## Introduction

Why use `linkerd-checker` when we already have `linkerd check`? </br> While the `check` command in `linkerd` serves its purpose, it currently lacks the functionality to run checks against specific categories.

What does this mean for us? If we wish to conduct checks after installing `linkerd` in our environment, we're obliged to run all checks without the ability to selectively ignore certain components or categories.

`linkerd` does provide a means to generate test output in JSON format, and `linkerd-checker` takes advantage of this by allowing users to filter the output based on chosen categories.

This enables users to run or fail tests only for selected components.

While this may not be beneficial to everyone, it can certainly aid certain groups that intend to test linkerd post-installation.

## Requirements

* Basic understanding of [linkerd](https://linkerd.io/) and running
  the [checks](https://linkerd.io/2.14/reference/cli/check/) using `linkerd`.

## Usage

The command `analyse` would help in analysing the json output of `linkerd` check command.

```shell
# Running the below command would analyse the entire output
linkerd2 check -o json | linkerd-checker analyse --all

# Running the below command would analyse output for linkerd components linkerd-smi and linkerd-multicluster only
linkerd2 check -o json | linkerd-checker analyse --category linkerd-smi --category linkerd-multicluster

# Analysing output of linkerd viz checks
linkerd2 viz check -o json | linkerd-checker analyse --all

# Analysing output of linkerd multicluster check
linkerd multicluster check -o json | linkerd-checker analyse --all
```

## Documentation

Updated documentation on all available commands and flags can be found [here](https://github.com/nikhilsbhat/linkerd-checker/blob/main/docs/doc/linkerd-checker.md).

## Installation

* Recommend installing released versions. Release binaries are available on the [releases](https://github.com/nikhilsbhat/linkerd-checker/releases) page.
* Can always build it locally by running `go build` against cloned repo.

#### Docker

Latest version of docker images are published to [ghcr.io](https://github.com/nikhilsbhat/linkerd-checker/pkgs/container/linkerd-checker), all available images can be
found there. </br>

```bash
docker pull ghcr.io/nikhilsbhat/linkerd-checker:latest
docker pull ghcr.io/nikhilsbhat/linkerd-checker:<github-release-tag>
```
