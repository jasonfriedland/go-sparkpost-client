# go-sparkpost-client

[![Lint](https://github.com/jasonfriedland/go-sparkpost-client/workflows/Lint/badge.svg)](https://github.com/jasonfriedland/go-sparkpost-client/actions/workflows/lint.yml)
[![Test](https://github.com/jasonfriedland/go-sparkpost-client/workflows/Test/badge.svg)](https://github.com/jasonfriedland/go-sparkpost-client/actions/workflows/test.yml)
[![Build](https://github.com/jasonfriedland/go-sparkpost-client/workflows/Build/badge.svg)](https://github.com/jasonfriedland/go-sparkpost-client/actions/workflows/build.yml)

A simple test SparkPost client.

## Usage

Lint, test and build the CLI binary:

    # All the things (clean lint test build):
    make all
    # Just build:
    make

Export SparkPost config:

    export SPARKPOST_API_KEY=<your API key>
    export SPARKPOST_API_URL=<your API URL>

Usage (an empty `--return-path` will use the `from` arg):

    usage: sp [<flags>] <from> <to>

    Flags:
        --help             Show context-sensitive help (also try --help-long and --help-man).
    -r, --return-path=""   Return path address.
    -s, --subject=SUBJECT  Email subject.

    Args:
    <from>  From email address.
    <to>    To email address.

Example:

    echo "Test email." | ./sp -s Test -r bounce@example.com from@example.com to@example.com
