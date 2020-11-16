package main

import (
	"fmt"
  "strconv"
	"github.com/fatih/color"
	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
  "github.com/floscodes/golang-thousands"
)

// Heavily used singltons
// allow package wide availability
var root *sciter.Element
var rootSelectorErr error
var w *window.Window
var windowErr error

func init() {
    fmt.Println("Initializing GUI")

	// initializing window
	rect := sciter.NewRect(100, 100, 300, 220)
	w, windowErr = window.New(sciter.SW_TITLEBAR|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_GLASSY,
		rect)

    // handle window error
	if windowErr != nil {
		fmt.Println("Cannot create new window")
		return
	}

	// Loading main html file for app
	htloadErr := w.LoadFile("./simpleGuiCalc.html")
	if htloadErr != nil {
		fmt.Println("Cannot load html in window", htloadErr.Error())
		return
	}

	// Initializng selector
	root, rootSelectorErr = w.GetRootElement()
	if rootSelectorErr != nil {
		fmt.Println("Cannot find root element")
		return
	}

	// Set title of the application window
	w.SetTitle("My Simple Calc")
	fmt.Println("Initializing GUI complete")
}

// Primary GUI loop
func main() {
  fmt.Println("Adding application controls")

  addbutton, _ := root.SelectById("add")

	result, errres := root.SelectById("result")
  if (errres != nil) {
    fmt.Println("Failed to bind result")
  }


  addbutton.OnClick(func() {
		output := add()
    out1 := addCommas(output)
    result.SetHtml(out1, sciter.SIH_REPLACE_CONTENT)
	})

  fmt.Println("Adding application controls complete")
  fmt.Println("Showing window")
	w.Show()
	w.Run()
}

func addCommas(number int) string {
  commaNumberStr := strconv.Itoa(number)
  commaNumber := thousands.Separate(commaNumberStr, "en")
  return commaNumber;
}

// Primary action: add
func add() int {
    fmt.Println("Add action event")
	// Refreshing and fetching inputs()
	in1, errin1 := root.SelectById("input1")
	if errin1 != nil {
		color.Red("Failed to bind first input", errin1.Error())
	}
	in2, errin2 := root.SelectById("input2")
	if errin2 != nil {
		color.Red("Failed to bind second input", errin2.Error())
	}

	in1val, errv1 := in1.GetValue()
	//color.Green(in1val.String()) // prints the input to terminal

	if errv1 != nil {
		color.Red(errv1.Error())
	}
	in2val, errv2 := in2.GetValue()
	if errv2 != nil {
		color.Red(errv2.Error())
	}
	//color.Green(in2val.String())  // prints the input to terminal

    // clear input fields
	in1.SetValue(sciter.NewValue(""))
	in2.SetValue(sciter.NewValue(""))

    // set display labels
	di1, _ := root.SelectById("disp1")
	di2, _ := root.SelectById("disp2")

	di1.SetHtml(addCommas(in1val.Int()), sciter.SIH_REPLACE_CONTENT)
	di2.SetHtml(addCommas(in2val.Int()), sciter.SIH_REPLACE_CONTENT)

	return in1val.Int() + in2val.Int()
}

