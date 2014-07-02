package sdl

// #cgo LDFLAGS: SDL2.dll
// #include "include/SDL.h"
//
// int getSdlEventType(SDL_Event event){
//	return event.type;
//}
//
// SDL_MouseMotionEvent eventAsMouseMotion(SDL_Event event){
//	return event.motion;
//}
//
// SDL_KeyboardEvent eventAsKeyboardEvent(SDL_Event event){
//	return event.key;
//}
import "C"
import (
	"unsafe"
)

type Window struct {
	w *C.SDL_Window
	r *C.SDL_Renderer
}
type Context *C.SDL_GLContext

type WindowFlag uint32

const (
	WindowFullscreen        WindowFlag = C.SDL_WINDOW_FULLSCREEN
	WindowFullscreenDesktop WindowFlag = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	WindowOpengl            WindowFlag = C.SDL_WINDOW_OPENGL
	WindowShown             WindowFlag = C.SDL_WINDOW_SHOWN
	WindowHidden            WindowFlag = C.SDL_WINDOW_HIDDEN
	WindowBorderless        WindowFlag = C.SDL_WINDOW_BORDERLESS
	WindowResizable         WindowFlag = C.SDL_WINDOW_RESIZABLE
	WindowMinimized         WindowFlag = C.SDL_WINDOW_MINIMIZED
	WindowMaximized         WindowFlag = C.SDL_WINDOW_MAXIMIZED
	WindowInputGrabbed      WindowFlag = C.SDL_WINDOW_INPUT_GRABBED
	WindowInputFocus        WindowFlag = C.SDL_WINDOW_INPUT_FOCUS
	WindowMouseFocus        WindowFlag = C.SDL_WINDOW_MOUSE_FOCUS
	WindowForeign           WindowFlag = C.SDL_WINDOW_FOREIGN
	WindowHighDPI           WindowFlag = C.SDL_WINDOW_ALLOW_HIGHDPI
)

func sdlErr() error {
	errInfo := C.GoString(C.SDL_GetError())
	C.SDL_ClearError()
	return &sdlError{errInfo}
}

func checksdlErr(code C.int) error {
	if code != 0 {
		return sdlErr()
	} else {
		return nil
	}
}

type sdlError struct {
	info string
}

func (err *sdlError) Error() string {
	return err.info
}

func CreateWindowRenderer(name string, posX, posY, height, width int, flags WindowFlag) (*Window, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	window := Window{w: C.SDL_CreateWindow(cname, C.int(posX), C.int(posY),
		C.int(height), C.int(width), C.Uint32(flags))}

	if window.w == nil {
		return nil, sdlErr()
	}
	window.r = C.SDL_CreateRenderer(window.w, -1, 0)
	if window.r == nil {
		return nil, sdlErr()
	}

	return &window, nil
}

func (w *Window) Close() {
	C.SDL_DestroyWindow(w.w)
}

func (w *Window) Present() {
	C.SDL_RenderPresent(w.r)
}

func (w *Window) Clear() {
	C.SDL_RenderClear(w.r)
}

func (w *Window) SetDrawColor(r, g, b, a uint8) error {
	return checksdlErr(C.SDL_SetRenderDrawColor(w.r, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

func (w *Window) RenderDrawLine(x1, y1, x2, y2 int) error {
	return checksdlErr(C.SDL_RenderDrawLine(w.r, C.int(x1), C.int(y1), C.int(x2), C.int(y2)))
}

type Texture *C.SDL_Texture

func (w *Window) CreateTexture(width, height int) (Texture, error) {
	tex := Texture(C.SDL_CreateTexture(w.r, C.SDL_PIXELFORMAT_RGBA8888,
		C.SDL_TEXTUREACCESS_STATIC,
		C.int(width), C.int(height)))
	if tex == nil {
		return nil, sdlErr()
	}
	return tex, nil
}

func (w *Window) RenderCopy(tex Texture, x, y, width, height int) error {
	//src := C.SDL_Rect{0, 0, C.int(width), C.int(height)}
	dst := C.SDL_Rect{C.int(x), C.int(y), C.int(width), C.int(height)}
	return checksdlErr(C.SDL_RenderCopy(w.r, tex, nil, &dst))
}

func UpdateTexture(tex Texture, pixels []byte, width int) error {
	return checksdlErr(C.SDL_UpdateTexture(tex, nil,
		unsafe.Pointer(&pixels[0]), C.int(width*4)))
}

func SdlInit() error {
	return checksdlErr(C.SDL_Init(C.SDL_INIT_EVERYTHING))
}

func Quit() {
	C.SDL_Quit()
}

type GlContext struct{ g C.SDL_GLContext }

type Event interface {
	event()
}

type QuitEvent struct{}
type unkownEvent struct{}
type MouseMoveEvent struct{ position [2]int }
type KeyupEvent struct{ Key string }
type KeydownEvent struct{ Key string }

func (*QuitEvent) event()      {}
func (*unkownEvent) event()    {}
func (*MouseMoveEvent) event() {}
func (*KeyupEvent) event()     {}
func (*KeydownEvent) event()   {}

func PollEvent() Event {
	var event C.SDL_Event
	if C.SDL_PollEvent(&event) == 1 {
		eventType := C.getSdlEventType(event)
		switch {
		case eventType == C.SDL_QUIT:
			return &QuitEvent{}
		case eventType == C.SDL_MOUSEMOTION:
			mouseEvent := C.eventAsMouseMotion(event)
			return &MouseMoveEvent{[2]int{int(mouseEvent.x), int(mouseEvent.y)}}
		case eventType == C.SDL_KEYUP:
			return &KeyupEvent{getKey(C.eventAsKeyboardEvent(event))}
		case eventType == C.SDL_KEYDOWN:
			return &KeydownEvent{getKey(C.eventAsKeyboardEvent(event))}
		}
		return &unkownEvent{}
	}
	return nil
}

func getKey(event C.SDL_KeyboardEvent) string {
	return C.GoString(C.SDL_GetKeyName(event.keysym.sym))
}
