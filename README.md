<img width="1200" alt="invoice" src="https://github.com/maaslalani/invoice/assets/42545625/a719e38e-0721-4c4b-9224-b6bc58de0edf">

# Invoice

Generate invoices from the command line.

## Text-based User Interface

```bash
invoice generate
```

View the generated PDF at `invoice.pdf`, you can customize the output location
with `--output`.

```bash
open invoice.pdf
```

<img width="574" alt="invoice-preview" src="https://github.com/maaslalani/invoice/assets/42545625/e02d1e29-f0c7-431c-b183-c9beaab1ac44">

## Command Line Interface

```bash
# Generate an invoice from information.
invoice generate --title "Invoice" \
    --id 2 \
    --logo ./images/logo.png \
    --from "Dream, Inc." \
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

<img src="https://vhs.charm.sh/vhs-66CMd4UQuXkuxX9djHUnGX.gif" width="600" />

Save repeated information with environment variables:

```bash
export INVOICE_LOGO=/path/to/image.png
export INVOICE_FROM="Dream, Inc."
export INVOICE_TO="Imagine, Inc."
export INVOICE_TAX=0.13
export INVOICE_RATE=25
```


Generate new invoice:

```bash
invoice generate \
    --item "Yellow Rubber Duck" --quantity 5 \
    --item "Special Edition Plaid Rubber Duck" --quantity 1 \
    --notes "For debugging purposes."
    --output duck-invoice.pdf
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
