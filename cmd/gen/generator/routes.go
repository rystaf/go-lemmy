package generator

import (
	"io"
	"strings"

	"github.com/dave/jennifer/jen"
	"go.arsenm.dev/go-lemmy/cmd/gen/parser"
)

type RoutesGenerator struct {
	w       io.Writer
	PkgName string
}

func NewRoutes(w io.Writer, pkgName string) *RoutesGenerator {
	return &RoutesGenerator{w, pkgName}
}

func (r *RoutesGenerator) Generate(routes []parser.Route, impls map[string]string) error {
	f := jen.NewFile(r.PkgName)
	for _, r := range routes {
		resStruct := impls[r.Struct]

		f.Func().Params(
			jen.Id("c").Id("*Client"),
		).Id(strings.TrimPrefix(r.Struct, "Get")).Params(
			jen.Id("ctx").Qual("context", "Context"),
			jen.Id("data").Qual("go.arsenm.dev/go-lemmy/types", r.Struct),
		).Params(
			jen.Op("*").Qual("go.arsenm.dev/go-lemmy/types", resStruct),
			jen.Error(),
		).BlockFunc(func(g *jen.Group) {
			g.Id("resData").Op(":=").Op("&").Qual("go.arsenm.dev/go-lemmy/types", resStruct).Block()

			var funcName string
			switch r.Method {
			case "GET":
				funcName = "getReq"
			default:
				funcName = "req"
			}

			g.List(jen.Id("res"), jen.Err()).Op(":=").Id("c").Dot(funcName).Params(
				jen.Id("ctx"), jen.Lit(r.Method), jen.Lit(r.Path), jen.Id("data"), jen.Op("&").Id("resData"),
			)
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()),
			)

			g.Err().Op("=").Id("resError").Params(jen.Id("res"), jen.Id("resData").Dot("LemmyResponse"))
			g.If(jen.Err().Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Err()),
			)

			g.Return(jen.Id("resData"), jen.Nil())
		})
	}

	return f.Render(r.w)
}
