package model

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"github.com/gioapp/icons/ico"
)

type Screen struct {
	TextSize          float32
	IconSize          float32
	BgColor           string
	AccentColor       string
	GroupsIco         []string
	IcoBank           ico.Ico
	NavList           *layout.List
	List              *layout.List
	NavButtonsGroup   *widget.Enum
	BgColorLight      *widget.Clickable
	BgColorYellow     *widget.Clickable
	BgColorPurple     *widget.Clickable
	BgColorRed        *widget.Clickable
	BgColorBlue       *widget.Clickable
	BgColorDark       *widget.Clickable
	BgColorOrange     *widget.Clickable
	BgColorGreen      *widget.Clickable
	AccentColorLight  *widget.Clickable
	AccentColorYellow *widget.Clickable
	AccentColorPurple *widget.Clickable
	AccentColorRed    *widget.Clickable
	AccentColorBlue   *widget.Clickable
	AccentColorDark   *widget.Clickable
	AccentColorOrange *widget.Clickable
	AccentColorGreen  *widget.Clickable
	Colors            Colors
}
type Colors map[string]string
