package pdfprint 

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func PdfPrint(path string) error {
	// Generate HTML from your project files
	html := generateHTMLFromProject(path)

	// Create PDF from HTML
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))

	// Create PDF
	if err := pdfg.Create(); err != nil {
		return err
	}

	// Write PDF to file
	if err := pdfg.WriteFile("./project_code.pdf"); err != nil {
		return err
	}

	fmt.Println("PDF created successfully")
	return nil
}

func generateHTMLFromProject(projectPath string) string {
	var html strings.Builder
	html.WriteString("<html><body><pre>")

	filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			content, err := io.ReadAll(file)
			if err != nil {
				return err
			}
			html.WriteString(fmt.Sprintf("<h2>%s</h2>", path))
			html.WriteString(string(content))
			html.WriteString("\n\n")
		}
		return nil
	})

	html.WriteString("</pre></body></html>")
	return html.String()
}
