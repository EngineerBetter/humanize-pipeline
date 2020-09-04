# Humanize Pipeline

Sort the keys in a Concourse pipeline YAML file to be more human-readable.

Heavily based on [humanize-manifest](https://github.com/cloudfoundry-community/humanize-manifest).

## TL;DR

```yaml
# cat fixtures/simple-pipeline-unordered.yml
jobs:
- plan:
  - trigger: true
    get: 2hours
  name: nothing
resources:
- source:
    interval: 2h
  type: time
  name: 2hours

# ./humanize-pipeline fixtures/simple-pipeline-unordered.yml
---
resources:
- name: 2hours
  type: time
  source:
    interval: 2h
jobs:
- name: nothing
  plan:
  - get: 2hours
    trigger: true
```

## Limitations

- Makes no attempt to validate the pipeline before or after reordering
- Does not preserve linebreaks between keys
- `in_parallel` must have nested `steps` key
- Tests do not cover all pipeline schema options (WIP)

## Installation

```sh
go get github.com/EngineerBetter/humanize-pipeline
```

## Development

Build with

```sh
go build
```

Run tests with

```sh
ginkgo -r
```
