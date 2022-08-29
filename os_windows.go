package terem

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type InputEvent struct {
	EventType uint16
	_         [2]byte
	Event     [16]byte
}

type InputCombo struct {
	Key    InputKey
	Char   rune
	Ctrl   bool
	Alt    bool
	Shift  bool
	Caps   bool
	Num    bool
	Scroll bool
}

type InputKey uint8

const (
	KeyLButton InputKey = 0x01
	KeyRButton InputKey = 0x02
	KeyCancel  InputKey = 0x03
	KeyMButton InputKey = 0x04

	KeyEscape InputKey = 0x1b
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

	// winIn = windows.Handle(os.Stdin.Fd())

	if err = windows.SetConsoleMode(winIn, 0x8); err != nil {
		panic(err)
	}

	if winOut, err = windows.Open("CONOUT$", windows.O_RDWR, 0); err != nil {
		panic(err)
	}

	// winOut = windows.Handle(os.Stdout.Fd())

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
