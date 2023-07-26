<img width="1200" alt="Invoice" src="https://github.com/maaslalani/nap/assets/42545625/16dae9d9-390c-49b6-aedd-3f882b17f57b">

# Invoice

Generate invoices from the command line.

## Command Line Interface

```bash
invoice generate --from "Dream, Inc." --to "Imagine, Inc." \
    --from-address "1 main road,newyark,567 999" \
    --to-address "234/4566 second-main street,washington,568 9899" \
    --item "Rubber Duck" --quantity 2 --rate 25 \
    --tax 0.13 --discount 0.15 \
    --note "For debugging purposes."
```

<img src="https://vhs.charm.sh/vhs-66CMd4UQuXkuxX9djHUnGX.gif" width="600" />

View the generated PDF at `invoice.pdf`, you can customize the output location
with `--output`.

```bash
open invoice.pdf
```


#### address 
To use an address which includes special separators like `,` 
enclose the text using a single quote `"` 


must be provided to cli as below 

`--from-address '"1865 N Charlotte st",Suite32,"Nothaman, SN 02317"`

`--to-address "135 Wehawken St.,\"Raleigh, NC 34217\""`



<img width="574" alt="Example invoice" src="https://github.com/maaslalani/nap/assets/42545625/13153de2-dfa1-41e6-a18e-4d3a5cea5b74">



### Environment

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
    --note "For debugging purposes." \
    --output duck-invoice.pdf
```

### Configuration File

Or, save repeated information with JSON / YAML:

```json
{
    "logo": "/path/to/image.png",
    "from": "Dream, Inc.",
    "from_address": "1 main road,newyark,567 8989",
    "to_address": "345/66 main-second street,washington,699-090",
    "to": "Imagine, Inc.",
    "tax": 0.13,
    "items": ["Yellow Rubber Duck", "Special Edition Plaid Rubber Duck"],
    "quantities": [5, 1],
    "rates": [25, 25],
}
```

Generate new invoice by importing the configuration file:

```bash
invoice generate --import path/to/data.json \
    --output duck-invoice.pdf
```

### Custom Templates

If you would like a custom invoice template for your business or company, please
reach out via:

* [Email](mailto:maas@lalani.dev)
* [Twitter](https://twitter.com/maaslalani)

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
