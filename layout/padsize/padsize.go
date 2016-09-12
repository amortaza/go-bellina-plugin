package padsize

import (
	"github.com/amortaza/go-bellina"
)

func Id() (*State) {
	bl.RequireSettledBoundaries()
	bl.RequireSettledKids()

	state := ensureState(bl.Current_Node.Id)

	return state
}

func runLogic(node *bl.Node, state *State) {

	padLeft := state.Z_Left
	padTop := state.Z_Top
	padRight := state.Z_Right
	padBottom := state.Z_Bottom

	for e := node.Kids.Front(); e != nil; e = e.Next() {
	    	kid := e.Value.(*bl.Node)

		if kid.Left < padLeft {
			//kid.OwnLeft("pad")
			kid.Left = padLeft
		}

		if kid.Top < padTop {
			//kid.OwnTop("pad")
			kid.Top = padTop
		}

		right := kid.Left + kid.Width - 1

		if right > node.Width - padRight {
			//kid.OwnWidth("pad")
			kid.Width = node.Width - padRight - kid.Left
		}

		bottom := kid.Top + kid.Height - 1

		if bottom > node.Height - padBottom {
			//kid.OwnHeight("pad")
			kid.Height = node.Height - padBottom - kid.Top
		}
	}
}
