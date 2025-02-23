# TReX

**TReX** is a Text User Interface (TUI) tool for writing, visualizing, and
testing Regular Expressions (RegEx). This terminal-based application uses a
keyboard-driven interface to allow users to create and evaluate RegEx patterns
efficiently.

<!-- toc -->

- [Why TReX?](#why-trex)
- [Features](#features)
- [Installation](#installation)
  * [Prerequisites](#prerequisites)
  * [Steps](#steps)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Implementation](#implementation)
- [Authors](#authors)

<!-- tocstop -->

## Why TReX?

**T** from the TUI and **REX** from the **RegEx**, hence the TReX, :t-rex: roar!

Because sometimes you just want to quickly test out a regex without fussing
through multiple browser tabs and searching for the right online tool. TReX was
created to let you see exactly how your regex interacts with your string—in one
simple, elegant terminal application.

## Features

- Written in Go, for efficient performance
- Capability to load external files for testing purposes
- Keyboard-driven interface, so that you don't need to use a mouse :)
- Mouse-support for those who still want to use it

## Installation

### Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/)

### Steps

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

- To start TReX:

  ```sh
  trex
  ```

- To load a file into TReX:

  ```sh
  trex file.txt
  ```

## Roadmap

- Editable text area instead of read-only view
- Local history, similar to your shell history, which can be navigated via
  up & down arrow key
- Syntax highlighting for the RegEx input.

## Implementation

TReX is developed in **Go** and leverages the following libraries:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for building the TUI.
- [Cobra](https://github.com/spf13/cobra) for command-line functionality.
- [Lipgloss](https://github.com/charmbracelet/lipgloss) for styling the interface.
- [Bubbles](https://github.com/charmbracelet/bubbles) for additional TUI utilities

## Authors

- [Samyak Bardiya](https://links.samyakbardiya.dev)
- [Mital Sapkale](https://github.com/mitalrs)
