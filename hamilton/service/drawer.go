package service

import (
	"fmt"
	svg "github.com/ajstarks/svgo"
	_ "github.com/boombuler/barcode"
	code128 "github.com/boombuler/barcode/code128"
	"golang.org/x/crypto/bcrypt"
	"image/color"
	"math"
	"strings"
)

type Drawer struct {
}

func NewDrawer() Drawer {
	return Drawer{}
}

func (d Drawer) Spin(cost int) error {
	_, err := bcrypt.GenerateFromPassword([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), cost)
	return err
}

func (d Drawer) DrawTicket(ticket Ticket) (string, error) {
	err := d.Spin(16)
	if err != nil {
		return "", err
	}
	w := &strings.Builder{}

	width := 1200
	height := 518
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Desc(ticket.ID)
	canvas.Rect(0, 0, width, height, "fill:rgb(255,255,255)")
	canvas.Rect(0, 0, 200, height, "fill:rgb(92,158,173)")
	canvas.Rect(200, 0, 20, height, "fill:rgb(50,98,115)")
	canvas.Circle(210, 100, 50, "fill:rgb(50,98,115)")
	canvas.Text(300, 100, ticket.Concert.Name, "text-anchor:left;font-size:36px;fill:black")
	canvas.Text(300, 150, ticket.Concert.Date, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 200, ticket.Concert.Venue, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 250, ticket.Concert.ID, "text-anchor:left;font-size:24px;fill:black")
	canvas.Line(300, 300, 900, 300, "stroke:rgb(0,0,0);stroke-width:2")
	canvas.Text(300, 350, ticket.Name, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 400, ticket.Email, "text-anchor:left;font-size:24px;fill:black")
	canvas.Text(300, 450, ticket.ID, "text-anchor:left;font-size:24px;fill:black")
	canvas.Rect(1190, 0, 10, height, "fill:rgb(255,215,0)")
	err = d.barcode(canvas, ticket.ID)
	if err != nil {
		return "", err
	}
	canvas.End()
	return w.String(), nil
}

func (d Drawer) DrawConcert(concert Concert) (string, error) {
	err := d.Spin(18)
	if err != nil {
		return "", err
	}
	w := &strings.Builder{}

	size := 20
	if concert.Seats.Max > 1000 {
		size = int(20 - math.Log(float64(concert.Seats.Max)))
	}

	rowSize := 1000 / size
	width := 1200
	height := 250 + (concert.Seats.Max / rowSize * size) + 50

	canvas := svg.New(w)
	canvas.Start(width, height)

	canvas.Desc(fmt.Sprintf("%s|%d|%d", concert.ID, concert.Seats.Max, concert.Seats.Purchased))
	canvas.Rect(0, 0, width, height, "fill:rgb(255,255,255)")
	canvas.Text(600, 100, concert.Name, "text-anchor:middle;font-size:36px;fill:black")
	canvas.Text(600, 150, concert.Venue, "text-anchor:middle;font-size:24px;fill:black")
	canvas.Line(300, 170, 900, 170, "stroke:rgb(0,0,0);stroke-width:2")

	err = d.seats(canvas, concert.Seats.Max, concert.Seats.Purchased, size, rowSize)
	if err != nil {
		return "", err
	}

	canvas.End()
	return w.String(), nil
}

func (d Drawer) barcode(canvas *svg.SVG, uuid string) error {
	bc, err := code128.Encode(uuid)
	if err != nil {
		return err
	}
	for i := 0; i < bc.Bounds().Dx(); i++ {
		if bc.At(i, 0) == color.Black {
			canvas.Rect(1000, 80+i, 150, 1, "fill:rgb(0,0,0)")
		}
	}

	return nil
}

func (d Drawer) seats(canvas *svg.SVG, total int, purchased int, size int, rowSize int) error {
	row := 0
	seat := 0
	shift := 0
	canvas.Style("text/css", "rect {fill: white; stroke-width: 1; stroke: black;}")
	for i := 0; i < total; i++ {

		row = i / rowSize
		seat = (i % rowSize) + 1

		if seat%2 == 1 {
			seat = int(math.Ceil(float64(seat)/2.0) - 1)
			seat = seat * -1
		} else {
			seat = seat / 2
		}

		shift = (row % 3) - 1

		if i < purchased {
			canvas.CenterRect(600+(seat*size)+(shift*6), 250+row*size, size, size, "fill:rgb(50,98,115);")
			continue
		}
		canvas.CenterRect(600+(seat*size)+(shift*6), 250+row*size, size, size)
	}
	return nil
}
