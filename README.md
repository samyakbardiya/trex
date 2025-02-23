# TReX

**TReX** is a TUI (Text User Interface) tool designed for writing, visualizing
and testing Regular Expressions (RegEx). Developed in Go, TReX features a
keyboard-driven interface that allows users to create and test RegEx patterns
effectively.

<!-- TODO: Add gif/video -->
<!-- @see https://github.com/icholy/ttygif -->
<!-- @see https://asciinema.org/ -->

<!-- toc -->

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Authors](#authors)

<!-- tocstop -->

## Features

- Written in Go, for efficient performance
- Capability to load external files for testing purposes
- Keyboard-driven interface, so that you don't need to use a mouse :)
- Mouse-support for those who still want to use it

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/)

## Installation

- **Clone the repository:**

  ```sh
  git clone https://github.com/samyakbardiya/trex.git
  cd trex
  ```

- **Build the application:**

  ```sh
  go install
  go build
  ```

- **Optionally, you can copy the binary to your PATH:**

  ```sh
  cp ./trex ~/.local/bin
  ```

- **Verify the installation:**

  ```sh
  ./trex --help
  ./trex --version
  ```

## Usage

- Simply open and start RegEx-ing

  ```sh
  ./trex
  ```

- Or, bring-your-own file

  ```sh
  ./trex file.txt
  ```

## Roadmap

- Editable text area instead of read-only view
- Pre-defined templates, for popular use-cases, like email-validation,
  url-validation, etc.
- Local history, similar to your shell history, which can be navigated via
  up & down arrow key
- Syntax highlighting for the RegEx input.

## Authors

- [Samyak Bardiya](https://links.samyakbardiya.dev)
- [Mital Sapkale](https://github.com/mitalrs)
