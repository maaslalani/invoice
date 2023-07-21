package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/signintech/gopdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed "Inter/Inter Variable/Inter.ttf"
var interFont []byte

//go:embed "Inter/Inter Hinted for Windows/Desktop/Inter-Bold.ttf"
var interBoldFont []byte

type Invoice struct {
	Id    string `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`

	Logo    string `json:"logo" yaml:"logo"`
	Details string `json:"details" yaml:"details"`
	From    string `json:"from" yaml:"from"`
	To      string `json:"to" yaml:"to"`
	Date    string `json:"date" yaml:"date"`
	Due     string `json:"due" yaml:"due"`

	Items      []string  `json:"items" yaml:"items"`
	Quantities []int     `json:"quantities" yaml:"quantities"`
	Rates      []float64 `json:"rates" yaml:"rates"`

	Tax      float64 `json:"tax" yaml:"tax"`
	Discount float64 `json:"discount" yaml:"discount"`
	Currency string  `json:"currency" yaml:"currency"`

	Note string `json:"note" yaml:"note"`

	Fonts Fonts `json:"fonts" yaml:"fonts"`
}

type Fonts struct {
	Regular Font `json:"regular" yaml:"regular"`
	Bold    Font `json:"bold" yaml:"bold"`
}

type Font struct {
	Name    string `json:"name" yaml:"name"`
	Path    string `json:"path" yaml:"path"`
	Content []byte `json:"-" yaml:"-"`
}

var defaultFonts = Fonts{
	Regular: Font{
		Name:    "Inter",
		Content: interFont,
	},
	Bold: Font{
		Name:    "Inter-Bold",
		Content: interBoldFont,
	},
}

func DefaultInvoice() Invoice {
	return Invoice{
		Id:         time.Now().Format("20060102"),
		Title:      "INVOICE",
		Rates:      []float64{25},
		Quantities: []int{2},
		Items:      []string{"Paper Cranes"},
		From:       "Project Folded, Inc.",
		To:         "Untitled Corporation, Inc.",
		Date:       time.Now().Format("Jan 02, 2006"),
		Due:        time.Now().AddDate(0, 0, 14).Format("Jan 02, 2006"),
		Tax:        0,
		Discount:   0,
		Currency:   "USD",
		Fonts:      defaultFonts,
	}
}

var (
	importPath     string
	output         string
	file           = Invoice{}
	defaultInvoice = DefaultInvoice()
)

func init() {
	viper.AutomaticEnv()

	generateCmd.Flags().StringVar(&importPath, "import", "", "Imported file (.json/.yaml)")
	generateCmd.Flags().StringVar(&file.Id, "id", time.Now().Format("20060102"), "ID")
	generateCmd.Flags().StringVar(&file.Title, "title", "INVOICE", "Title")

	generateCmd.Flags().Float64SliceVarP(&file.Rates, "rate", "r", defaultInvoice.Rates, "Rates")
	generateCmd.Flags().IntSliceVarP(&file.Quantities, "quantity", "q", defaultInvoice.Quantities, "Quantities")
	generateCmd.Flags().StringSliceVarP(&file.Items, "item", "i", defaultInvoice.Items, "Items")

	generateCmd.Flags().StringVarP(&file.Logo, "logo", "l", defaultInvoice.Logo, "Company logo")
	generateCmd.Flags().StringVar(&file.Details, "details", defaultInvoice.Logo, "Company details (address, phone number, etc)")
	generateCmd.Flags().StringVarP(&file.From, "from", "f", defaultInvoice.From, "Issuing company")
	generateCmd.Flags().StringVarP(&file.To, "to", "t", defaultInvoice.To, "Recipient company")
	generateCmd.Flags().StringVar(&file.Date, "date", defaultInvoice.Date, "Date")
	generateCmd.Flags().StringVar(&file.Due, "due", defaultInvoice.Due, "Payment due date")

	generateCmd.Flags().Float64Var(&file.Tax, "tax", defaultInvoice.Tax, "Tax")
	generateCmd.Flags().Float64VarP(&file.Discount, "discount", "d", defaultInvoice.Discount, "Discount")
	generateCmd.Flags().StringVarP(&file.Currency, "currency", "c", defaultInvoice.Currency, "Currency")


	generateCmd.Flags().StringVar(&file.Fonts.Regular.Name, "fonts.regular.name",  "", "Regular font name")
	generateCmd.Flags().StringVar(&file.Fonts.Regular.Path, "fonts.regular.path",  "", "Regular font path")
	generateCmd.Flags().StringVar(&file.Fonts.Bold.Name, "fonts.bold.name",  "", "Bold font name")
	generateCmd.Flags().StringVar(&file.Fonts.Bold.Path, "fonts.bold.path",  "", "Bold font path")

	generateCmd.Flags().StringVarP(&file.Note, "note", "n", "", "Note")
	generateCmd.Flags().StringVarP(&output, "output", "o", "invoice.pdf", "Output file (.pdf)")

	flag.Parse()
}

var rootCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Invoice generates invoices from the command line.",
	Long:  `Invoice generates invoices from the command line.`,
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate an invoice",
	Long:  `Generate an invoice`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if importPath != "" {
			err := importData(importPath, &file, cmd.Flags())
			if err != nil {
				return err
			}
		}

		file.Fonts.Regular = firstNonEmpty(file.Fonts.Regular, defaultFonts.Regular)
		file.Fonts.Bold = firstNonEmpty(file.Fonts.Bold, defaultFonts.Bold)
		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{
			PageSize: *gopdf.PageSizeA4,
		})
		pdf.SetMargins(40, 40, 40, 40)
		pdf.AddPage()

		if err := addFont(&pdf, file.Fonts.Regular); err != nil {
			return err
		}

		if err := addFont(&pdf, file.Fonts.Bold); err != nil {
			return err
		}

		writeLogo(&pdf, file.Fonts, file.Logo, file.From, file.Details)
		writeTitle(&pdf, file.Fonts, file.Title, file.Id, file.Date)
		writeBillTo(&pdf, file.Fonts, file.To)
		writeHeaderRow(&pdf, file.Fonts)
		subtotal := 0.0
		for i := range file.Items {
			q := 1
			if len(file.Quantities) > i {
				q = file.Quantities[i]
			}

			r := 0.0
			if len(file.Rates) > i {
				r = file.Rates[i]
			}

			writeRow(&pdf, file.Fonts, file.Items[i], q, r)
			subtotal += float64(q) * r
		}
		if file.Note != "" {
			writeNotes(&pdf, file.Fonts, file.Note)
		}
		writeTotals(&pdf, file.Fonts, subtotal, subtotal*file.Tax, subtotal*file.Discount)
		if file.Due != "" {
			writeDueDate(&pdf, file.Fonts, file.Due)
		}
		writeFooter(&pdf, file.Fonts, file.Id)
		output = strings.TrimSuffix(output, ".pdf") + ".pdf"
		if err := pdf.WritePdf(output); err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", output)

		return nil
	},
}

func main() {
	rootCmd.AddCommand(generateCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func addFont(pdf *gopdf.GoPdf, font Font) error {
	content := font.Content
	if len(content) == 0 {
		if font.Path == "" {
			return fmt.Errorf("missing font path")
		}
		bts, err := os.ReadFile(font.Path)
		if err != nil {
			return err
		}
		content = bts
	}
	fmt.Println("Using font", font.Name, len(content))
	return pdf.AddTTFFontData(font.Name, content)
}

func firstNonEmpty(f1, f2 Font) Font {
	if f1.Name != "" {
		return f1
	}
	return f2
}
