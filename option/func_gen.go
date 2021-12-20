package option

import (
	"github.com/csgura/fp"
)

func Compose3[A1, A2, A3, R any](f1 fp.Func1[A1, fp.Option[A2]], f2 fp.Func1[A2, fp.Option[A3]], f3 fp.Func1[A3, fp.Option[R]]) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose2(f2, f3))
}

func Compose4[A1, A2, A3, A4, R any](f1 fp.Func1[A1, fp.Option[A2]], f2 fp.Func1[A2, fp.Option[A3]], f3 fp.Func1[A3, fp.Option[A4]], f4 fp.Func1[A4, fp.Option[R]]) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose3(f2, f3, f4))
}

func Compose5[A1, A2, A3, A4, A5, R any](f1 fp.Func1[A1, fp.Option[A2]], f2 fp.Func1[A2, fp.Option[A3]], f3 fp.Func1[A3, fp.Option[A4]], f4 fp.Func1[A4, fp.Option[A5]], f5 fp.Func1[A5, fp.Option[R]]) fp.Func1[A1, fp.Option[R]] {
	return Compose2(f1, Compose4(f2, f3, f4, f5))
}
