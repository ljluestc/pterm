package main

import "github.com/pterm/pterm"

func main() {
	// Restyle the default header on the fly. The padding adds horizontal
	// space on both sides of the text.
	pterm.DefaultHeader.WithPadding(15).WithBackgroundStyle(pterm.NewStyle(pterm.BgCyan)).WithTextStyle(pterm.NewStyle(pterm.FgBlack)).Println("This is a custom header!")

	// Alternatively, build a HeaderPrinter from scratch instead of deriving
	// from DefaultHeader.
	newHeader := pterm.HeaderPrinter{
		TextStyle:       pterm.NewStyle(pterm.FgBlack),
		BackgroundStyle: pterm.NewStyle(pterm.BgRed),
		Padding:         20,
	}

	newHeader.Println("This is a custom header!")
}
