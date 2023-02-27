<h1 align="center">
    <span style="font-family: monospace;">prompti</span>
</h1>
<h2 align="center">
Lightweight and customizable interactive prompt components for Go based CLI.
</h2>
<p align="center">
    <a href="https://github.com/sveltinio/prompti/blob/main/LICENSE" target="_blank">
        <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square&logo=none" alt="license" />
    </a>
     &nbsp;
     <a href="https://goreportcard.com/report/github.com/sveltinio/prompti/" target="_blank">
        <img src="https://goreportcard.com/badge/github.com/sveltinio/prompti" alt="go report card" />
    </a>
    &nbsp;
    <a href="https://pkg.go.dev/github.com/sveltinio/prompti/" target="_blank">
        <img src="https://pkg.go.dev/badge/github.com/sveltinio/prompti/.svg" alt="go reference" />
    </a>
</p>

**prompti** is a collection of TUI blocks initially created to be used by [sveltin](https://github.com/sveltinio/sveltin) and its use cases. We worked on it to allow customizations and make it available as standalone package. Here it is!

## Install

Install the latest version by using the `go get` command. Ensure to have **(Go v1.18 or higher)** installed on your machine.

```bash
go get github.com/sveltinio/prompti@latest
```

## Prompts

### Input

`input` is a terminal input prompt component supporting default values, validations (`type ValidateFunc func(string) error`), password input echo and _customizable_ styles.

It also provides ready to use validation functions for some common use cases:

- alphanumeric
- digits only
- integers
- floats
- email address
- url

#### Default

[Source Code](_examples/input/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/input/input-default.gif" alt="Input example">

### Default value and validation

<img src="https://statics.sveltin.io/github/prompti/input/input-initial-value.gif" alt="Input example with default value and validation">

#### Custom styles

[Source Code](_examples/input/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/input/input-styled.gif" alt="Input example with custom styles">

#### Other Examples

- [Email](_examples/input/email/main.go)
- [Password](_examples/input/password/main.go)

### Choose

`choose` is a customizable component for browsing a set of items.

#### Default

[Source Code](_examples/choose/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/choose/choose-default.gif" alt="Choose example">

#### Custom styles

[Source Code](_examples/choose/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/choose/choose-styled.gif" alt="Choose example with custom styles">

### Confirm

`confirm` is a customizable component to confirm an anction.

#### Default

[Source Code](_examples/confirm/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/confirm/confirm-default.gif" alt="Confirm example">

#### Custom styles

[Source Code](_examples/confirm/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/confirm/confirm-styled.gif" alt="Confirm example with custom styles">

### Toggle

`toggle` is a customizable component to confirm an anction. It works like `confirm` but it renders the options inline and not in a box.

#### Default

[Source Code](_examples/toggle/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/toggle/toggle-default.gif" alt="Toggle example">

#### Custom styles

[Source Code](_examples/toggle/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/toggle/toggle-styled.gif" alt="Toggle example with custom styles">

### ProgressBar

`progressbar` is a customizable component for progress meter.

#### Default

[Source Code](_examples/progressbar/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/progressbar/progressbar-default.gif" alt="ProgressBar example">

#### Styled

[Source Code](_examples/progressbar/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/progressbar/progressbar-styled.gif" alt="ProgressBar example">

## :free: License

prompti is free and open-source software licensed under the MIT License.

***
Made with [Charm](https://charm.sh).
