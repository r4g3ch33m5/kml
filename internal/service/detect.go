package service

import (
	"fmt"
	"strings"

	parser "thinh.lth/kml/internal/ast"
)

// class polygon constructor
type KmlListener struct {
	parser.BaseXMLParserListener

	MultiPolygons []MultiPolygon

	CurrentPolygons []Polygon
}

type ElementNameVisitor struct {
	parser.BaseXMLParserListener

	mapElementName map[string]bool
}

func NewElementNameListener() ElementNameVisitor {
	return ElementNameVisitor{mapElementName: make(map[string]bool)}
}

func (v *ElementNameVisitor) EnterElement(ctx *parser.ElementContext) {
	v.mapElementName[ctx.Name(0).GetText()] = true
}

func (v *ElementNameVisitor) ExitDocument(ctx *parser.DocumentContext) {
	for name := range v.mapElementName {
		fmt.Println(name)
	}
}

var (
	CoordinatesName  = strings.ToLower("coordinates")
	MultiPolygonName = strings.ToLower("MultiGeometry")
)

type Point struct {
	Longitude string
	Latitude  string
}

type Polygon struct {
	Points []Point
}

type MultiPolygon struct {
	Polygons []Polygon
}

func (v *KmlListener) ExitDocument(ctx *parser.DocumentContext) {
	for _, multiPolygon := range v.MultiPolygons {
		for _, polygon := range multiPolygon.Polygons {
			fmt.Println(polygon)
		}
	}
}

func (v *KmlListener) ExitElement(ctx *parser.ElementContext) {
	name := strings.ToLower(ctx.Name(0).GetSymbol().GetText()) // name of attribute, name 1 and name 2 should be same
	switch name {
	case CoordinatesName:
		rawText := ctx.Content().GetText()
		points := strings.Fields(rawText)
		pointStructs := make([]Point, len(points))
		for idx, point := range points {
			val := strings.Split(point, ",")
			pointStructs[idx] = Point{
				Longitude: val[0],
				Latitude:  val[1],
			}
		}
		v.CurrentPolygons = append(v.CurrentPolygons, Polygon{Points: pointStructs})

	case MultiPolygonName:
		//
		v.MultiPolygons = append(v.MultiPolygons, MultiPolygon{
			Polygons: v.CurrentPolygons,
		})

		v.CurrentPolygons = make([]Polygon, 0) // clear current polygon
	}
}
