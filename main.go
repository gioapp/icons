package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioapp/icons/controller"
	"github.com/gioapp/icons/model"
	"image/color"
	"log"
	"sort"
)

type (
	D = layout.Dimensions
	C = layout.Context
)

var (
	scr = controller.NewScreen()
)

func main() {

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(800)),
			app.Title("ParallelCoin"),
		)

		if err := loop(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	var o op.Ops
	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				c := layout.NewContext(&o, e)
				appMain(c, th)
				e.Frame(c.Ops)
			}
		}
	}
}
func appMain(gtx layout.Context, th *material.Theme) D {
	DrawRectangle(gtx, gtx.Constraints.Max.X, gtx.Constraints.Max.Y, HexARGB(scr.BgColor))
	return layout.Flex{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return material.Body1(th, "Background").Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						renderColors(gtx, th, scr.BgColorLight, scr, "bg", scr.Colors["light"]),
						renderColors(gtx, th, scr.BgColorYellow, scr, "bg", scr.Colors["yellow"]),
						renderColors(gtx, th, scr.BgColorPurple, scr, "bg", scr.Colors["purple"]),
						renderColors(gtx, th, scr.BgColorRed, scr, "bg", scr.Colors["red"]),
						renderColors(gtx, th, scr.BgColorBlue, scr, "bg", scr.Colors["blue"]),
						renderColors(gtx, th, scr.BgColorDark, scr, "bg", scr.Colors["dark"]),
						renderColors(gtx, th, scr.BgColorOrange, scr, "bg", scr.Colors["orange"]),
						renderColors(gtx, th, scr.BgColorGreen, scr, "bg", scr.Colors["green"]))
				}),
				layout.Rigid(func(gtx C) D {
					return material.Body1(th, "Accent").Layout(gtx)
				}),
				layout.Rigid(func(gtx C) D {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						renderColors(gtx, th, scr.AccentColorLight, scr, "accent", scr.Colors["light"]),
						renderColors(gtx, th, scr.AccentColorYellow, scr, "accent", scr.Colors["yellow"]),
						renderColors(gtx, th, scr.AccentColorPurple, scr, "accent", scr.Colors["purple"]),
						renderColors(gtx, th, scr.AccentColorRed, scr, "accent", scr.Colors["red"]),
						renderColors(gtx, th, scr.AccentColorBlue, scr, "accent", scr.Colors["blue"]),
						renderColors(gtx, th, scr.AccentColorDark, scr, "accent", scr.Colors["dark"]),
						renderColors(gtx, th, scr.AccentColorOrange, scr, "accent", scr.Colors["orange"]),
						renderColors(gtx, th, scr.AccentColorGreen, scr, "accent", scr.Colors["green"]))
				}),
				layout.Rigid(func(gtx C) D {
					DrawRectangle(gtx, 160, gtx.Constraints.Max.Y, HexARGB(scr.AccentColor))
					return scr.NavList.Layout(gtx, len(scr.GroupsIco), func(gtx C, i int) D {
						return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx C) D {
							return material.RadioButton(th, scr.NavButtonsGroup, scr.GroupsIco[i], scr.GroupsIco[i]).Layout(gtx)
						})
					})
				}))
		}),
		layout.Flexed(1, func(gtx C) D {
			icoS := make([]string, 0, len(scr.IcoBank[scr.NavButtonsGroup.Value]))
			for k := range scr.IcoBank[scr.NavButtonsGroup.Value] {
				icoS = append(icoS, k)
			}
			sort.Strings(icoS)
			return scr.List.Layout(gtx, len(scr.IcoBank[scr.NavButtonsGroup.Value]), func(gtx C, i int) D {
				return layout.UniformInset(unit.Dp(8)).Layout(gtx, renderIcon(gtx, th, scr.IcoBank[scr.NavButtonsGroup.Value][icoS[i]], scr, icoS[i]))
			})
		}))
}

func renderIcon(gtx layout.Context, th *material.Theme, icon *widget.Icon, scr *model.Screen, iconLabel string) func(gtx C) D {
	return func(gtx C) D {
		return layout.Flex{
			Axis:    layout.Vertical,
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {

				return layout.Flex{
					Spacing: layout.SpaceBetween,
				}.Layout(gtx,
					layout.Rigid(func(gtx C) D {
						icon.Color = HexARGB(scr.AccentColor)
						return icon.Layout(gtx, unit.Dp(float32(scr.IconSize)))
					}),
					layout.Rigid(func(gtx C) D {
						return material.Label(th, th.TextSize.Scale(8.0/float32(scr.TextSize)), iconLabel).Layout(gtx)
					}),
				)
			}))
	}
}

func renderColors(gtx C, th *material.Theme, controllerButton *widget.Clickable, scr *model.Screen, part, color string) layout.FlexChild {
	return layout.Rigid(func(gtx C) D {
		return layout.Flex{
			Spacing: layout.SpaceBetween,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				var linkButton material.ButtonStyle
				linkButton = material.Button(th, controllerButton, "")
				linkButton.Background = HexARGB(color)
				for controllerButton.Clicked() {
					switch part {
					case "bg":
						scr.BgColor = color
					case "accent":
						scr.AccentColor = color
					}
				}
				return linkButton.Layout(gtx)
			}),
		)
	})
}

func DrawRectangle(gtx layout.Context, w, h int, color color.RGBA) {
	square := f32.Rectangle{
		Max: f32.Point{
			X: float32(w),
			Y: float32(h),
		},
	}
	paint.ColorOp{Color: color}.Add(gtx.Ops)
	paint.PaintOp{Rect: square}.Add(gtx.Ops)
	//layout.Dimensions{Size: image.Point{X: w, Y: h}}
}

func HexARGB(s string) (c color.RGBA) {
	_, _ = fmt.Sscanf(s, "%02x%02x%02x%02x", &c.A, &c.R, &c.G, &c.B)
	return
}
