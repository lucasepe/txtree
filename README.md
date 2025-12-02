
# `txtree`

Visualize your data, branch by branch.

> A command-line tool to **visualize structured data as ASCII/Unicode trees**.

It works both with plain indented text and JSON input.

## Features

* Convert JSON or indented text into a clean tree view
* Multiple layout styles (right-aligned or top-down)
* Optional alphabetical sorting of JSON object keys
* Works with stdin or file inputs
* Unicode connectors for readable output

## Usage

```
txtree [flags] [file]
```

If no file is provided, txtree reads from [stdin].


## Examples

### Pipe JSON from [stdin]

```sh
cat api-response.json | txtree
```

### Read JSON from file

```sh
txtree api-response.json
```

### Parse indented text

```sh
txtree -i text notes.txt
```

### Use a different layout

```sh
txtree -l 3 data.json
```

## Flags

```
-h    Show help
        (default: false)

-i    Input format (auto, text, json)
       ↳ (default: auto)

-l    Layout style:
        0 = right-center
        1 = right-top
        2 = right-bottom
        3 = top-down
        (default: 0)

-s    Sort keys alphabetically (JSON only)
        (default: false)
```

Note: `-s` applies only when `-i=json`.
Sorting does not affect plain text mode.


## Support

All tools are completely free to use, with every feature fully unlocked and accessible.

If you find one or more of these tool helpful, please consider supporting its development with a donation.

Your contribution, no matter the amount, helps cover the time and effort dedicated to creating and maintaining these tools, ensuring they remain free and receive continuous improvements.

Every bit of support makes a meaningful difference and allows me to focus on building more tools that solve real-world challenges.

Thank you for your generosity and for being part of this journey!

[![Donate with PayPal](https://img.shields.io/badge/💸-Tip%20me%20on%20PayPal-0070ba?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=FV575PVWGXZBY&source=url)

## How To Install

### Using the _install.sh_ script (macOS & linux)

Simply run the following command in your terminal:

```sh
curl -sL https://raw.githubusercontent.com/lucasepe/txtree/main/install.sh | bash
```

This script will:

- Detect your operating system and architecture
- Download the latest release binary
- Install it into _/usr/local/bin_ (requires sudo)
  - otherwise fallback to _$HOME/.local/bin_ 
- Make sure the install directory is in your _PATH_ environment variable


### Manually download the latest binaries from the [releases page](https://github.com/lucasepe/txtree/releases/latest):

- [macOS](https://github.com/lucasepe/txtree/releases/latest)
- [Windows](https://github.com/lucasepe/txtree/releases/latest)
- [Linux (arm64)](https://github.com/lucasepe/txtree/releases/latest)
- [Linux (amd64)](https://github.com/lucasepe/txtree/releases/latest)

Unpack the binary into any directory that is part of your _PATH_.

## If you have [Go](https://go.dev/dl/) installed

You can also install it using:

```bash
go install github.com/lucasepe/txtree@latest
```

Make sure your `$GOPATH/bin` is in your PATH to run `txtree` from anywhere.



[stdin]: https://en.wikipedia.org/wiki/Standard_streams