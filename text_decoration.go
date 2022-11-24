// Copyright ©2021-2022 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package unison

// TextDecoration holds the decorations that can be applied to text when drawn.
type TextDecoration struct {
	Font           Font
	Paint          *Paint
	Background     *Paint
	BaselineOffset float32
	Underline      bool
	StrikeThrough  bool
}

// Equivalent returns true if this TextDecoration is equivalent to the other.
func (d *TextDecoration) Equivalent(other *TextDecoration) bool {
	if d == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	return d.Underline == other.Underline && d.StrikeThrough == other.StrikeThrough &&
		d.BaselineOffset == other.BaselineOffset && d.Paint.Equivalent(other.Paint) &&
		d.Background.Equivalent(other.Background) && d.Font.Descriptor() == other.Font.Descriptor()
}

// Clone the TextDecoration.
func (d *TextDecoration) Clone() *TextDecoration {
	if d == nil {
		return nil
	}
	other := *d
	if other.Paint != nil {
		other.Paint = other.Paint.Clone()
	}
	if other.Background != nil {
		other.Background = other.Background.Clone()
	}
	return &other
}

// DrawText draws the given text using this TextDecoration.
func (d *TextDecoration) DrawText(canvas *Canvas, text string, x, y, width float32) {
	if d.Background != nil {
		canvas.DrawRect(NewRect(x, y-d.Font.Baseline(), width, d.Font.LineHeight()), d.Background)
	}
	y += d.BaselineOffset
	canvas.DrawSimpleString(text, x, y, d.Font, d.Paint)
	if d.Underline || d.StrikeThrough {
		paint := d.Paint.Clone()
		y++
		if d.StrikeThrough {
			yy := y + 0.5 - d.Font.Baseline()/2
			paint.SetStrokeWidth(1)
			canvas.DrawLine(x, yy, x+width, yy, paint)
		}
		if d.Underline {
			paint.SetStrokeWidth(1)
			canvas.DrawLine(x, y+1, x+width, y+1, paint)
		}
	}
}
