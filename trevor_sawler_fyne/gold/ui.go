package main

import "fyne.io/fyne/v2/container"

func (app *Config) makeUI() {
	// get the current price of gold - we do this with a separate function that returns
	// 3 objects. We could do it all here, but it'll be harder to test later.
	openPrice, currentPrice, priceChange := app.getPriceText()
	// put price information into a container
	priceContent := container.NewGridWithColumns(3,
		openPrice,
		currentPrice,
		priceChange,
	)

	app.PriceContainer = priceContent

	// get toolbar
	toolBar := app.getToolBar(app.MainWindow) // NOTE: here I'm starting to see the patterns. Even
	// with the separate files they are all in package 'main' and all connected with methods
	// off of *Config. This connects everything as all data from all packages are connecting
	// back to the storage struct and then I have access throughout the app.

	// add container to window
	finalContent := container.NewVBox(priceContent, toolBar)

	app.MainWindow.SetContent(finalContent)

}
