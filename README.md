<!-- markdownlint-disable MD033 MD041 MD013 -->

# SpinGO

![GitHub Action](https://img.shields.io/github/workflow/status/talwat/spingo/test)
[![GitHub license](https://img.shields.io/github/license/talwat/spingo)](https://github.com/talwat/spingo)
[![Go Report Card](https://goreportcard.com/badge/github.com/talwat/spingo)](https://goreportcard.com/report/github.com/talwat/spingo)
[![GoDoc](https://godoc.org/github.com/talwat/spingo?status.svg)](https://godoc.org/github.com/talwat/spingo)

SpinGO is a very basic library for making progress spinners which has no dependencies.

## Table of contents

- [SpinGO](#spingo)
  - [Table of contents](#table-of-contents)
  - [Demo](#demo)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Customization](#customization)
    - [Prefix](#prefix)
    - [Suffix](#suffix)
    - [SpinnerArt](#spinnerart)

## Demo

<image src="assets/demo.gif" width="300px"/>

**Note:** Framerate are not fully accurate. [Here](https://asciinema.org/a/429ZlUdM4IAhjHVnCQiacnrWb) is another, more accurate demo using asciinema.

## Installation

```bash
go get -u github.com/talwat/spingo
```

## Usage

```go
spinner := spingo.Spinner{}

for i := 0; i < 100; i++ {
    spinner.DisplaySpinner()
    time.Sleep(40 * time.Millisecond)
}

spinner.End()
```

## Customization

These are fields in the `Spinner` struct.

### Prefix

Prefix to display before the spinner.

### Suffix

Suffix to display after the spinner.

### SpinnerArt

Spinner art to display,
eg. `["/", "|", "\\", "-"]`
