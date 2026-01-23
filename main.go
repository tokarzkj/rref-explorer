package main

import (
	"os"
	"strconv"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.kristokarz.test.simple", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	builder := gtk.NewBuilderFromFile("./UI/main_ui.xml")

	window := builder.GetObject("MainWindow").Cast().(*gtk.Window)

	createMtxBtn := builder.GetObject("create_mtx_btn").Cast().(*gtk.Button)
	createMtxBtn.Connect("clicked", func() {
		rows_txt := builder.GetObject("mtx_rows").Cast().(*gtk.Text)
		row, _ := strconv.Atoi(rows_txt.Text())

		cols_txt := builder.GetObject("mtx_cols").Cast().(*gtk.Text)
		cols, _ := strconv.Atoi(cols_txt.Text())

		for i := 0; i < row; i++ {
			for j := 0; j < cols; j++ {

			}
		}

		//mtxGrid := builder.GetObject("grid_mtx").Cast().(*gtk.Grid)

	})

	app.AddWindow(window)

	window.SetVisible(true)
}
