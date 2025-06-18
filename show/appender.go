package show

import (
	"fmt"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/option"
	"github.com/csgura/fp/seq"
	"github.com/csgura/fp/slice"
)

type Appender func(buf []string, opt fp.ShowOption) []string

func appendSeq(buf []string, typeName string, apdseq []Appender, opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	if opt.OmitEmpty && len(apdseq) == 0 {
		return nil
	}

	if opt.OmitObjectBrace {
		showseq := seq.Map(apdseq, func(v Appender) []string {
			return v(nil, opt)
		})

		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), omitBrace("\n", opt), opt.CurrentIndent()),
				makeArrayString(showseq, "- ", structFieldSeparator(opt))...,
			),
			trailingComma(opt), omitBrace("\n", opt), omitBrace(opt.CurrentIndent(), opt), arrayClose(opt),
		)
	}

	showseq := slice.Map(apdseq, func(v Appender) []string {
		return v(nil, childOpt)
	})

	if opt.Indent != "" && slice.Exists(showseq, func(v []string) bool {
		return as.Seq(v).Exists(fp.Test(as.Func2(strings.Contains), "\n"))
	}) {

		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), omitBrace("\n", opt), childOpt.CurrentIndent()),
				makeString(showseq, structFieldSeparator(childOpt))...,
			),
			trailingComma(opt), "\n", opt.CurrentIndent(), arrayClose(opt),
		)
		//		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {

		if slice.IsEmpty(showseq) {
			return append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), arrayClose(opt))
		}

		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), spaceWithinBrace(opt)),
				makeString(showseq, spaceAfterComma(opt))...,
			),
			spaceWithinBrace(opt), arrayClose(opt),
		)
	}
}

func appendMap(buf []string, typeName string, apdseq []Appender, opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	if opt.OmitEmpty && slice.IsEmpty(apdseq) {
		return nil
	}

	showseq := slice.Map(apdseq, func(v Appender) []string {
		return v(nil, childOpt)
	})

	if opt.Indent != "" {
		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), omitBrace("{\n", opt), childOpt.CurrentIndent()),
				makeString(showseq, structFieldSeparator(childOpt))...,
			),
			trailingComma(opt), omitBrace("\n", opt), omitBrace(opt.CurrentIndent(), opt), omitBrace("}", opt),
		)
		//		return fmt.Sprintf("%s {\n%s%s\n%s}", typeName, childOpt.CurrentIndent(), showseq.MakeString(",\n"+childOpt.CurrentIndent()), opt.CurrentIndent())
	} else {

		if slice.IsEmpty(showseq) {
			return append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), " {}")
		}

		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), "{", spaceWithinBrace(opt)),
				makeString(showseq, spaceAfterComma(opt))...,
			),
			spaceWithinBrace(opt), "}",
		)
	}
}

func FormatStruct(typeName string, opt fp.ShowOption, fields ...fp.Entry[Appender]) string {
	return strings.Join(AppendStruct(nil, typeName, opt, fields...), "")
}

func AppendCommaSperated(buf []string, list []Appender, opt fp.ShowOption) []string {
	showseq := slice.FilterNot(slice.Map(list, func(v Appender) []string {
		return v(nil, opt)
	}), isEmptyString)
	return append(buf, makeString(showseq, structFieldSeparator(opt))...)
}

func AppendStruct(buf []string, typeName string, opt fp.ShowOption, fields ...fp.Entry[Appender]) []string {

	childOpt := opt.IncreaseIndent()

	itr := iterator.FilterMap(iterator.FromSeq(fields), func(t fp.Entry[Appender]) fp.Option[Appender] {
		valuestr := t.I2(nil, childOpt)
		if isEmptyString(valuestr) {
			return option.None[Appender]()
		}
		return option.Some[Appender](func(buf []string, opt fp.ShowOption) []string {
			return append([]string{quoteNames(t.I1, opt), spaceAfterColon(opt)}, valuestr...)
		})
	}).ToSeq()

	return appendMap(buf, typeName, itr, opt)

}

func StructAppender(typeName string, fields ...fp.Entry[Appender]) Appender {
	return func(buf []string, opt fp.ShowOption) []string {
		return AppendStruct(buf, typeName, opt, fields...)
	}
}

