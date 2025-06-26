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
	HeaderCase NamingCase = "Header-Case"
)

func spanSlice[T any](s []T, p func(T) bool) ([]T, []T) {
	left := []T{}
	right := []T{}

	span := false
	for _, v := range s {
		if span {
			right = append(right, v)
		} else {
			if p(v) {
				left = append(left, v)
			} else {
				span = true
				right = append(right, v)
			}
		}
	}

	return left, right
}

// func firstLower(name string) string {
// 	if name == "" {
// 		return ""
// 	}
// 	return strings.ToLower(name[:1]) + name[1:]
// }

func firstSeriesLower(name string) string {
	l, r := spanSlice([]rune(name), unicode.IsUpper)
	return strings.ToLower(string(l)) + string(r)
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
		list         []string
		builder      *strings.Builder
		includeLower bool
	}

	buf := foldSlice([]rune(name), buffer{builder: &strings.Builder{}}, func(b buffer, a rune) buffer {
		if unicode.IsUpper(a) {
			if b.includeLower {
				nl := append(b.list, b.builder.String())
				b.builder.Reset()
				b.builder.WriteRune(a)
				return buffer{
					list:    nl,
					builder: b.builder,
				}
			} else {
				b.builder.WriteRune(a)
			}
		} else {
			b.includeLower = true
			b.builder.WriteRune(a)
		}

		return b
	})
	ret := append(buf.list, buf.builder.String())
	return IteratorOfSeq(ret)
}

func runeAt(v string, idx int) Option[rune] {
	if idx < len(v) {
		return Some(([]rune(v))[idx])
	}
	return None[rune]()
}

func isUpper(v string, idx int) bool {
	r := runeAt(v, idx)
	if r.IsDefined() {
		return unicode.IsUpper(r.Get())
	}
	return false
}

func CheckNaming(name string) Option[NamingCase] {

	none := None[NamingCase]()

	if name == "" {
		return none
	}

	if strings.Contains(name, "-") {
		if strings.Contains(name, "_") {
			return none
		}
		if isUpper(name, 0) {
			return Some(HeaderCase)
		}
		return Some(KebabCase)
	} else if strings.Contains(name, "_") {
		return Some(SnakeCase)
	} else if isUpper(name, 0) {
		return Some(PascalCase)
	}
	return Some(CamelCase)
}

func ConvertNaming(name string, namingCase NamingCase) string {
	if name == "" {
		return ""
	}

	inferred := CheckNaming(name)

	switch namingCase {
	case CamelCase:
		switch inferred {
		case Some(KebabCase), Some(SnakeCase), Some(HeaderCase):
			return firstSeriesLower(splitSnakeAndKebab(name).Map(firstUpper).MakeString(""))
		case Some(PascalCase):
			return firstSeriesLower(name)
		default:
			return name
		}
	case PascalCase:
		switch inferred {
		case Some(KebabCase), Some(SnakeCase), Some(HeaderCase):
			return splitSnakeAndKebab(name).Map(firstUpper).MakeString("")
		case Some(CamelCase):
			return firstUpper(name)
		default:
			return name
		}
	case SnakeCase:
		return splitSnakeAndKebab(name).FlatMap(splitUpper).Map(strings.ToLower).MakeString("_")
	case KebabCase:
		return splitSnakeAndKebab(name).FlatMap(splitUpper).Map(strings.ToLower).MakeString("-")
	case HeaderCase:
		return splitSnakeAndKebab(name).FlatMap(splitUpper).Map(firstUpper).MakeString("-")
	default:
		return name
	}
}
