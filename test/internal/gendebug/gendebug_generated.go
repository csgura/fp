// Code generated by gombok, DO NOT EDIT.
package gendebug

type InvokerAdaptor struct {
	DoInvoke func(a1 any)
}

func (r *InvokerAdaptor) Invoke(a1 any) {

	if r.DoInvoke != nil {
		r.DoInvoke(a1)
		return
	}

	panic("InvokerAdaptor.Invoke not implemented")
}
