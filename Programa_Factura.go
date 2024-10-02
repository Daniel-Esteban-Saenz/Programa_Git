package main

import (
	"fmt"
	"time"
	"github.com/jung-kurt/gofpdf"
)

type Product struct {
	Name    string
	Price   float64
	Quantity int
}

type Invoice struct {
	ClientName string
	Date       string
	Products   []Product
	Total      float64
}

func main() {
	// Solicitar información al usuario
	var clientName string
	fmt.Print("Ingrese el nombre del cliente: ")
	fmt.Scanln(&clientName)

	
	date := time.Now().Format("17/10/2003")

	var products []Product
	for {
		var productName string
		fmt.Print("Ingrese el nombre del producto (o 'fin' para terminar): ")
		fmt.Scanln(&productName)
		if productName == "fin" {
			break
		}

		var price float64
		fmt.Print("Ingrese el precio unitario del producto: ")
		fmt.Scanln(&price)

		var quantity int
		fmt.Print("Ingrese la cantidad del producto: ")
		fmt.Scanln(&quantity)

		product := Product{Name: productName, Price: price, Quantity: quantity}
		products = append(products, product)
	}

	// Calcular el total
	var total float64
	for _, product := range products {
		total += product.Price * float64(product.Quantity)
	}

	// Crear la factura
	invoice := Invoice{ClientName: clientName, Date: date, Products: products, Total: total}


	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.SetLineWidth(0.5)
	pdf.SetLineJoinStyle("D")
	pdf.Cell(0, 10, "Nombre")
	pdf.Cell(0, 10, "Cantidad")
	pdf.Cell(0, 10, "Precio Unitario")
	pdf.CellFormat(0, 10, "Total", "", 0, "L", false, 0, "")
	pdf.Ln(10)

	for _, product := range invoice.Products {
		pdf.Cell(0, 10, fmt.Sprintf(product.Name, 1, 0, "L", false, 0))
		pdf.Cell(0, 10, fmt.Sprintf("%.2f", float64(product.Quantity)))
		pdf.CellFormat(0, 10, fmt.Sprintf("%.2f", product.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(0, 10, fmt.Sprintf("%.2f", product.Price*float64(product.Quantity)), "1", 0, "C", false, 0,"")
		pdf.Ln(10)
	}

	pdf.SetFont("Arial", "B", 12)
	totalString := fmt.Sprintf("Total: %.2f", invoice.Total)
	pdf.Cell(0, 10, totalString) 
	pdf.Ln(10)

	err := pdf.OutputFileAndClose("factura.pdf")
	if err != nil {
		fmt.Println(err)
	}
}