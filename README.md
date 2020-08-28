# Humanize Pipeline

Sort the keys in a Concourse pipeline YAML file to be more human-readable.

Heavily based on [humanize-manifest](https://github.com/cloudfoundry-community/humanize-manifest).

## Limitations

- Makes no attempt to validate the pipeline before or after reordering
- Does not preserve linebreaks between keys
- Tests do not cover all pipeline schema options (WIP)

## Installation

```sh
go get github.com/EngineerBetter/humanize-pipeline
```

## Development

Run locally with

```sh
go run *.go pipeline.yml
```

Build with

```sh
go build
```

Run tests with

```sh
ginkgo -r
```
