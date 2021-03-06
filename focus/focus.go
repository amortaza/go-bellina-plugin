package focus

import (
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-hal"
)

type Event struct {
	FocusNodeId      string
	LoseFocusNodeId  string
	IsGainFocusEvent bool
	IsKeyEvent       bool
	KeyEvent         *bl.KeyEvent
}

func init() {
	g_onKeyByNodeId = make(map[string] func(interface{}))
	g_onLoseFocusByNodeId = make(map[string] func(interface{}))
	g_onGainFocusByNodeId = make(map[string] func(interface{}))

	bl.Register_LifeCycle_Init(onBlInit)
}

func On(onKey func(interface{})) {
	On_LifeCycle(onKey, nil, nil)
}

func On_LifeCycle(onKey func(interface{}), onGainFocus func(interface{}), onLoseFocus func(interface{})) {

	nodeId := bl.Current_Node.Id

	click.On( func(i interface{}) {
		e := i.(click.Event)

		if g_lastNodeId != e.Target.Id {
			if onGainFocus != nil && e.Target.Id == nodeId {
				onGainFocus(newFocusGainLoseEvent(nodeId, g_lastNodeId))
			}
		}

		if nodeId == e.Target.Id {
			g_lastNodeId = e.Target.Id
		}
	})

	if onKey != nil {
		g_onKeyByNodeId[nodeId] = onKey
	}

	if onLoseFocus != nil {
		g_onLoseFocusByNodeId[nodeId] = onLoseFocus
	}

	if onGainFocus != nil {
		g_onGainFocusByNodeId[nodeId] = onGainFocus
	}

}

func onBlInit() {

	bl.RegisterLongTerm(bl.EventType_Key, func(e bl.Event) {

		if g_lastNodeId == "" {
			return
		}

		onKey, ok := g_onKeyByNodeId[g_lastNodeId]

		if ok {
			onKey(newFocusKeyEvent(g_lastNodeId, e.(*bl.KeyEvent)))
		}
	})

	bl.RegisterLongTerm(bl.EventType_Mouse_Button, func(mbe bl.Event) {
		if g_lastNodeId == "" {
			return
		}

		e := mbe.(*bl.MouseButtonEvent)

		if e.ButtonAction == hal.Button_Action_Down {
			return
		}

		if e.Target.Id != g_lastNodeId {

			onLoseFocus, ok := g_onLoseFocusByNodeId[g_lastNodeId]

			if ok {
				onLoseFocus(newFocusGainLoseEvent(e.Target.Id, g_lastNodeId))
			}

			newId := e.Target.Id

			onGainFocus, ok2 := g_onGainFocusByNodeId[newId]

			if ok2 {
				onGainFocus(newFocusGainLoseEvent(newId, g_lastNodeId))
			}

			g_lastNodeId = newId
		}
	})
}