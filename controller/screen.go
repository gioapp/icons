package controller

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/icons/ico"
	"github.com/gioapp/icons/model"
)

func Colors() model.Colors {
	return model.Colors{
		"light":  "ffcfcfcf",
		"yellow": "ffcfcf30",
		"purple": "ff803080",
		"red":    "ffcf3030",
		"green":  "ff30cf30",
		"blue":   "ff3080cf",
		"orange": "ffcf8030",
		"dark":   "ff303030",
	}
}

func NewScreen() *model.Screen {
	return &model.Screen{
		TextSize:    8.0,
		IconSize:    16,
		BgColor:     "ffcfcfcf",
		AccentColor: "ff303030",
		GroupsIco:   ico.NewIco().GroupsIco(),
		IcoBank:     ico.NewIco(),
		NavList: &layout.List{
			Axis:      layout.Vertical,
			Alignment: layout.Start,
		},
		List: &layout.List{
			Axis:      layout.Vertical,
			Alignment: layout.Start,
		},
		NavButtonsGroup:   new(widget.Enum),
		BgColorLight:      new(widget.Button),
		BgColorYellow:     new(widget.Button),
		BgColorPurple:     new(widget.Button),
		BgColorRed:        new(widget.Button),
		BgColorBlue:       new(widget.Button),
		BgColorDark:       new(widget.Button),
		BgColorOrange:     new(widget.Button),
		BgColorGreen:      new(widget.Button),
		AccentColorLight:  new(widget.Button),
		AccentColorYellow: new(widget.Button),
		AccentColorPurple: new(widget.Button),
		AccentColorRed:    new(widget.Button),
		AccentColorBlue:   new(widget.Button),
		AccentColorDark:   new(widget.Button),
		AccentColorOrange: new(widget.Button),
		AccentColorGreen:  new(widget.Button),
		Colors:            Colors(),
	}
}
