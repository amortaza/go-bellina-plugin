package side_resize

import (
	"fmt"
	"github.com/amortaza/go-bellina"
)

var _LEFT uint32 = 1 << 0
var _TOP uint32 = 1 << 1
var _RIGHT uint32 = 1 << 2
var _BOTTOM uint32 = 1 << 3

var g_startWidth, g_startHeight int
var g_flags uint32
var g_otherId string

type State struct {
}

func (state *State) Left() (*State) {
    g_flags |= _LEFT

    return state
}

func (state *State) Top() (*State) {
    g_flags |= _TOP

    return state
}

func (state *State) Right() (*State) {
    g_flags |= _RIGHT

    return state
}

func (state *State) Bottom() (*State) {
    g_flags |= _BOTTOM

    return state
}

func (state *State) End() {
    logic()
}

func set(shadowOther, shadowCur *bl.ShadowNode) {
    if g_flags & _LEFT != 0 {
        shadowOther.Left__Node_Only("side-resize-other")
        shadowOther.Width__Node_Only("side-resize-other")
    }

    if g_flags & _TOP != 0 {
        shadowOther.Top__Node_Only("side-resize-other")
        shadowOther.Height__Node_Only("side-resize-other")
    }

    if g_flags & _RIGHT != 0 {
        shadowOther.Width__Node_Only("side-resize-other")
        shadowCur.Left__Node_Only("side-resize-other")
    }

    if g_flags & _BOTTOM != 0 {
        shadowOther.Height__Node_Only("side-resize-other")
        shadowCur.Top__Node_Only("side-resize-other")
    }
}

func validate() {
    if g_flags & _LEFT != 0 && g_flags & _RIGHT != 0 {
        fmt.Println("Side Resize Other plugin can not have LEFT and RIGHT flags at the same time")
        panic("See print out")
    }

    if g_flags & _TOP != 0 && g_flags & _BOTTOM != 0 {
        fmt.Println("Side Resize Other plugin can not have TOP and BOTTOM flags at the same time")
        panic("See print out")
    }
}
