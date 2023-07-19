package banner

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func WriteBanner() {
	// Entry Header
	header := figure.NewFigure(" FinTech API", "", true)
	color.Red(header.String())
	fmt.Println(" ")
}
