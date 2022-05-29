package bot

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/lxn/win"
)

type Mouse int

const (
    LEFT Mouse = iota
    RIGHT
    MIDDLE
)

func Move(x, y int){
	win.SetCursorPos(int32(x), int32(y))
}

func MouseDown(mouse Mouse) {  
	in := []win.MOUSE_INPUT{
		{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{},
		},
	}

  	switch mouse {
    case LEFT:
        in[0].Mi.DwFlags = win.MOUSEEVENTF_LEFTDOWN
    case RIGHT:
        in[0].Mi.DwFlags = win.MOUSEEVENTF_RIGHTDOWN
    case MIDDLE:
        in[0].Mi.DwFlags = win.MOUSEEVENTF_MIDDLEDOWN
    default:
        fmt.Println("Woops")
    }

    win.SendInput(1,  unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

func MouseUp(mouse Mouse) {  
	in := []win.MOUSE_INPUT{
		{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{},
		},
	}

  	switch mouse {
		case LEFT:
			in[0].Mi.DwFlags = win.MOUSEEVENTF_LEFTUP
		case RIGHT:
			in[0].Mi.DwFlags = win.MOUSEEVENTF_RIGHTUP
		case MIDDLE:
			in[0].Mi.DwFlags = win.MOUSEEVENTF_MIDDLEUP
		default:
			fmt.Println("Woops")
    }

    win.SendInput(1,  unsafe.Pointer(&in[0]), int32(unsafe.Sizeof(in[0])));
}

func RightClick(){
	MouseDown(RIGHT)
    time.Sleep(20 * time.Millisecond)
	MouseUp(RIGHT)
}

func MiddleClick(){
	MouseDown(MIDDLE)
    time.Sleep(20 * time.Millisecond)
	MouseUp(MIDDLE)
}

func Click(){
	MouseDown(LEFT)
    time.Sleep(20 * time.Millisecond)
	MouseUp(LEFT)
}