# go-sparkpost-client

[![Build Status](https://travis-ci.org/jasonfriedland/go-sparkpost-client.svg?branch=master)](https://travis-ci.org/jasonfriedland/go-sparkpost-client)

A simple test SparkPost client.

## Usage

    make
    export SPARKPOST_API_KEY=<your API key>
    export SPARKPOST_API_URL=<your API URL>
    echo Testing | ./sp -s Test from@example.com to@example.com
