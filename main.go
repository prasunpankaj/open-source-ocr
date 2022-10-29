package main

import (
    "bytes"
	"github.com/ledongthuc/pdf"
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("example/image/test1.png")
    //client.SetImage("example/pdf/test1.pdf")
	text, _ := client.Text()
	fmt.Println(text)
	// Hello, World!

    pdf.DebugOn = true
	content, err := readPdf("example/pdf/test1.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Printf("***********")
	fmt.Println(content)
	return

}

func readPdf(path string) (string, error) {
    f, r, err := pdf.Open(path)
	// remember close file
    defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
    b, err := r.GetPlainText()
    if err != nil {
        return "", err
    }
    buf.ReadFrom(b)
	return buf.String(), nil
}