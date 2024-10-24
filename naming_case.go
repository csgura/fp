package fp

import (
	"strings"
	"unicode"
)

type NamingCase string

const (
	CamelCase  NamingCase = "camelCase"
	PascalCase NamingCase = "PascalCase"
	SnakeCase  NamingCase = "snake_case"
	KebabCase  NamingCase = "kebab-case"
)

func firstLower(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToLower(name[:1]) + name[1:]
}

func firstUpper(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[:1]) + name[1:]
}

func splitSnakeAndKebab(name string) Iterator[string] {
	return IteratorOfSeq(strings.Split(name, "_")).FlatMap(func(v string) Iterator[string] {
		return IteratorOfSeq(strings.Split(v, "-"))
	})
}

func foldSlice[A, B any](s []A, zero B, f func(B, A) B) B {
	sum := zero
	for _, v := range s {
		sum = f(sum, v)
	}
	return sum
}

func splitUpper(name string) Iterator[string] {

	type buffer struct {
		list    []string
		builder *strings.Builder
	}

	buf := foldSlice([]rune(name), buffer{builder: &strings.Builder{}}, func(b buffer, a rune) buffer {
		if unicode.IsUpper(a) && b.builder.Len() > 0 {
			nl := append(b.list, b.builder.String())
			b.builder.Reset()
			b.builder.WriteRune(a)
			return buffer{
				list:    nl,
				builder: b.builder,
			}
		}
		b.builder.WriteRune(a)
		return b
	})
	ret := append(buf.list, buf.builder.String())
	return IteratorOfSeq(ret)
}

func ConvertNaming(name string, namingCase NamingCase) string {
	if name == "" {
		return ""
	}

	switch namingCase {
	case CamelCase:
		return firstLower(splitSnakeAndKebab(name).Map(firstUpper).MakeString(""))
	case PascalCase:
		return splitSnakeAndKebab(name).Map(firstUpper).MakeString("")
	case SnakeCase:
		return splitSnakeAndKebab(name).FlatMap(splitUpper).Map(strings.ToLower).MakeString("_")
	case KebabCase:
		return splitSnakeAndKebab(name).FlatMap(splitUpper).Map(strings.ToLower).MakeString("-")
	default:
		return name
	}
}
