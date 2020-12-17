package style

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

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
			Color:   []string{"#FFFFFF", "#E0EBF5"},
			Shading: 1,
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
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern", // gradient, pattern
			Pattern: 8,
			Color:   []string{"#ed2939"},
			// Shading: 1,
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}
