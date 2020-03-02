// This file is part of Gopher2600.
//
// Gopher2600 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Gopher2600 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Gopher2600.  If not, see <https://www.gnu.org/licenses/>.
//
// *** NOTE: all historical versions of this file, as found in any
// git repository, are also covered by the licence, even when this
// notice is not present ***

package sdlimgui

import "github.com/inkyblackness/imgui-go/v2"

// requires the minimum Vec2{} required to fit any of the string values
// listed in the arguments
func imguiGetFrameDim(s string, t ...string) imgui.Vec2 {
	w := imgui.CalcTextSize(s, false, 0)
	for i := range t {
		y := imgui.CalcTextSize(t[i], false, 0)
		if y.X > w.X {
			w = y
		}
	}
	w.Y = imgui.FontSize() + (imgui.CurrentStyle().FramePadding().Y * 2.0)
	return w
}

// draw toggle button at current cursor position
func imguiToggleButton(id string, v *bool, col imgui.Vec4) {
	bg := imgui.PackedColorFromVec4(col)
	p := imgui.CursorScreenPos()
	dl := imgui.WindowDrawList()

	height := imgui.FrameHeight()
	width := height * 1.55
	radius := height * 0.50
	t := float32(0.0)
	if *v {
		t = 1.0
	}

	// const animSpeed = 0.08
	// ctx, _ := imgui.CurrentContext()
	// if ctx.LastActiveId == ctx.CurrentWindow.GetID(id) {
	// 	tanim := ctx.LastActiveIdTimer / animSpeed
	// 	if tanim < 0.0 {
	// 		tanim = 0.0
	// 	} else if tanim > 1.0 {
	// 		tanim = 1.0
	// 	}
	// 	if *v {
	// 		t = tanim
	// 	} else {
	// 		t = 1.0 - tanim
	// 	}
	// }

	imgui.InvisibleButtonV(id, imgui.Vec2{width, height})
	if imgui.IsItemClicked() {
		*v = !*v
	}

	dl.AddRectFilledV(p, imgui.Vec2{p.X + width, p.Y + height}, bg, radius, imgui.DrawCornerFlagsAll)
	dl.AddCircleFilled(imgui.Vec2{p.X + radius + t*(width-radius*2.0), p.Y + radius},
		radius-1.5, imgui.PackedColorFromVec4(imgui.Vec4{1.0, 1.0, 1.0, 1.0}))
}

// input text that accepts a maximum number of hex digits
func imguiHexInput(label string, aggressiveUpdate bool, digits int, content *string, update func()) {
	cb := func(d imgui.InputTextCallbackData) int32 {
		b := string(d.Buffer())

		// restrict length of input to two characters. note that restriction to
		// hexadecimal characters is handled by imgui's CharsHexadecimal flag
		// given to InputTextV()
		if len(b) > digits {
			d.DeleteBytes(0, len(b))
			b = b[:digits]
			d.InsertBytes(0, []byte(b))
			d.MarkBufferModified()
		}

		return 0
	}

	// flags used with InputTextV()
	flags := imgui.InputTextFlagsCharsHexadecimal |
		imgui.InputTextFlagsCallbackAlways |
		imgui.InputTextFlagsAutoSelectAll

	// with aggressiveUpdate the values entered will be given to the onEnter()
	// function immediately and not just when the enter key is pressed.
	if aggressiveUpdate {
		flags |= imgui.InputTextFlagsEnterReturnsTrue
	}

	if imgui.InputTextV(label, content, flags, cb) {
		update()
	}
}

func imguiLeftlabel(label string) {
	imgui.AlignTextToFramePadding()
	imgui.Text(label)
	imgui.SameLine()
}
