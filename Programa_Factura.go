package main

import (
	"fmt"
	"time"
	"github.com/jung-kurt/gofpdf"
)

type Product struct {
	Name     string
	Price    float64
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

	date := time.Now().Format("02/01/2006 15:04:05")

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

	fmt.Println("Nombre: ", clientName)
	fmt.Println("Fecha: ", date)
	fmt.Println("Productos:")
	for _, product := range invoice.Products {
		fmt.Printf("%s x %d = %.2f\n", product.Name, product.Quantity, product.Price*float64(product.Quantity))
	}
	fmt.Printf("Total a pagar: %.2f\n", total)

	// Generar el PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Factura")
	pdf.Ln(12)

	// Datos del cliente
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Cliente: "+invoice.ClientName)
	pdf.Ln(8)

	// Fecha
	pdf.Cell(0, 10, "Fecha: "+invoice.Date)
	pdf.Ln(12)

	// Título de tabla
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(40, 10, "Producto", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Cantidad", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Precio Unitario", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Total", "1", 0, "C", false, 0, "")
	pdf.Ln(-1)

	// Rellenar la tabla con productos
	pdf.SetFont("Arial", "", 12)
	for _, product := range invoice.Products {
		pdf.CellFormat(40, 10, product.Name, "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%d", product.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", product.Price), "1", 0, "C", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", product.Price*float64(product.Quantity)), "1", 0, "C", false, 0, "")
		pdf.Ln(-1)
	}

	// Total
	pdf.SetFont("Arial", "B", 12)
	pdf.Ln(10)
	pdf.Cell(0, 10, "Total a pagar: "+fmt.Sprintf("%.2f", invoice.Total))

	// Guardar el PDF
	err := pdf.OutputFileAndClose("factura.pdf")
	if err != nil {
		fmt.Println(err)
	}
}