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

	// Generar el PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	pdf.Cell(0, 10, "Factura")
	pdf.Ln(10)

	pdf.Cell(0, 10, "Cliente: "+invoice.ClientName)
	pdf.Ln(10)

	pdf.Cell(0, 10, "Fecha: "+invoice.Date)
	pdf.Ln(10)

	pdf.Cell(0, 10, "Productos:")
	pdf.Ln(10)

	for _, product := range invoice.Products {
		pdf.Cell(0, 10, fmt.Sprintf("%s x %d = %.2f", product.Name, product.Quantity, product.Price*float64(product.Quantity)))
		pdf.Ln(10)
	}

	pdf.Cell(0, 10, "Total: "+fmt.Sprintf("%.2f", invoice.Total))
	pdf.Ln(10)

	// Guardar el PDF
	err := pdf.OutputFileAndClose("factura.pdf")
	if err != nil {
		fmt.Println(err)
	}
}