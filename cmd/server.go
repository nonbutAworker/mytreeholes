package cmd

import (
	"fmt"
	"html/template"
	"io"

	m "bricking.com/mytreeholes/management"
	md "bricking.com/mytreeholes/mongo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCommand)
}

var serverCommand = &cobra.Command{
	Use:   "server",
	Short: "start web server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Zz Start Echo Web Server v0.9 -- HEAD")
		initserver()
	},
}

func initserver() {
	// Echo instance
	e := echo.New()

	// Pre Middleware
	e.Pre(middleware.RemoveTrailingSlash())
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Static("/", "assets")
	e.Static("/admin", "assets")
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("").Delims("[[", "]]").ParseGlob("assets/*.html")),
	}

	e.Renderer = renderer

	s := initStorage()
	initApi(e, s)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func initStorage() md.Storage {
	return md.New()
}

func initApi(e *echo.Echo, s md.Storage) {
	m.New(e, s)
	//add other api package ...
}
