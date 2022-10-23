package main

/*This is a quick notetakingcli(qote), that lets you quickly take notes in the Terminal and saves them in a .txt file with the current time and date as name.
Features:
	- Save and quit with ctrl-w
	- Save files as .txt and with the current time and date as name
	- Press ctrl-c to just exit but not write
*/

import (
	"log"
	"os"
	"time"

	"github.com/jroimartin/gocui"
)

func main() {
	//create new gui
	g, err := gocui.NewGui(gocui.OutputNormal)
	//error handling
	if err != nil {
		log.Panicln(err)
	}
	//defer the close function ^^
	defer g.Close()

	//set a GUI manager and key bindings
	g.SetManagerFunc(Editorlayout)
	g.Mouse = true

	// key bindings with error handling
	if err := g.SetKeybinding("", gocui.KeyCtrlW, gocui.ModNone, save); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	//error handling and invoking mainloop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// layout with relative coordinates or GUI manager
func Editorlayout(g *gocui.Gui) error {
	//set relative cooridinates
	//get the terminal width(tw) and the terminal height (th)
	twidth, theight := g.Size()

	//set the view size
	if v, err := g.SetView("Editor", 1, 1, twidth-1, theight-1); err != nil {
		if err != gocui.ErrUnknownView {
			log.Panicln(err)
		}

		//set the current view
		_, err := g.SetCurrentView("Editor")
		if err != nil {
			return err
		}

		//enable Editor
		v.Editable = true

		//customization

		//enable/disable the frame
		v.Frame = false

		//set color
		v.FgColor = gocui.ColorGreen
		v.BgColor = gocui.ColorDefault

		//enable cursor
		g.Cursor = true
	}
	return nil
}

// callback function for keybindings
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func save(g *gocui.Gui, v *gocui.View) error {
	//get the current date and time
	currentTime := time.Now()

	//create a new .txt file with the name being the current date and time
	f, _ := os.Create(string(currentTime.Format("15:04:05-2006-01-02")) + ".txt")

	//write the Buffer to the file with error handling
	_, err := f.WriteString(v.Buffer())
	if err != nil {
		return err
	}

	//close the file with error handling
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	return gocui.ErrQuit
}
