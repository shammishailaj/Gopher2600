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
	"gopher2600/gui"

	"github.com/inkyblackness/imgui-go/v2"
	"github.com/veandco/go-sdl2/sdl"
)

// Service implements GuiCreator interface
func (img *SdlImgui) Service() {
	// do not check for events if no event channel has been set
	if img.events != nil {

		// loop until there are no more events to retreive. this loop is
		// intimately connected with the framelimiter below. what we don't want
		// to loop for too long servicing events. however:
		//
		// 1. servicing just one event per frame is not enough, queued events
		//    will take one frame on longer to resolve
		//
		// 2. limiting the number of events serviced per frame has the same
		//    problem for very long queues
		//
		// 3. truncating events is not wanted because we may miss important
		//    user input
		//
		// best solution is the poll loop
		for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {

			switch ev := ev.(type) {
			// close window
			case *sdl.QuitEvent:
				img.events <- gui.EventWindowClose{}

			case *sdl.KeyboardEvent:
				mod := gui.KeyModNone

				if sdl.GetModState()&sdl.KMOD_LALT == sdl.KMOD_LALT ||
					sdl.GetModState()&sdl.KMOD_RALT == sdl.KMOD_RALT {
					mod = gui.KeyModAlt
				} else if sdl.GetModState()&sdl.KMOD_LSHIFT == sdl.KMOD_LSHIFT ||
					sdl.GetModState()&sdl.KMOD_RSHIFT == sdl.KMOD_RSHIFT {
					mod = gui.KeyModShift
				} else if sdl.GetModState()&sdl.KMOD_LCTRL == sdl.KMOD_LCTRL ||
					sdl.GetModState()&sdl.KMOD_RCTRL == sdl.KMOD_RCTRL {
					mod = gui.KeyModCtrl
				}

				switch ev.Type {
				case sdl.KEYDOWN:
					if ev.Repeat == 0 {
						img.events <- gui.EventKeyboard{
							Key:  sdl.GetKeyName(ev.Keysym.Sym),
							Mod:  mod,
							Down: true}
					}
				case sdl.KEYUP:
					if ev.Repeat == 0 {
						img.events <- gui.EventKeyboard{
							Key:  sdl.GetKeyName(ev.Keysym.Sym),
							Mod:  mod,
							Down: false}
					}
				}

			case *sdl.MouseButtonEvent:
				button := gui.MouseButtonLeft
				if ev.Button == sdl.BUTTON_RIGHT {
					button = gui.MouseButtonRight
				}

				img.events <- gui.EventMouseButton{
					Button: button,
					Down:   ev.Type == sdl.MOUSEBUTTONDOWN}
			}
		}

		// mouse motion
		if img.isCaptured {
			mx, my, _ := sdl.GetMouseState()
			if mx != img.mx || my != img.my {
				w, h := img.plt.window.GetSize()

				// reduce mouse x and y coordintes to the range 0.0 to 1.0
				//  no need to worry about negative numbers and numbers greater
				//  than 1.0 because we (should) have restricted mouse movement
				//  to the window (with window.SetGrab(). see the ReqCaptureMouse
				//  case in the SetFeature() function)
				x := float32(mx) / float32(w)
				y := float32(my) / float32(h)

				img.events <- gui.EventMouseMotion{X: x, Y: y}
				img.mx = mx
				img.my = my
			}
		}
	}

	// Signal start of a new frame
	img.plt.newFrame()
	imgui.NewFrame()

	// imgui commands
	img.screen.draw()

	// Rendering
	imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

	clearColor := [4]float32{0.0, 0.0, 0.0, 1.0}
	img.glsl.preRender(clearColor)
	img.screen.render()
	img.glsl.render(img.plt.displaySize(), img.plt.framebufferSize(), imgui.RenderedDrawData())
	img.plt.postRender()

	// wait for frame limiter
	img.lmtr.Wait()

	// run any outstanding service functions
	select {
	case f := <-img.service:
		f()
	default:
	}
}