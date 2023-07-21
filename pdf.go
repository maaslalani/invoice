package main

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
	subtotalLabel = "Subtotal"
	discountLabel = "Discount"
	taxLabel      = "Tax"
	totalLabel    = "Total"
)

func writeLogo(pdf *gopdf.GoPdf, fonts Fonts, logo, from, details string) {
	if logo != "" {
		width, height := getImageDimension(logo)
		scaledWidth := 100.0
		scaledHeight := float64(height) * scaledWidth / float64(width)
		_ = pdf.Image(logo, pdf.GetX(), pdf.GetY(), &gopdf.Rect{W: scaledWidth, H: scaledHeight})
		pdf.Br(scaledHeight + 24)
	}
	_ = pdf.SetFont(fonts.Bold.Name, "", 15)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, from)

	if details != "" {
		pdf.Br(25)
		_ = pdf.SetFont(fonts.Regular.Name, "", 10)
		pdf.SetTextColor(95, 95, 95)
		lines := strings.Split(details, "\n")
		for i, line := range lines {
			_ = pdf.Cell(nil, line)
			if i == len(lines)-1 {
				continue
			}
			pdf.Br(11)
		}
	}
	pdf.Br(36)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX(), pdf.GetY(), 100, pdf.GetY())
	pdf.Br(36)
}

func writeTitle(pdf *gopdf.GoPdf, fonts Fonts, title, id, date string) {
	_ = pdf.SetFont(fonts.Bold.Name, "", 24)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.Cell(nil, title)
	pdf.Br(36)
	_ = pdf.SetFont(fonts.Regular.Name, "", 12)
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, "#")
	_ = pdf.Cell(nil, id)
	pdf.SetTextColor(150, 150, 150)
	_ = pdf.Cell(nil, "  ·  ")
	pdf.SetTextColor(100, 100, 100)
	_ = pdf.Cell(nil, date)
	pdf.Br(48)
}

func writeDueDate(pdf *gopdf.GoPdf, fonts Fonts, due string) {
	_ = pdf.SetFont(fonts.Regular.Name, "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, "Due Date")
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(11)
	pdf.SetX(amountColumnOffset - 15)
	_ = pdf.Cell(nil, due)
	pdf.Br(12)
}

func writeBillTo(pdf *gopdf.GoPdf, fonts Fonts, to string) {
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont(fonts.Regular.Name, "", 9)
	_ = pdf.Cell(nil, "BILL TO")
	pdf.Br(18)
	pdf.SetTextColor(75, 75, 75)
	_ = pdf.SetFont(fonts.Regular.Name, "", 15)
	_ = pdf.Cell(nil, to)
	pdf.Br(64)
}

func writeHeaderRow(pdf *gopdf.GoPdf, fonts Fonts) {
	_ = pdf.SetFont(fonts.Regular.Name, "", 9)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, "ITEM")
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, "QTY")
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, "RATE")
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, "AMOUNT")
	pdf.Br(24)
}

func writeNotes(pdf *gopdf.GoPdf, fonts Fonts, notes string) {
	pdf.SetY(650)

	_ = pdf.SetFont(fonts.Regular.Name, "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, "Notes")
	pdf.Br(18)
	_ = pdf.SetFont(fonts.Regular.Name, "", 8)
	pdf.SetTextColor(0, 0, 0)

	formattedNotes := strings.ReplaceAll(notes, `\n`, "\n")
	notesLines := strings.Split(formattedNotes, "\\n")

	for i := 0; i < len(notesLines); i++ {
		_ = pdf.Cell(nil, notesLines[i])
		pdf.Br(15)
	}

	pdf.Br(48)
}

func writeFooter(pdf *gopdf.GoPdf, fonts Fonts, id string) {
	pdf.SetY(800)

	_ = pdf.SetFont(fonts.Regular.Name, "", 10)
	pdf.SetTextColor(55, 55, 55)
	_ = pdf.Cell(nil, id)
	pdf.SetStrokeColor(225, 225, 225)
	pdf.Line(pdf.GetX()+10, pdf.GetY()+6, 550, pdf.GetY()+6)
	pdf.Br(48)
}

func writeRow(pdf *gopdf.GoPdf, fonts Fonts, item string, quantity int, rate float64) {
	_ = pdf.SetFont(fonts.Regular.Name, "", 11)
	pdf.SetTextColor(0, 0, 0)

	total := float64(quantity) * rate
	amount := strconv.FormatFloat(total, 'f', 2, 64)

	_ = pdf.Cell(nil, item)
	pdf.SetX(quantityColumnOffset)
	_ = pdf.Cell(nil, strconv.Itoa(quantity))
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[file.Currency]+strconv.FormatFloat(rate, 'f', 2, 64))
	pdf.SetX(amountColumnOffset)
	_ = pdf.Cell(nil, currencySymbols[file.Currency]+amount)
	pdf.Br(24)
}

func writeTotals(pdf *gopdf.GoPdf, fonts Fonts, subtotal float64, tax float64, discount float64) {
	pdf.SetY(650)

	writeTotal(pdf, fonts, subtotalLabel, subtotal)
	if tax > 0 {
		writeTotal(pdf, fonts, taxLabel, tax)
	}
	if discount > 0 {
		writeTotal(pdf, fonts, discountLabel, discount)
	}
	writeTotal(pdf, fonts, totalLabel, subtotal+tax-discount)
}

func writeTotal(pdf *gopdf.GoPdf, fonts Fonts, label string, total float64) {
	_ = pdf.SetFont(fonts.Regular.Name, "", 9)
	pdf.SetTextColor(75, 75, 75)
	pdf.SetX(rateColumnOffset)
	_ = pdf.Cell(nil, label)
	pdf.SetTextColor(0, 0, 0)
	_ = pdf.SetFontSize(12)
	pdf.SetX(amountColumnOffset - 15)
	if label == totalLabel {
		_ = pdf.SetFont(fonts.Bold.Name, "", 11.5)
	}
	_ = pdf.Cell(nil, currencySymbols[file.Currency]+strconv.FormatFloat(total, 'f', 2, 64))
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
