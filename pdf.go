package invoiceservice

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/signintech/gopdf"
)

const (
	quantityColumnOffset = 360
	rateColumnOffset     = 405
	amountColumnOffset   = 480
)

const (
	subtotalLabel = "Zwischensumme"  // subtotalLabel = "Subtotal"
	discountLabel = "Ermäßigung"     // discountLabel = "Discount"
	taxLabel      = "Inkl. 20% Ust." // taxLabel      = "Tax"
	totalLabel    = "Gesamtbetrag"   // totalLabel    = "Total"
)

func writeLogo(pdf *gopdf.GoPdf, logo string, from string) {
	if logo != "" {
		width, height := getImageDimension(logo)
		scaledWidth := 100.0
		scaledHeight := float64(height) * scaledWidth / float64(width)
		_ = pdf.Image(logo, pdf.GetX(), pdf.GetY(), &gopdf.Rect{W: scaledWidth, H: scaledHeight})
		pdf.Br(scaledHeight + 24)
	}
	_ = pdf.SetFont("Inter", "", 12)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, from)
	pdf.Br(36)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX(), pdf.GetY(), 100, pdf.GetY())
	pdf.Br(36)
}

func writeTitle(pdf *gopdf.GoPdf, title, id, date string, performanceDate string, localizationService LocalizationService) {
	_ = pdf.SetFont("Inter-Bold", "", 24)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Cell(nil, title)
	pdf.Br(36)
	_ = pdf.SetFont("Inter", "", 12)
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, "#")
	_ = pdf.Cell(nil, id)
	pdf.SetTextColor(150, 150, 150)
	_ = pdf.Cell(nil, "  ·  ")
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, date)
	pdf.SetX(rateColumnOffset - 15)
	_ = pdf.Cell(nil, fmt.Sprintf("%s %s", localizationService.serviceDate(), performanceDate))
	pdf.Br(48)
}

func writeDueDate(pdf *gopdf.GoPdf, due string, l LocalizationService) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset - 30)
	_ = pdf.Cell(nil, l.dueDate())
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(11)
	pdf.SetX(amountColumnOffset - 15)
	_ = pdf.Cell(nil, due)
	pdf.Br(12)
}

func writeBillTo(pdf *gopdf.GoPdf, to string, localizationService LocalizationService) {
	if to == "" {
		return
	}
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont("Inter", "", 9)
	_ = pdf.Cell(nil, localizationService.billTo())
	pdf.Br(18)
	_ = pdf.SetFont("Inter", "", 8)
	pdf.SetTextColor(0, 0, 0)

	formattedTo := strings.ReplaceAll(to, `\n`, "\n")
	toLines := strings.Split(formattedTo, "\n")

	for i := 0; i < len(toLines); i++ {
		_ = pdf.Cell(nil, toLines[i])
		pdf.Br(15)
	}

	pdf.Br(64)
}

func writeHeaderRow(pdf *gopdf.GoPdf, l LocalizationService) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, l.item())
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, l.qty())
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, l.rate())
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, l.amount())
	pdf.Br(24)
}

func writeComment(pdf *gopdf.GoPdf, comment string, l LocalizationService) {
	pdf.SetY(550)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, fmt.Sprintf("%s:", l.note()))
	pdf.Br(18)
	_ = pdf.SetFont("Inter", "", 8)
	pdf.SetTextColor(0, 0, 0)

	formattedNotes := strings.ReplaceAll(comment, `\n`, "\n")
	notesLines := strings.Split(formattedNotes, "\n")

	for i := 0; i < len(notesLines); i++ {
		_ = pdf.Cell(nil, notesLines[i])
		pdf.Br(15)
	}
}

func writeNotes(pdf *gopdf.GoPdf, notes string, l LocalizationService) {
	pdf.SetY(650)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, l.from())
	pdf.Br(18)
	_ = pdf.SetFont("Inter", "", 8)
	pdf.SetTextColor(0, 0, 0)

	formattedNotes := strings.ReplaceAll(notes, `\n`, "\n")
	notesLines := strings.Split(formattedNotes, "\n")

	for i := 0; i < len(notesLines); i++ {
		_ = pdf.Cell(nil, notesLines[i])
		pdf.Br(15)
	}

	pdf.Br(48)
}
func writeFooter(pdf *gopdf.GoPdf, id string) {
	pdf.SetY(800)

	_ = pdf.SetFont("Inter", "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, id)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX()+10, pdf.GetY()+6, 550, pdf.GetY()+6)
	pdf.Br(48)
}

func writeRow(pdf *gopdf.GoPdf, item string, quantity int, rate float64, currency string) {
	_ = pdf.SetFont("Inter", "", 11)
	pdf.SetTextColor(0, 0, 0)

	total := float64(quantity) * rate
	amount := strconv.FormatFloat(total, 'f', 2, 64)

	_ = pdf.Cell(nil, item)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, strconv.Itoa(quantity))
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[currency]+strconv.FormatFloat(rate, 'f', 2, 64))
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[currency]+amount)
	pdf.Br(24)
}

func writeTotals(pdf *gopdf.GoPdf, subtotal float64, tax float64, includeTax bool, discount float64, currency string, l LocalizationService) {
	pdf.SetY(650)

	writeTotal(pdf, l.subtotal(), subtotal, currency)
	if tax > 0 {
		writeTotal(pdf, l.taxLabelATInclVat(), tax, currency)
	}
	if discount > 0 {
		writeTotal(pdf, l.discount(), discount, currency)
	}
	var total float64 = subtotal - discount
	if !includeTax {
		total = total + tax
	}
	writeTotal(pdf, l.total(), total, currency)
}

func writeTotal(pdf *gopdf.GoPdf, label string, total float64, currency string) {
	_ = pdf.SetFont("Inter", "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset - 30)
	_ = pdf.Cell(nil, label)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(12)
	pdf.SetX(amountColumnOffset - 15)
	if label == totalLabel {
		_ = pdf.SetFont("Inter-Bold", "", 11.5)
	}
	_ = pdf.Cell(nil, currencySymbols[currency]+strconv.FormatFloat(total, 'f', 2, 64))
	pdf.Br(24)
}

func getImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	defer file.Close()

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}
