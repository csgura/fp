// Code generated by gombok, DO NOT EDIT.
package genintf

import (
	"context"
	"github.com/csgura/fp"
)

// @fp.Getter
// @fp.AllArgsConstructor
type MessageUniverse struct {
	ctx    context.Context
	galaxy string
	ResponseType[string]
}

func (r handler) Universe(ctx context.Context, galaxy string) fp.Try[string] {
	return NewMessageUniverse(ctx, galaxy).SendRequest(r.ref, r.timeout)
}

// @fp.Getter
// @fp.AllArgsConstructor
type MessageWorld struct {
	address string
	count   int
	ResponseType[string]
}

func (r handler) World(address string, count int) fp.Try[string] {
	return NewMessageWorld(address, count).SendRequest(r.ref, r.timeout)
}

// @fp.Getter
// @fp.AllArgsConstructor
type MessageToday struct {
	ResponseType[string]
}

func (r handler) Today() fp.Try[string] {
	return NewMessageToday().SendRequest(r.ref, r.timeout)
}

// World HasMethod Size :true
// World HasMethod Head :false

// World has method Size
