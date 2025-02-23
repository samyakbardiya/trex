# TReX

**TReX** is a terminal-based tool for writing, visualizing, and testing Regular
Expressions. Designed for efficiency, it provides a keyboard-driven interface
for rapid feedback on your regex experiments—all within your terminal.

[![asciicast](https://asciinema.org/a/704948.svg)](https://asciinema.org/a/704948)

[![Go Report Card](https://goreportcard.com/badge/github.com/samyakbardiya/trex)](https://goreportcard.com/report/github.com/samyakbardiya/trex) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE) [![GitHub release](https://img.shields.io/github/v/release/samyakbardiya/trex)](https://github.com/samyakbardiya/trex/releases)

<!-- toc -->

- [Why TReX?](#why-trex)
- [Features](#features)
- [Installation](#installation)
  * [Prerequisites](#prerequisites)
  * [Steps](#steps)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [Implementation](#implementation)

<!-- tocstop -->

## Why TReX?

[![xkcd comic about Regular Expressions](https://imgs.xkcd.com/comics/regular_expressions.png)](https://xkcd.com/208)

Sometimes you just want to quickly test out a regex without switching between
multiple browser tabs or online tools. TReX lets you see how your regex
interacts with your text in real time—all within your terminal.

- **Quick feedback:** Validate and debug regex patterns instantly.
- **Integrated testing:** Load files and experiment with regex combinations.
- **Efficient workflow:** Stay in your terminal and keep your focus on writing code.

Btw, `T` is from **T**UI and `ReX` from **R**eg**Ex**, hence **_TReX_**. :t-rex: roar!

## Features

- **Written in Go:** Fast and portable.
- **External file loading:** Test regex patterns against real-world data.
- **Keyboard-driven interface:** Navigate without the need for a mouse.
- **Mouse support:** For users who prefer it or need it.

## Installation

### Prerequisites

Make sure you have the following installed:

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

  - **_Optionally_, you can copy the binary to your `PATH`:**

    ```sh
    cp ./trex ~/.local/bin
    ```

- **Verify the installation:**

  ```sh
  ./trex --version
  ```

## Usage

- **Start TReX:**

  ```sh
  trex
  ```

- **Load a file into TReX:**

  ```sh
  trex file.txt
  ```

- **Advanced usage:** Check out the help flag for more commands:

  ```sh
  trex --help
  ```

## Roadmap

- [ ] **Editable Text Area**: Replace the read-only view with an editable interface.
- [ ] **Local History**: Implement local history similar to shell history,
      navigable with arrow keys.
- [ ] **Syntax Highlighting**: Add syntax highlighting for the RegEx input.
- [ ] **Toggleable Flags**: Implement quick toggling for RegEx flags, such as:
  - `g` (global)
  - `m` (multi-line)
  - `i` (case-insensitive)
  - `U` (ungreedy)

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
For major changes, please open an issue first to discuss what you'd like to
change.

## Implementation

Developed in Go, **TReX** leverages:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) for building the TUI.
- [Cobra](https://github.com/spf13/cobra) for command-line functionality.
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) for styling.
- [Bubbles](https://github.com/charmbracelet/bubbles) for additional utilities

---

<p align="center">
  Made during <a href="https://fossunited.org/fosshack/2025">FOSS HACK 2025</a> in India :india:
</p>

<p align="center">
<sup>By <a href="https://links.samyakbardiya.dev">Samyak Bardiya</a> &amp; <a href="https://github.com/mitalrs">Mital Sapkale</a></sup>
</p>
