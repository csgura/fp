package main_test

import (
	"testing"

	"github.com/csgura/fp"
	"github.com/csgura/fp/internal/assert"
)

func TestNaming(t *testing.T) {

	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.CamelCase), "helloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.CamelCase), "helloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.CamelCase), "helloWorldTowerCrain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.PascalCase), "HelloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.PascalCase), "HelloWorldTowerCrain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.PascalCase), "HelloWorldTowerCrain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.SnakeCase), "hello_world_tower_crain")
	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.SnakeCase), "hello_world_tower_crain")
	assert.Equal(fp.ConvertNaming("hello-world-tower-crain", fp.SnakeCase), "hello_world_tower_crain")

	assert.Equal(fp.ConvertNaming("helloWorldTowerCrain", fp.KebabCase), "hello-world-tower-crain")
	assert.Equal(fp.ConvertNaming("HelloWorldTowerCrain", fp.KebabCase), "hello-world-tower-crain")
	assert.Equal(fp.ConvertNaming("hello_world_tower_crain", fp.KebabCase), "hello-world-tower-crain")
}
