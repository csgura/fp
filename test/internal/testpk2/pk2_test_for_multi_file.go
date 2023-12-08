package testpk2

// @fp.Value
type NotIgnored struct {
	ig int
}

//lint:file-ignore U1000 test code

type Ignored struct {
	ig int
}
