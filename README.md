<!-- markdownlint-disable MD033 MD041 MD013 -->

# SpinGO

SpinGO is a very basic library for making progress spinners which has no dependencies.

## Table of contents

- [SpinGO](#spingo)
  - [Table of contents](#table-of-contents)
  - [Demo](#demo)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Customization](#customization)
    - [ArtIndex](#artindex)
    - [Prefix](#prefix)
    - [Suffix](#suffix)
    - [SpinnerArt](#spinnerart)

## Demo

<image src="assets/demo.gif" width="300px"/>

**Note:** Framerate are not fully accurate. [Here](https://asciinema.org/a/429ZlUdM4IAhjHVnCQiacnrWb) is another, more accurate demo using asciinema.

## Installation

```bash
go get -u github.com/talwat/spingo/v1
```

## Usage

```go
spinner := spingo.Spinner{}

for i := 0; i < 100; i++ {
    spinner.DisplaySpinner()
    time.Sleep(40 * time.Millisecond)
}
```

## Customization

These are fields in the `Spinner` struct.

### ArtIndex

Index of the current spinner art, it's fine to not set this when creating a new spinner.

### Prefix

Prefix to display before the spinner.

### Suffix

Suffix to display after the spinner.

### SpinnerArt

Spinner art to display,
eg. `["/", "|", "\\", "-"]`
