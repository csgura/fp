package show

import (
	"fmt"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
	"github.com/csgura/fp/seq"
)

type Appender func(buf []string, opt fp.ShowOption) []string

func appendSeq(buf []string, typeName string, itr fp.Iterator[Appender], opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	apdseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && apdseq.IsEmpty() {
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

	showseq := seq.Map(apdseq, func(v Appender) []string {
		return v(nil, childOpt)
	})

	if opt.Indent != "" && showseq.Exists(func(v []string) bool {
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

		if showseq.IsEmpty() {
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

func appendMap(buf []string, typeName string, itr fp.Iterator[[]string], opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	showseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && showseq.IsEmpty() {
		return nil
	}

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

		if showseq.IsEmpty() {
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

func AppendStruct(buf []string, typeName string, opt fp.ShowOption, fields ...fp.Entry[Appender]) []string {

	childOpt := opt.IncreaseIndent()

	itr := iterator.Map(iterator.FromSeq(fields), func(t fp.Entry[Appender]) []string {
		valuestr := t.I2(nil, childOpt)
		if isEmptyString(valuestr) {
			return nil
		}
		return append([]string{quoteNames(t.I1, opt), spaceAfterColon(opt)}, valuestr...)
	}).FilterNot(isZero)

	return appendMap(buf, typeName, itr, opt)

}

func StructAppender(typeName string, fields ...fp.Entry[Appender]) Appender {
	return func(buf []string, opt fp.ShowOption) []string {
		childOpt := opt.IncreaseIndent()

		itr := iterator.Map(iterator.FromSeq(fields), func(t fp.Entry[Appender]) []string {
			valuestr := t.I2(nil, childOpt)
			if isEmptyString(valuestr) {
				return nil
			}
			return append([]string{quoteNames(t.I1, opt), spaceAfterColon(opt)}, valuestr...)
		}).FilterNot(isZero)

		return appendMap(buf, typeName, itr, opt)
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

func AppendSpaceAfterComma(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceAfterComma(opt))
}

func AppendSpaceAfterColon(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceAfterColon(opt))
}

func AppendSpaceWithinBrace(buf []string, opt fp.ShowOption) []string {
	return append(buf, spaceWithinBrace(opt))
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
