package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-oob"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/drag-other"
	"github.com/amortaza/go-bellina-plugins/layout/docker"
)

func fake() {
	var _ = drag_other.Use
}

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(1024, 768)

		bl.Div()
		{
			bl.Id("green")
			bl.Pos(100,100)
			bl.Dim(400,400)

			border.Fill(0,50,0)

			bl.Div()
						{
							bl.Id("blue")
							bl.Pos(50,50)
							bl.Dim(100,100)

							border.Fill(50,0,50)
						}
						bl.End()

			bl.Div()
			{
				bl.Id("red")
				bl.Pos(50,50)
				bl.Dim(100,100)
				bl.SettleBoundary()

				border.Fill(50,0,0)

				drag_other.Use("green")
				//drag_other.Use("blue")
				docker.Id().AnchorBottom(10).AnchorLeft(10).AnchorRight(10).End()
			}
			bl.End()
		}
		bl.End()
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( haloob.New(), 1280, 1024, "Bellina v0.2", initialize, tick, uninit )

	fmt.Println("bye!")
}

