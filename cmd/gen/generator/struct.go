package generator

import (
	"io"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/rystaf/go-lemmy/cmd/gen/extractor"
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
					g.Id(transformFieldName(field.Name)).Add(getType(field)).Tag(map[string]string{
						"json": field.Name,
						"url":  field.Name + ",omitempty",
					})
				}

				if strings.HasSuffix(item.Name, "Response") {
					g.Id("Error").Id("Optional").Types(jen.String()).Tag(map[string]string{"json": "error"})
				}
			})
		}
	}

	return f.Render(s.w)
}

func getType(f extractor.Field) jen.Code {
	// Some time fields are strings in the JS client,
	// use time.Time for those
	switch f.Name {
	case "published", "updated", "when_":
		return jen.Qual("time", "Time")
	}

	// Rank types such as hot_rank and hot_rank_active may be floats.
	if strings.Contains(f.Name, "rank") {
		return jen.Float64()
	}

	t := transformType(f.Name, f.Type)
	if f.IsArray {
		t = "[]" + t
	}
	if f.IsOptional {
		t = "Optional[" + t + "]"
	}
	return jen.Id(t)
}

func transformType(name, t string) string {
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
		"Png", "PNG",
		"Uuid", "UUID",
		"Wav", "WAV",
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
