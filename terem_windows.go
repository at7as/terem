package terem

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type InputEvent struct {
	EventType InputType
	_         [2]byte
	Event     [16]byte
}

type InputType uint16

const (
	InputTypeKey    InputType = 0x0001
	InputTypeMouse  InputType = 0x0002
	InputTypeBuffer InputType = 0x0004
	InputTypeMenu   InputType = 0x0008
	InputTypeFocus  InputType = 0x0010
)

type InputCombo struct {
	Pressed bool
	Key     InputKey
	Char    rune
	Ctrl    bool
	Alt     bool
	Shift   bool
	Caps    bool
	Num     bool
	Scroll  bool
}

type InputKey uint8

const (
	KeyLButton     InputKey = 0x01
	KeyRButton     InputKey = 0x02
	KeyCancel      InputKey = 0x03
	KeyMButton     InputKey = 0x04
	KeyBackspace   InputKey = 0x08
	KeyTab         InputKey = 0x09
	KeyClear       InputKey = 0x0c
	KeyReturn      InputKey = 0x0d
	KeyShift       InputKey = 0x10
	KeyControl     InputKey = 0x11
	KeyAlt         InputKey = 0x12
	KeyPause       InputKey = 0x13
	KeyCapsLock    InputKey = 0x14
	KeyEscape      InputKey = 0x1b
	KeySpace       InputKey = 0x20
	KeyPageUp      InputKey = 0x21
	KeyPageDown    InputKey = 0x22
	KeyEnd         InputKey = 0x23
	KeyHome        InputKey = 0x24
	KeyLeftArrow   InputKey = 0x25
	KeyUpArrow     InputKey = 0x26
	KeyRightArrow  InputKey = 0x27
	KeyDownArrow   InputKey = 0x28
	KeySelect      InputKey = 0x29
	KeyPrint       InputKey = 0x2a
	KeyExecute     InputKey = 0x2b
	KeyPrintScreen InputKey = 0x2c
	KeyInsert      InputKey = 0x2d
	KeyDelete      InputKey = 0x2e
	KeyHelp        InputKey = 0x2f
	Key0           InputKey = 0x30
	Key1           InputKey = 0x31
	Key2           InputKey = 0x32
	Key3           InputKey = 0x33
	Key4           InputKey = 0x34
	Key5           InputKey = 0x35
	Key6           InputKey = 0x36
	Key7           InputKey = 0x37
	Key8           InputKey = 0x38
	Key9           InputKey = 0x39
	KeyA           InputKey = 0x41
	KeyB           InputKey = 0x42
	KeyC           InputKey = 0x43
	KeyD           InputKey = 0x44
	KeyE           InputKey = 0x45
	KeyF           InputKey = 0x46
	KeyG           InputKey = 0x47
	KeyH           InputKey = 0x48
	KeyI           InputKey = 0x49
	KeyJ           InputKey = 0x4a
	KeyK           InputKey = 0x4b
	KeyL           InputKey = 0x4c
	KeyM           InputKey = 0x4d
	KeyN           InputKey = 0x4e
	KeyO           InputKey = 0x4f
	KeyP           InputKey = 0x50
	KeyQ           InputKey = 0x51
	KeyR           InputKey = 0x52
	KeyS           InputKey = 0x53
	KeyT           InputKey = 0x54
	KeyU           InputKey = 0x55
	KeyV           InputKey = 0x56
	KeyW           InputKey = 0x57
	KeyX           InputKey = 0x58
	KeyY           InputKey = 0x59
	KeyZ           InputKey = 0x5a
	KeyLWin        InputKey = 0x5b
	KeyRWin        InputKey = 0x5c
	KeyApps        InputKey = 0x5d
	KeySleep       InputKey = 0x5f
	KeyNumpad0     InputKey = 0x60
	KeyNumpad1     InputKey = 0x61
	KeyNumpad2     InputKey = 0x62
	KeyNumpad3     InputKey = 0x63
	KeyNumpad4     InputKey = 0x64
	KeyNumpad5     InputKey = 0x65
	KeyNumpad6     InputKey = 0x66
	KeyNumpad7     InputKey = 0x67
	KeyNumpad8     InputKey = 0x68
	KeyNumpad9     InputKey = 0x69
	KeyMultiply    InputKey = 0x6a
	KeyAdd         InputKey = 0x6b
	KeySeparator   InputKey = 0x6c
	KeySubtract    InputKey = 0x6d
	KeyDecimal     InputKey = 0x6e
	KeyDivide      InputKey = 0x6f
	KeyF1          InputKey = 0x70
	KeyF2          InputKey = 0x71
	KeyF3          InputKey = 0x72
	KeyF4          InputKey = 0x73
	KeyF5          InputKey = 0x74
	KeyF6          InputKey = 0x75
	KeyF7          InputKey = 0x76
	KeyF8          InputKey = 0x77
	KeyF9          InputKey = 0x78
	KeyF10         InputKey = 0x79
	KeyF11         InputKey = 0x7a
	KeyF12         InputKey = 0x7b
	KeyNumLock     InputKey = 0x90
	KeyScrollLock  InputKey = 0x91
	KeyLShift      InputKey = 0xa0
	KeyRShift      InputKey = 0xa1
	KeyLControl    InputKey = 0xa2
	KeyRControl    InputKey = 0xa3
	KeyLAlt        InputKey = 0xa4
	KeyRAlt        InputKey = 0xa5
	KeyOEM1        InputKey = 0xba
	KeyOEMPlus     InputKey = 0xbb
	KeyOEMComma    InputKey = 0xbc
	KeyOEMMinus    InputKey = 0xbd
	KeyOEMPeriod   InputKey = 0xbe
	KeyOEM2        InputKey = 0xbf
	KeyOEM3        InputKey = 0xc0
	KeyOEM4        InputKey = 0xdb
	KeyOEM5        InputKey = 0xdc
	KeyOEM6        InputKey = 0xdd
	KeyOEM7        InputKey = 0xde
	KeyOEM8        InputKey = 0xdf
	KeyOEM102      InputKey = 0xe2
)

var (
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	procReadConsoleInputW = kernel32.NewProc("ReadConsoleInputW")
	winIn                 windows.Handle
	winOut                windows.Handle
	winOutMode            uint32
	winRead               uint32
)

func init() {

	var err error

	if winIn, err = windows.Open("CONIN$", windows.O_RDWR, 0); err != nil {
		panic(err)
	}

	if err = windows.SetConsoleMode(winIn, 0x8); err != nil {
		panic(err)
	}

	winOut = windows.Handle(os.Stdout.Fd())

	if err = windows.GetConsoleMode(winOut, &winOutMode); err != nil {
		panic(err)
	}

	if err = windows.SetConsoleMode(winOut, winOutMode|0x4); err != nil {
		panic(err)
	}

}

func readConsoleInput(e *InputEvent) error {

	if r1, _, err := syscall.SyscallN(procReadConsoleInputW.Addr(), uintptr(winIn), uintptr(unsafe.Pointer(e)), 1, uintptr(unsafe.Pointer(&winRead)), 0, 0); r1 == 0 {
		return err
	}
	return nil

}

func getConsoleSize() (int, int, error) {

	var info windows.ConsoleScreenBufferInfo
	err := windows.GetConsoleScreenBufferInfo(winOut, &info)
	return int(info.MaximumWindowSize.X), int(info.MaximumWindowSize.Y), err

}
