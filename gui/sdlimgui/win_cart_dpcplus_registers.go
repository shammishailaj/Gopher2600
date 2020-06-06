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

import (
	"fmt"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/jetsetilly/gopher2600/hardware/memory/cartridge/harmony"
)

const winDPCplusRegistersTitle = "DPC+ Registers"

type winDPCplusRegisters struct {
	windowManagement
	widgetDimensions

	img *SdlImgui
}

func newWinDPCplusRegisters(img *SdlImgui) (managedWindow, error) {
	win := &winDPCplusRegisters{
		img: img,
	}

	return win, nil
}

func (win *winDPCplusRegisters) init() {
	win.widgetDimensions.init()
}

func (win *winDPCplusRegisters) destroy() {
}

func (win *winDPCplusRegisters) id() string {
	return winDPCplusRegistersTitle
}

func (win *winDPCplusRegisters) draw() {
	if !win.open {
		return
	}

	// do not open window if there is no valid cartridge debug bus available
	r, ok := win.img.lz.Cart.Registers.(harmony.DPCplusRegisters)
	if !win.img.lz.Cart.HasDebugBus || !ok {
		return
	}

	imgui.SetNextWindowPosV(imgui.Vec2{659, 35}, imgui.ConditionFirstUseEver, imgui.Vec2{0, 0})
	imgui.BeginV(winDPCplusRegistersTitle, &win.open, imgui.WindowFlagsAlwaysAutoResize)

	// random number generator value
	rng := fmt.Sprintf("%08x", r.RNG.Value)
	imguiText("Random Number Generator")
	imgui.PushItemWidth(win.eightDigitDim.X)
	if imguiHexInput("##rng", !win.img.paused, 8, &rng) {
		win.img.lz.Dbg.PushRawEvent(func() {
			b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
			b.PutRegister("rng", rng)
		})
	}
	imgui.PopItemWidth()

	imgui.SameLineV(0, 20)
	imguiText("Fast Fetch")
	ff := r.FastFetch
	if imgui.Checkbox("##fastfetch", &ff) {
		win.img.lz.Dbg.PushRawEvent(func() {
			b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
			b.PutRegister("fastfetch", fmt.Sprintf("%v", ff))
		})
	}

	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	// *** data fetchers grouping ***
	imgui.BeginGroup()

	// loop over data fetchers
	imgui.Text("Data Fetchers")
	imgui.Spacing()
	for i := 0; i < len(r.Fetcher); i++ {
		f := i

		imguiText(fmt.Sprintf("#%d", f))

		label := fmt.Sprintf("##d%dlow", i)
		low := fmt.Sprintf("%02x", r.Fetcher[i].Low)
		imguiText("Low")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &low) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("fetcher::%d::low", f), low)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##d%dhi", i)
		hi := fmt.Sprintf("%02x", r.Fetcher[i].Hi)
		imguiText("Hi")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &hi) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("fetcher::%d::hi", f), hi)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##d%dtop", i)
		top := fmt.Sprintf("%02x", r.Fetcher[i].Top)
		imguiText("Top")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &top) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("fetcher::%d::top", f), top)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##d%dbottom", i)
		bottom := fmt.Sprintf("%02x", r.Fetcher[i].Bottom)
		imguiText("Bottom")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &bottom) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("fetcher::%d::bottom", f), bottom)
			})
		}
		imgui.PopItemWidth()
	}
	imgui.EndGroup()

	// *** fraction fetchers grouping ***
	imgui.SameLineV(0, 20)
	imgui.BeginGroup()

	// loop over fractional fetchers
	imgui.Text("Fractional Fetchers")
	imgui.Spacing()
	for i := 0; i < len(r.FracFetcher); i++ {
		f := i

		imguiText(fmt.Sprintf("#%d", f))

		label := fmt.Sprintf("##f%dlow", i)
		low := fmt.Sprintf("%02x", r.FracFetcher[i].Low)
		imguiText("Low")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &low) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("frac::%d::low", f), low)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##f%dhi", i)
		hi := fmt.Sprintf("%02x", r.FracFetcher[i].Hi)
		imguiText("Hi")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &hi) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("frac::%d::hi", f), hi)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##f%dincrement", i)
		increment := fmt.Sprintf("%02x", r.FracFetcher[i].Increment)
		imguiText("Increment")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &increment) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("frac::%d::increment", f), increment)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##f%dcount", i)
		count := fmt.Sprintf("%02x", r.FracFetcher[i].Count)
		imguiText("Count")
		imgui.PushItemWidth(win.twoDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 2, &count) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("frac::%d::count", f), count)
			})
		}
		imgui.PopItemWidth()
	}
	imgui.EndGroup()

	// *** music fetchers grouping ***
	imgui.Spacing()
	imgui.Separator()
	imgui.Spacing()

	imgui.BeginGroup()

	// loop over music fetchers
	imgui.Text("Music Fetchers")
	imgui.Spacing()
	for i := 0; i < len(r.MusicFetcher); i++ {
		f := i

		imguiText(fmt.Sprintf("#%d", f))

		label := fmt.Sprintf("##m%dwaveform", i)
		waveform := fmt.Sprintf("%08x", r.MusicFetcher[i].Waveform)
		imguiText("Waveform")
		imgui.PushItemWidth(win.eightDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 8, &waveform) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("music::%d::waveform", f), waveform)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##m%dfeq", i)
		freq := fmt.Sprintf("%08x", r.MusicFetcher[i].Freq)
		imguiText("Freq")
		imgui.PushItemWidth(win.eightDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 8, &freq) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("music::%d::freq", f), freq)
			})
		}
		imgui.PopItemWidth()

		imgui.SameLine()
		label = fmt.Sprintf("##m%dcount", i)
		count := fmt.Sprintf("%08x", r.MusicFetcher[i].Count)
		imguiText("Count")
		imgui.PushItemWidth(win.eightDigitDim.X)
		if imguiHexInput(label, !win.img.paused, 8, &count) {
			win.img.lz.Dbg.PushRawEvent(func() {
				b := win.img.lz.Dbg.VCS.Mem.Cart.GetDebugBus()
				b.PutRegister(fmt.Sprintf("music::%d::count", f), count)
			})
		}
		imgui.PopItemWidth()
	}
	imgui.EndGroup()

	imgui.End()
}