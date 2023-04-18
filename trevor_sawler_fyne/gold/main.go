package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App      fyne.App
	Infolog  *log.Logger
	ErrorLog *log.Logger
}

var myApp Config // to store our configuration

func main() {
	// create a fyne application
	fyneApp := app.NewWithID("test.ladoblanco.ID.unique.preferences")
	myApp.App = fyneApp

	// create our loggers
	myApp.Infolog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.Infolog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// open a connection to the database - sqlite to embed the dbase in the application

	// create a database respository (simply pattern for interacting with things)

	// create and size a fyne window
	win := fyneApp.NewWindow("GoldWatcher")

	// show and run the application
	win.ShowAndRun()
}
