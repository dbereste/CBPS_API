package invoice

import (
	"bytes"
	"fmt"
	template "html/template"
	"io/ioutil"
	"main/models/invoice"
	"net/http"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

func GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		username, _, _ := c.Request.BasicAuth()
		data := invoice.GetAll(username)

		c.JSON(http.StatusOK, gin.H{
			"Status":  "OK",
			"content": data,
		})
	}
}

func GetPaid() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, _ := c.Request.BasicAuth()
		data := invoice.GetPaid(username)

		c.JSON(http.StatusOK, gin.H{
			"Status":  "OK",
			"content": data,
		})
	}
}

func GetUnpaid() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, _ := c.Request.BasicAuth()
		data := invoice.GetUnpaid(username)

		c.JSON(http.StatusOK, gin.H{
			"Status":  "OK",
			"content": data,
		})
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, _ := c.Request.BasicAuth()
		var id string = c.Param("id")
		data := invoice.GetByID(id, username)
		// TODO: Add return fir no Content

		if data.Id != 0 {
			c.JSON(http.StatusOK, gin.H{
				"Status":  "OK",
				"content": data,
			})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}

	}
}

func GenerateHTML(writer http.ResponseWriter, tmpl string, user string, id string) string {
	var tmplt *template.Template

	tmplt, _ = template.ParseFiles(tmpl)
	data := invoice.GetByID(id, user)

	parsed := "/var/www/API/tmp/invoice" + id + ".html"

	f, _ := os.Create(parsed)
	err := tmplt.Execute(f, data)

	if err != nil {
		fmt.Print(err)
	}
	f.Close()

	return parsed
}

func Download() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, _, _ := c.Request.BasicAuth()
		var id string = c.Param("id")
		template := "/var/www/API/invoice_tmpl/template.html"

		parsed := GenerateHTML(c.Writer, template, username, id)

		html, err := ioutil.ReadFile(parsed)

		filename := "Invoice" + id + ".pdf"
		path := "/var/www/API/tmp/"

		var pdfFile string = path + filename
		if err != nil {
			fmt.Print("ReadFile:")
			fmt.Println(err)
		}

		page := wkhtmltopdf.NewPageReader(bytes.NewReader(html))
		page.NoBackground.Set(true)
		page.DisableExternalLinks.Set(false)

		pdfg, err := wkhtmltopdf.NewPDFGenerator()
		if err != nil {
			fmt.Print("NewPDFGenerator:")
			fmt.Println(err)
		}

		pdfg.AddPage(page)

		pdfg.Dpi.Set(350)

		pdfg.MarginBottom.Set(0)
		pdfg.MarginTop.Set(0)
		pdfg.MarginLeft.Set(0)
		pdfg.MarginRight.Set(0)

		err = pdfg.Create()
		if err != nil {
			fmt.Print("Create:")
			fmt.Println(err)
		}

		err = pdfg.WriteFile(pdfFile)

		if err != nil {
			fmt.Print("WriteFile:")
			fmt.Println(err)
		}

		c.File(path + filename)
	}
}

func Add() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{})
	}
}

func Edit() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{})
	}
}

func Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotImplemented, gin.H{})
	}
}
