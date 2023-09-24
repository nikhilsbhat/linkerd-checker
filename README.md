# A utility to analyse linkerd check's output

[![Go Report Card](https://goreportcard.com/badge/github.com/nikhilsbhat/linkerd-checker)](https://goreportcard.com/report/github.com/nikhilsbhat/linkerd-checker)
[![shields](https://img.shields.io/badge/license-MIT-blue)](https://github.com/nikhilsbhat/linkerd-checker/blob/main/LICENSE)
[![shields](https://godoc.org/github.com/nikhilsbhat/linkerd-checker?status.svg)](https://godoc.org/github.com/nikhilsbhat/linkerd-checker)
[![shields](https://img.shields.io/github/v/tag/nikhilsbhat/linkerd-checker.svg)](https://github.com/nikhilsbhat/linkerd-checker/tags)
[![shields](https://img.shields.io/github/downloads/nikhilsbhat/linkerd-checker/total.svg)](https://github.com/nikhilsbhat/linkerd-checker/releases)

command-line utility for analyse linkerd check's [output](https://linkerd.io/2.14/reference/cli/check/).

## Introduction

Why `linkerd-checker` when we already have `linkerd check`? </br> Yes, the `check` command in `linkerd` does the work. But running the checks against certain categories is not working at the moment.

What does it mean to us? If I want to run the checks post-installation of `linkerd` in our environment, we are forced to run all checks. </br>And we do not have a native way to ignore certain components or categories while running checks.

`linkerd` has a way to generate the output of tests as JSON, and the `linkerd-checker` leverages this and helps in filtering the output based on selected categories so that one can run or fail tests only for the selected components.

This might not help everyone, but it will definitely help some groups that are considering testing `linkerd` post-installation.

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
