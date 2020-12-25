package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type style struct {
	Left   int
	Center int
	Title  int
	Button int
}

func (xl *Excel) setStyle() {
	xl.Style = style{
		Left:   left(xl.File),
		Center: center(xl.File),
		Title:  title(xl.File),
		Button: button(xl.File),
	}
}

func left(xl *excelize.File) int {
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

func center(xl *excelize.File) int {
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

func title(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "gradient",
			Color:   []string{"#beff32", "#6ba543"},
			Shading: 2,
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}

func button(xl *excelize.File) int {
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