func AsAppender[T any](tshow fp.Show[T], t T) Appender {
	return func(buf []string, opt fp.ShowOption) []string {
		return tshow.Append(buf, t, opt)
	}
}

func StringAppender(str string) Appender {
	return AsAppender(String, str)
}

func NumberAppender[T fp.ImplicitNum](v T) Appender {
	return AsAppender(Number[T](), v)
}

func StringerAppender[T fmt.Stringer](v T) Appender {
	return AsAppender(Given[T](), v)
}

func AppendNil(buf []string, opt fp.ShowOption) []string {
	return append(buf, nullForNil(opt))
}

func AppendSpaceBetweenTypeAndBrace(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceBetweenTypeAndBrace(opt))
}

func AppendStringLiteral(buf []string, literal string, opt fp.ShowOption) []string {
	return String.Append(buf, literal, opt)
}

func AppendComma(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceAfterComma(opt))
}

func AppendColon(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceAfterColon(opt))
}

func AppendSpaceWithinBrace(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceWithinBrace(opt))
}

func AppendSpaceBeforeHCons(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceBeforeHCons(opt))
}

func AppendSpaceAfterHCons(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceAfterHCons(opt))
}

func AppendTypeName(buf []string, typeName string, opt fp.ShowOption) []string {
	return append(buf, omitTypeName(typeName, opt))
}

func AppendFieldName(buf []string, fieldName string, opt fp.ShowOption) []string {
	return append(buf, quoteNames(fieldName, opt))
}

func AppendArrayBegin(buf []string, opt fp.ShowOption) []string {
	return append(buf, arrayOpen(opt))
}

func AppendArrayEnd(buf []string, opt fp.ShowOption) []string {
	return append(buf, arrayClose(opt))
}

func AppendTrailingComma(buf []string, opt fp.ShowOption) []string {
	return append(buf, trailingComma(opt))
}

func AppendNewLineWithIndent(buf []string, opt fp.ShowOption) []string {
	return append(buf, "\n", opt.CurrentIndent())
}

func AppendSlice(buf []string, typeName string, sl []Appender, opt fp.ShowOption) []string {
	return appendSeq(buf, typeName, sl, opt)
}

func AppendMap(buf []string, typeName string, sl []Appender, opt fp.ShowOption) []string {
	return appendMap(buf, typeName, sl, opt)
}

func AppendGeneric(buf []string, name string, kind string, reprAppend Appender, opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()
	valueStr := reprAppend(nil, childOpt)
	if opt.OmitEmpty && isEmptyString(valueStr) {
		return nil
	}

	if kind == fp.GenericKindConversion {
		return reprAppend(buf, opt)
	}
	if kind == fp.GenericKindNewType {
		if opt.OmitTypeName {
			return append(buf, valueStr...)
		}
		return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt),
			omitBrace("(", opt), spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), omitBrace(")", opt))
	} else if kind == fp.GenericKindTuple {
		if opt.SquareBracketForArray {
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt),
				omitBrace("[", opt), spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), omitBrace("]", opt))

		} else {
			return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), omitBrace("(", opt), spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), omitBrace(")", opt))

		}
	}

	if opt.Indent != "" {
		return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), omitBrace("{\n", opt), childOpt.CurrentIndent()), valueStr...), trailingComma(opt), omitBrace("\n", opt), omitBrace(opt.CurrentIndent(), opt), omitBrace("}", opt))
	} else {
		return append(append(append(buf, omitTypeName(name, opt), spaceBetweenTypeAndBrace(opt), "{", spaceWithinBrace(opt)), valueStr...), spaceWithinBrace(opt), "}")
	}
}

// var NullAppender = Appender(func(buf []string, opt fp.ShowOption) []string {
// 	return append(buf, nullForNil(opt))
// })

// var ColonAppender = Appender(func(buf []string, opt fp.ShowOption) []string {
// 	return append(buf, spaceAfterColon(opt))
// })

// var NameAppender = Appender(func(buf []string, opt fp.ShowOption) []string {
// 	return append(buf, spaceAfterColon(opt))
// })
