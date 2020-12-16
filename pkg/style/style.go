package style

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// Left ...
func Left(xl *excelize.File) int {
	styleID, err := xl.NewStyle(`{
		"alignment": {
			"horizontal": "left",
			"vertical": "center"
		}
	}`)
	if err != nil {
		return -1
	}
	return styleID
}

// Center ...
func Center(xl *excelize.File) int {
	styleID, err := xl.NewStyle(`{
		"alignment": {
			"horizontal": "center",
			"vertical": "center"
		}
	}`)
	if err != nil {
		return -1
	}
	return styleID
}
