# Invoice

<sub><sub>z</sub></sub><sub>z</sub>z

Generate invoices from the command line.

```bash
invoice generate --title "Invoice" \
    --id 2 \
    --logo ./images/logo.png \
    --from "LaLaLabs, Inc." \
    --to "Imagine, Inc." \
    --date "June 10, 2023" \
    --due "June 30, 2023" \
    --tax 0.13 \
    --discount 0.15 \
    --currency USD \
    --item "Rubber Duck" \
    --quantity 2 \
    --rate 25 \
    --notes "For debugging purposes."
```

## Installation

<!--

Use a package manager:

```bash
# macOS
brew install invoice

# Arch
yay -S invoice

# Nix
nix-env -iA nixpkgs.invoice
```

-->

Install with Go:

```sh
go install github.com/maaslalani/invoice@main
```

Or download a binary from the [releases](https://github.com/maaslalani/invoice/releases).


## License

[MIT](https://github.com/maaslalani/invoice/blob/master/LICENSE)

## Feedback

I'd love to hear your feedback on improving `invoice`.

Feel free to reach out via:
* [Email](mailto:maas@lalani.dev) 
* [Twitter](https://twitter.com/maaslalani)
* [GitHub issues](https://github.com/maaslalani/invoice/issues/new)

---

<sub><sub>z</sub></sub><sub>z</sub>z
