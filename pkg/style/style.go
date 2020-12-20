package style

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var (
	IDLeft   int
	IDCenter int
	IDTitle  int
	IDButton int
)

func init() {
	xl := excelize.NewFile()
	IDLeft = Left(xl)     // 1
	IDCenter = Center(xl) // 2
	IDTitle = Title(xl)   // 3
	IDButton = Button(xl) // 4
}

// Left ...
func Left(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "left",
			Vertical:   "center",
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}

// Center ...
func Center(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}

// Title ...
func Title(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "gradient",
			Color:   []string{"#ffffff", "#38b832"},
			Shading: 5,
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}

// Button ...
func Button(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size: 14,
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern", // gradient, pattern
			Pattern: 8,
			Color:   []string{"#c0c0c0"},
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}
