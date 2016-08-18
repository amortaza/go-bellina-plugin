package hover

import (
	"github.com/amortaza/go-adt"
	"github.com/amortaza/go-bellina/event"
	"github.com/amortaza/go-bellina"
)

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "hover"
}

func (c *Plugin) Init() {
	callbacksByNodeId = adt.NewCallbacksByID()

	event.RegisterLongTerm( bl.EventType_Mouse_Move, func(e event.Event) {
		event := e.(*bl.MouseMoveEvent)

		curId := event.Target.Id

		if lastNodeId != curId {

			inId, outId := curId, lastNodeId
			
			if callbacksByNodeId.HasId(inId) {
				eIn := newEvent(inId, outId, true)
				callbacksByNodeId.CallAll(inId, eIn)
			}

			if callbacksByNodeId.HasId(outId) {
				eOut := newEvent(inId, outId, false)
				callbacksByNodeId.CallAll(outId, eOut)
			}

			lastNodeId = curId
		}
	})
}

func (c *Plugin) Reset_ShortTerm() {
	callbacksByNodeId.ClearAll()
}

func (c *Plugin) Tick() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}

func (c *Plugin) On(cb func(interface{})) {
	if c == nil {
		panic("You did not load hover plugin")
	}

	callbacksByNodeId.Add(bl.Current_Node.Id, cb)
}

func (c *Plugin) On2(cb func(interface{}), start func(interface{}), end func(interface{})) {
	panic("On2 not suported in hover.Plugin")
}

