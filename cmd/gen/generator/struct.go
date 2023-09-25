package generator

import (
	"io"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"go.elara.ws/go-lemmy/cmd/gen/extractor"
)

type StructGenerator struct {
	w       io.Writer
	PkgName string
}

func NewStruct(w io.Writer, pkgName string) *StructGenerator {
	return &StructGenerator{w, pkgName}
}

func (s *StructGenerator) Generate(items []extractor.Struct) error {
	f := jen.NewFile(s.PkgName)
	f.HeaderComment("Code generated by go.elara.ws/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.")

	for _, item := range items {
		if len(item.UnionNames) > 0 {
			f.Type().Id(item.Name).String()

			f.Const().DefsFunc(func(g *jen.Group) {
				for _, member := range item.UnionNames {
					constName := strings.Replace(item.Name+string(member), " ", "", -1)
					g.Id(constName).Id(item.Name).Op("=").Lit(string(member))
				}
			})
		} else {
			f.Type().Id(item.Name).StructFunc(func(g *jen.Group) {
				for _, field := range item.Fields {
					g.Id(transformFieldName(field.Name)).Id(getType(field)).Tag(map[string]string{
						"json": field.Name,
						"url":  field.Name + ",omitempty",
					})
				}

				if strings.HasSuffix(item.Name, "Response") {
					g.Id("LemmyResponse")
				}
			})
		}
	}

	return f.Render(s.w)
}

func getType(f extractor.Field) string {
	t := transformType(f.Name, f.Type)
	if f.IsArray {
		t = "[]" + t
	}
	if f.IsOptional {
		t = "Optional[" + t + "]"
	}
	return t
}

func transformType(name, t string) string {
	// Some time fields are strings in the JS client,
	// use LemmyTime for those
	switch name {
	case "published", "updated":
		return "LemmyTime"
	}

	switch t {
	case "number":
		return "int64"
	case "boolean":
		return "bool"
	default:
		return t
	}
}

func transformFieldName(s string) string {
	s = snakeToCamel(s)
	s = strings.NewReplacer(
		"Id", "ID",
		"Url", "URL",
		"Nsfw", "NSFW",
		"Jwt", "JWT",
		"Crud", "CRUD",
		"Pm", "PM",
		"Totp", "TOTP",
		"2fa", "2FA",
	).Replace(s)
	return s
}

func snakeToCamel(s string) string {
	sb := &strings.Builder{}
	capitalizeNext := true
	for _, char := range s {
		if char == '_' {
			capitalizeNext = true
			continue
		}

		if capitalizeNext {
			sb.WriteRune(unicode.ToUpper(char))
			capitalizeNext = false
		} else {
			sb.WriteRune(char)
		}
	}
	return sb.String()
}
