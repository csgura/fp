package show

import (
	"fmt"
	"strings"

	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"github.com/csgura/fp/iterator"
)

type Appender func(buf []string, opt fp.ShowOption) []string

func appendSeq(buf []string, typeName string, itr fp.Iterator[[]string], opt fp.ShowOption) []string {
	childOpt := opt.IncreaseIndent()

	showseq := as.Seq(itr.ToSeq())
	if opt.OmitEmpty && showseq.IsEmpty() {
		return nil
	}
	if opt.Indent != "" && showseq.Exists(func(v []string) bool {
		return as.Seq(v).Exists(fp.Test(as.Func2(strings.Contains), "\n"))
	}) {
		return append(
			append(
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), arrayOpen(opt), "\n", childOpt.CurrentIndent()),
				makeString(showseq, ",\n"+childOpt.CurrentIndent())...,
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
				append(buf, omitTypeName(typeName, opt), spaceBetweenTypeAndBrace(opt), "{\n", childOpt.CurrentIndent()),
				makeString(showseq, ",\n"+childOpt.CurrentIndent())...,
			),
			trailingComma(opt), "\n", opt.CurrentIndent(), "}",
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

func FormatStruct(typeName string, opt fp.ShowOption, fields ...fp.Tuple2[string, Appender]) string {
	return strings.Join(AppendStruct(nil, typeName, opt, fields...), "")
}

func AppendStruct(buf []string, typeName string, opt fp.ShowOption, fields ...fp.Tuple2[string, Appender]) []string {

	childOpt := opt.IncreaseIndent()

	itr := iterator.Map(iterator.FromSeq(fields), func(t fp.Tuple2[string, Appender]) []string {
		valuestr := t.I2(nil, childOpt)
		if isEmptyString(valuestr) {
			return nil
		}
		return append([]string{t.I1, spaceAfterColon(opt)}, valuestr...)
	}).FilterNot(isZero)

	return appendMap(buf, typeName, itr, opt)

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
