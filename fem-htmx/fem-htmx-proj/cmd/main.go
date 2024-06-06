package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

type Contact struct {
	Name  string
	Email string
}

func NewContact(name string, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

type Contacts = []Contact

type Data struct {
	Contacts Contacts
}

func (d *Data) HasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

func NewData() Data {
	return Data{
		Contacts: []Contact{
			NewContact("Sefat", "sefatanam@email.com"),
			NewContact("Anam", "anamsefat@email.com"),
		},
	}
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

type Page struct {
	Data Data
	Form FormData
}

func NewPage() Page {
	return Page{
		Data: NewData(),
		Form: NewFormData(),
	}
}
func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	// count := Count{Count: 0}
	// data := NewData()

	page := NewPage()
	e.Renderer = newTemplate()

	/* e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	}) */
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", page)
	})

	/* e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	}) */

	e.POST("/contacts", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")

		if page.Data.HasEmail(email) {
			formData := NewFormData()
			formData.Values["Name"] = name
			formData.Values["Email"] = email
			formData.Errors["email"] = "Email already exists"

			return c.Render(422, "form", formData)
		}

		page.Data.Contacts = append(page.Data.Contacts, NewContact(name, email))
		return c.Render(200, "display", page)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))
}
