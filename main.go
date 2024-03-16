package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/antlr4-go/antlr/v4"
	"github.com/beevik/etree"
	"github.com/twpayne/go-kml"
	parser "thinh.lth/kml/internal/ast"
	"thinh.lth/kml/internal/service"
)

var (
	_ = etree.NewDocument()
	_ = kml.Namespace
)

func findFieldNames(kmlFile string) {
	// fieldNames := make(map[string]struct{})

	// Load the KML file
	doc := etree.NewDocument()
	err := doc.ReadFromFile(kmlFile)
	if err != nil {
		fmt.Println(err)
	}
	mapFieldTags := make(map[string]struct{})
	childs := doc.ChildElements()

	for n := 0; n < len(childs); n += 1 {
		childs = append(childs, childs[n].ChildElements()...)
		mapFieldTags[childs[n].FullTag()] = struct{}{}
	}

	for i := range mapFieldTags {
		fmt.Println(i)
	}
}

func main() {
	file, err := os.Open(filepath.Join(".", "eligible_area_yes.kml"))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	inputStream := antlr.NewIoStream(file)
	// construct lexer stream
	lx := parser.NewXMLLexer(inputStream)
	lxStream := antlr.NewCommonTokenStream(lx, antlr.LexerDefaultTokenChannel)

	// construct parser stream
	ps := parser.NewXMLParser(lxStream)

	visitor := service.KmlListener{}
	antlr.ParseTreeWalkerDefault.Walk(&visitor, ps.Document())
}
