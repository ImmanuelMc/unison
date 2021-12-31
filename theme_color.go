// Copyright ©2021 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package unison

import (
	"github.com/richardwilkes/toolbox/xmath/geom32"
)

var (
	_ ColorProvider = &ThemeColor{}
	_ Ink           = &ThemeColor{}
)

// Pre-defined theme colors.
var (
	BackgroundColor          = &ThemeColor{Light: RGB(238, 238, 238), Dark: RGB(50, 50, 50)}
	OnBackgroundColor        = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(221, 221, 221)}
	ContentColor             = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(32, 32, 32)}
	OnContentColor           = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(221, 221, 221)}
	BandingColor             = &ThemeColor{Light: RGB(235, 235, 220), Dark: RGB(42, 42, 42)}
	OnBandingColor           = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(221, 221, 221)}
	DividerColor             = &ThemeColor{Light: RGB(192, 192, 192), Dark: RGB(102, 102, 102)}
	HeaderColor              = &ThemeColor{Light: RGB(43, 43, 43), Dark: RGB(64, 64, 64)}
	OnHeaderColor            = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(192, 192, 192)}
	TabFocusedColor          = &ThemeColor{Light: RGB(224, 212, 175), Dark: RGB(68, 102, 0)}
	OnTabFocusedColor        = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(221, 221, 221)}
	TabCurrentColor          = &ThemeColor{Light: RGB(211, 207, 197), Dark: RGB(41, 61, 0)}
	OnTabCurrentColor        = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(221, 221, 221)}
	DropAreaColor            = &ThemeColor{Light: RGB(204, 0, 51), Dark: RGB(255, 0, 0)}
	EditableColor            = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(24, 24, 24)}
	OnEditableColor          = &ThemeColor{Light: RGB(0, 0, 160), Dark: RGB(100, 153, 153)}
	SelectionColor           = &ThemeColor{Light: RGB(0, 96, 160), Dark: RGB(0, 96, 160)}
	OnSelectionColor         = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(255, 255, 255)}
	InactiveSelectionColor   = &ThemeColor{Light: RGB(0, 64, 148), Dark: RGB(0, 64, 148)}
	OnInactiveSelectionColor = &ThemeColor{Light: RGB(228, 228, 228), Dark: RGB(228, 228, 228)}
	ScrollColor              = &ThemeColor{Light: ARGB(0.5, 192, 192, 192), Dark: ARGB(0.5, 128, 128, 128)}
	ScrollRolloverColor      = &ThemeColor{Light: RGB(192, 192, 192), Dark: RGB(128, 128, 128)}
	ScrollEdgeColor          = &ThemeColor{Light: RGB(128, 128, 128), Dark: RGB(160, 160, 160)}
	AccentColor              = &ThemeColor{Light: RGB(0, 102, 102), Dark: RGB(100, 153, 153)}
	ControlColor             = &ThemeColor{Light: RGB(248, 248, 255), Dark: RGB(64, 64, 64)}
	OnControlColor           = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(221, 221, 221)}
	ControlPressedColor      = &ThemeColor{Light: RGB(0, 96, 160), Dark: RGB(0, 96, 160)}
	OnControlPressedColor    = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(255, 255, 255)}
	ControlEdgeColor         = &ThemeColor{Light: RGB(96, 96, 96), Dark: RGB(96, 96, 96)}
	IconButtonColor          = &ThemeColor{Light: RGB(96, 96, 96), Dark: RGB(128, 128, 128)}
	IconButtonRolloverColor  = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(192, 192, 192)}
	IconButtonPressedColor   = &ThemeColor{Light: RGB(0, 96, 160), Dark: RGB(0, 96, 160)}
	TooltipColor             = &ThemeColor{Light: RGB(252, 252, 196), Dark: RGB(184, 204, 143)}
	OnTooltipColor           = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(32, 32, 32)}
	SearchListColor          = &ThemeColor{Light: RGB(224, 255, 255), Dark: RGB(0, 43, 43)}
	OnSearchListColor        = &ThemeColor{Light: RGB(0, 0, 0), Dark: RGB(204, 204, 204)}
	ErrorColor               = &ThemeColor{Light: RGB(192, 64, 64), Dark: RGB(115, 37, 37)}
	OnErrorColor             = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(221, 221, 221)}
	WarningColor             = &ThemeColor{Light: RGB(224, 128, 0), Dark: RGB(192, 96, 0)}
	OnWarningColor           = &ThemeColor{Light: RGB(255, 255, 255), Dark: RGB(221, 221, 221)}
)

// ThemeColor holds a pair of colors, one for light mode and one for dark mode.
type ThemeColor struct {
	Light Color
	Dark  Color
}

// GetColor returns the current color. Here to satisfy the ColorProvider interface.
func (t *ThemeColor) GetColor() Color {
	if IsDarkModeEnabled() {
		return t.Dark
	}
	return t.Light
}

// Paint returns a Paint for this ThemeColor. Here to satisfy the Ink interface.
func (t *ThemeColor) Paint(canvas *Canvas, rect geom32.Rect, style PaintStyle) *Paint {
	return t.GetColor().Paint(canvas, rect, style)
}
