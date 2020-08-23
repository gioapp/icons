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
		BgColorLight:      new(widget.Clickable),
		BgColorYellow:     new(widget.Clickable),
		BgColorPurple:     new(widget.Clickable),
		BgColorRed:        new(widget.Clickable),
		BgColorBlue:       new(widget.Clickable),
		BgColorDark:       new(widget.Clickable),
		BgColorOrange:     new(widget.Clickable),
		BgColorGreen:      new(widget.Clickable),
		AccentColorLight:  new(widget.Clickable),
		AccentColorYellow: new(widget.Clickable),
		AccentColorPurple: new(widget.Clickable),
		AccentColorRed:    new(widget.Clickable),
		AccentColorBlue:   new(widget.Clickable),
		AccentColorDark:   new(widget.Clickable),
		AccentColorOrange: new(widget.Clickable),
		AccentColorGreen:  new(widget.Clickable),
		Colors:            Colors(),
	}
}
