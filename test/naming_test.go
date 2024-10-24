package main_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
	"github.com/csgura/fp/option"
)

func TestNaming(t *testing.T) {

	assert.Equal(fp.CheckNaming("HelloWorldTowerCrain"), option.Some(fp.PascalCase))
	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.CamelCase), "helloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.CamelCase), "helloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.CamelCase), "helloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("Hello-World-Tower-Crain", fp.CamelCase), "helloWorldTowerCrain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.PascalCase), "HelloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.PascalCase), "HelloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.PascalCase), "HelloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("Hello-World-Tower-Crain", fp.PascalCase), "HelloWorldTowerCrain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.SnakeCase), "hello_world_tower_crain")
	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.SnakeCase), "hello_world_tower_crain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.SnakeCase), "hello_world_tower_crain")
	assert.Equal(fp.ConvertNaming("Hello-World-Tower-Crain", fp.SnakeCase), "hello_world_tower_crain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.KebabCase), "hello-world-tower-crain")
	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.KebabCase), "hello-world-tower-crain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.KebabCase), "hello-world-tower-crain")
	assert.Equal(fp.ConvertNaming("Hello-World-Tower-Crain", fp.KebabCase), "hello-world-tower-crain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.HeaderCase), "Hello-World-Tower-Crain")
	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.HeaderCase), "Hello-World-Tower-Crain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.HeaderCase), "Hello-World-Tower-Crain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.HeaderCase), "Hello-World-Tower-Crain")

	assert.Equal(fp.ConvertNaming("PDUSessionChargingInformation", fp.HeaderCase), "PDUSession-Charging-Information")
	assert.Equal(fp.ConvertNaming("PDUSessionChargingInformation", fp.KebabCase), "pdusession-charging-information")
	assert.Equal(fp.ConvertNaming("PDUSessionChargingInformation", fp.CamelCase), "pdusessionChargingInformation")

	assert.Equal(fp.ConvertNaming("PDUSession-Charging-Information", fp.CamelCase), "pdusessionChargingInformation")
	assert.Equal(fp.ConvertNaming("PDUSession-Charging-Information", fp.PascalCase), "PDUSessionChargingInformation")

}
