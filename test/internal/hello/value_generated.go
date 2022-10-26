// Code generated by gombok, DO NOT EDIT.
package hello

import (
	"encoding/json"
	"fmt"
	"github.com/csgura/fp"
	"github.com/csgura/fp/as"
	"net/http"
	"time"
)

type WorldBuilder World

type WorldMutable struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func (r WorldBuilder) Build() World {
	return World(r)
}

func (r World) Builder() WorldBuilder {
	return WorldBuilder(r)
}

func (r World) Message() string {
	return r.message
}

func (r World) WithMessage(v string) World {
	r.message = v
	return r
}

func (r WorldBuilder) Message(v string) WorldBuilder {
	r.message = v
	return r
}

func (r World) Timestamp() time.Time {
	return r.timestamp
}

func (r World) WithTimestamp(v time.Time) World {
	r.timestamp = v
	return r
}

func (r WorldBuilder) Timestamp(v time.Time) WorldBuilder {
	r.timestamp = v
	return r
}

func (r World) String() string {
	return fmt.Sprintf("World(message=%v, timestamp=%v)", r.message, r.timestamp)
}

func (r World) AsTuple() fp.Tuple2[string, time.Time] {
	return as.Tuple2(r.message, r.timestamp)
}

func (r World) AsMutable() WorldMutable {
	return WorldMutable{
		Message:   r.message,
		Timestamp: r.timestamp,
	}
}

func (r WorldMutable) AsImmutable() World {
	return World{
		message:   r.Message,
		timestamp: r.Timestamp,
	}
}

func (r WorldBuilder) FromTuple(t fp.Tuple2[string, time.Time]) WorldBuilder {
	r.message = t.I1
	r.timestamp = t.I2
	return r
}

func (r World) AsMap() map[string]any {
	return map[string]any{
		"message":   r.message,
		"timestamp": r.timestamp,
	}
}

func (r WorldBuilder) FromMap(m map[string]any) WorldBuilder {

	if v, ok := m["message"].(string); ok {
		r.message = v
	}

	if v, ok := m["timestamp"].(time.Time); ok {
		r.timestamp = v
	}

	return r
}

func (r World) AsLabelled() fp.Labelled2[string, time.Time] {
	return as.Labelled2(as.Field("message", r.message), as.Field("timestamp", r.timestamp))
}

func (r WorldBuilder) FromLabelled(t fp.Labelled2[string, time.Time]) WorldBuilder {
	r.message = t.I1.Value
	r.timestamp = t.I2.Value
	return r
}

func (r World) MarshalJSON() ([]byte, error) {
	m := r.AsMutable()
	return json.Marshal(m)
}

func (r *World) UnmarshalJSON(b []byte) error {
	if r == nil {
		return fp.Error(http.StatusBadRequest, "target ptr is nil")
	}
	m := r.AsMutable()
	err := json.Unmarshal(b, &m)
	if err == nil {
		*r = m.AsImmutable()
	}
	return err
}
