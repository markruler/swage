package xlsx

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type style struct {
	Left   int
	Center int
	Button int
	Title  int
	Column int
}

func (xl *Excel) setStyle() {
	xl.Style = style{
		Left:   left(xl.File),
		Center: center(xl.File),
		Button: button(xl.File),
		Title:  title(xl.File),
		Column: column(xl.File),
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

func button(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Size:      14,
			Italic:    true,
			Underline: "double",
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

func title(xl *excelize.File) int {
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
			Type:  "gradient",
			Color: []string{"#9CDC9D", "#3E9487"},
			// Color:   []string{"#beff32", "#6ba543"},
			Shading: 2,
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}

func column(xl *excelize.File) int {
	styleID, err := xl.NewStyle(&excelize.Style{
		// Font: &excelize.Font{
		// 	Family: "Arial",
		// },
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern", // gradient, pattern
			Pattern: 1,
			Color:   []string{"#d2ebd2"},
		},
	})
	if err != nil {
		return -1
	}
	return styleID
}
