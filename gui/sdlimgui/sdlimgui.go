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
	"io"

	"github.com/jetsetilly/gopher2600/debugger/terminal"
	"github.com/jetsetilly/gopher2600/gui"
	"github.com/jetsetilly/gopher2600/gui/sdlaudio"
	"github.com/jetsetilly/gopher2600/gui/sdlimgui/lazyvalues"
	"github.com/jetsetilly/gopher2600/paths"
	"github.com/jetsetilly/gopher2600/television"

	"github.com/inkyblackness/imgui-go/v2"
)

const imguiIniFile = "debugger_imgui.ini"

// SdlImgui is an sdl based visualiser using imgui
type SdlImgui struct {
	// the mechanical requirements for the gui
	io      imgui.IO
	context *imgui.Context
	plt     *platform
	glsl    *glsl

	// references to the emulation
	lazy *lazyvalues.Values
	tv   television.Television

	// terminal interface to the debugger
	term *term

	// implementations of screen television protocols
	screen *screen
	audio  *sdlaudio.Audio

	// imgui window management
	wm *windowManager

	// the colors used by the imgui system. includes the TV colors in a
	// suitable format
	cols *imguiColors

	// functions that need to be performed in the main thread should be queued
	// for service
	service    chan func()
	serviceErr chan error

	// SetFeature() hands off requests to the featureReq channel for servicing.
	// think of this as a special instance of the service chan
	featureReq chan featureRequest
	featureErr chan error

	// events channel is not created but assigned with SetEventChannel()
	events chan gui.Event

	// is emulation running
	paused bool

	// mouse coords at last frame
	mx, my int32
}

// NewSdlImgui is the preferred method of initialisation for type SdlImgui
//
// MUST ONLY be called from the #mainthread
func NewSdlImgui(tv television.Television) (*SdlImgui, error) {
	img := &SdlImgui{
		context:    imgui.CreateContext(nil),
		io:         imgui.CurrentIO(),
		tv:         tv,
		service:    make(chan func(), 1),
		serviceErr: make(chan error, 1),
		featureReq: make(chan featureRequest, 1),
		featureErr: make(chan error, 1),
	}

	// define colors
	img.cols = newColors()

	var err error

	img.plt, err = newPlatform(img)
	if err != nil {
		return nil, err
	}

	img.glsl, err = newGlsl(img.io, img)
	if err != nil {
		return nil, err
	}

	iniPath, err := paths.ResourcePath("", imguiIniFile)
	if err != nil {
		return nil, err
	}
	img.io.SetIniFilename(iniPath)

	// we don't have access to the Debugger, Disassembly or the VCS yet. those
	// fields in the lazy instance will be set when the requests come in
	img.lazy = lazyvalues.NewValues()

	img.screen = newScreen(img)
	img.term = newTerm()

	img.wm, err = newWindowManager(img)
	if err != nil {
		return nil, err
	}

	// connect some screen properties to other parts of the system
	img.glsl.screenTextureID = img.screen.screenTexture
	tv.AddPixelRenderer(img.screen)

	// this audio mixer produces the sound. there is another AudioMixer
	// implementation in winAudio which visualises the sound
	img.audio, err = sdlaudio.NewAudio()
	if err != nil {
		return nil, err
	}
	tv.AddAudioMixer(img.audio)

	return img, nil
}

// Destroy implements GuiCreator interface
//
// MUST ONLY be called from the #mainthread
func (img *SdlImgui) Destroy(output io.Writer) {
	img.wm.destroy()
	img.audio.EndMixing()
	img.glsl.destroy()

	err := img.plt.destroy()
	if err != nil {
		output.Write([]byte(err.Error()))
	}

	img.context.Destroy()
}

// GetTerminal implements terminal.Broker interface
func (img *SdlImgui) GetTerminal() terminal.Terminal {
	return img.term
}

func (img *SdlImgui) pause(set bool) {
	img.paused = set
}

func (img *SdlImgui) draw() {
	if img.lazy.Dbg == nil {
	} else {
		img.wm.drawWindows()
	}
}
