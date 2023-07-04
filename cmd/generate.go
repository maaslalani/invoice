package cmd

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/maaslalani/invoice/assets"
	"github.com/maaslalani/invoice/utils"
	"github.com/signintech/gopdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	importPath     string
	output         string
	file           = utils.Invoice{}
	defaultInvoice = utils.DefaultInvoice()
)

func init() {
	rootCmd.AddCommand(generateCmd)
	viper.AutomaticEnv()

	generateCmd.Flags().StringVar(&importPath, "import", "", "Imported file (.json/.yaml)")
	generateCmd.Flags().StringVar(&file.Id, "id", time.Now().Format("20060102"), "ID")
	generateCmd.Flags().StringVar(&file.Title, "title", "INVOICE", "Title")

	generateCmd.Flags().Float64SliceVarP(&file.Rates, "rate", "r", defaultInvoice.Rates, "Rates")
	generateCmd.Flags().IntSliceVarP(&file.Quantities, "quantity", "q", defaultInvoice.Quantities, "Quantities")
	generateCmd.Flags().StringSliceVarP(&file.Items, "item", "i", defaultInvoice.Items, "Items")

	generateCmd.Flags().StringVarP(&file.Logo, "logo", "l", defaultInvoice.Logo, "Company logo")
	generateCmd.Flags().StringVarP(&file.From, "from", "f", defaultInvoice.From, "Issuing company")
	generateCmd.Flags().StringVarP(&file.To, "to", "t", defaultInvoice.To, "Recipient company")
	generateCmd.Flags().StringVar(&file.Date, "date", defaultInvoice.Date, "Date")
	generateCmd.Flags().StringVar(&file.Due, "due", defaultInvoice.Due, "Payment due date")

	generateCmd.Flags().Float64Var(&file.Tax, "tax", defaultInvoice.Tax, "Tax")
	generateCmd.Flags().Float64VarP(&file.Discount, "discount", "d", defaultInvoice.Discount, "Discount")
	generateCmd.Flags().StringVarP(&file.Currency, "currency", "c", defaultInvoice.Currency, "Currency")

	generateCmd.Flags().StringVarP(&file.Note, "note", "n", "", "Note")
	generateCmd.Flags().StringVarP(&output, "output", "o", "invoice.pdf", "Output file (.pdf)")

	flag.Parse()
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate an invoice",
	Long:  `Generate an invoice`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if importPath != "" {
			err := importData(importPath, &file, cmd.Flags())
			if err != nil {
				return err
			}
		}

		pdf := gopdf.GoPdf{}
		pdf.Start(gopdf.Config{
			PageSize: *gopdf.PageSizeA4,
		})
		pdf.SetMargins(40, 40, 40, 40)
		pdf.AddPage()
		err := pdf.AddTTFFontData("Inter", assets.InterFont)
		if err != nil {
			return err
		}

		err = pdf.AddTTFFontData("Inter-Bold", assets.InterBoldFont)
		if err != nil {
			return err
		}

		utils.WriteLogo(&pdf, file.Logo, file.From)
		utils.WriteTitle(&pdf, file.Title, file.Id, file.Date)
		utils.WriteBillTo(&pdf, file.To)
		utils.WriteHeaderRow(&pdf)
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

			utils.WriteRow(&pdf, file.Items[i], q, r, file.Currency)
			subtotal += float64(q) * r
		}
		if file.Note != "" {
			utils.WriteNotes(&pdf, file.Note)
		}
		utils.WriteTotals(&pdf, subtotal, subtotal*file.Tax, subtotal*file.Discount, file.Currency)
		if file.Due != "" {
			utils.WriteDueDate(&pdf, file.Due)
		}
		utils.WriteFooter(&pdf, file.Id)
		output = strings.TrimSuffix(output, ".pdf") + ".pdf"
		err = pdf.WritePdf(output)
		if err != nil {
			return err
		}

		fmt.Printf("Generated %s\n", output)

		return nil
	},
}
