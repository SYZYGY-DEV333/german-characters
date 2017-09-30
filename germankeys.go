package main

import (
	"github.com/atotto/clipboard"
	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		//buttons:
		A := ui.NewButton("Ä")
		a := ui.NewButton("ä")
		O := ui.NewButton("Ö")
		o := ui.NewButton("ö")
		U := ui.NewButton("Ü")
		u := ui.NewButton("ü")
		S := ui.NewButton("ẞ")
		s := ui.NewButton("ß")
		
		//layout:
		help := ui.NewLabel("Copies pressed key to clipboard")
		caps := ui.NewLabel("Capital")
		lwrc := ui.NewLabel("Lowercase")
		author := ui.NewLabel("Coded by Joshua Sobel")
		
		row0 := ui.NewHorizontalBox()
		row0.Append(caps, true)
		row0.Append(lwrc, true)
		row0.SetPadded(true)
		
		row1 := ui.NewHorizontalBox()
		row1.Append(A, true)
		row1.Append(a, true)
		row1.SetPadded(true)
		
		row2 := ui.NewHorizontalBox()
		row2.Append(O, true)
		row2.Append(o, true)
		row2.SetPadded(true)
		
		row3 := ui.NewHorizontalBox()
		row3.Append(U, true)
		row3.Append(u, true)
		row3.SetPadded(true)
		
		row4 := ui.NewHorizontalBox()
		row4.Append(S, true)
		row4.Append(s, true)
		row4.SetPadded(true)
		
		mainbox := ui.NewVerticalBox()
		mainbox.Append(help, true)
		mainbox.Append(row0, false)
		mainbox.Append(row1, true)
		mainbox.Append(row2, true)
		mainbox.Append(row3, true)
		mainbox.Append(row4, true)
		mainbox.Append(author, true)
		
		window := ui.NewWindow("German Characters", 128, 384, false)
		window.SetChild(mainbox)
		mainbox.SetPadded(true)
		
		//controls:
		A.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("Ä")
		})
		a.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("ä")
		})
		O.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("Ö")
		})
		o.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("ö")
		})
		U.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("Ü")
		})
		u.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("ü")
		})
		S.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("ẞ")
		})
		s.OnClicked(func(*ui.Button) {
			clipboard.WriteAll("ß")
		})
		
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
