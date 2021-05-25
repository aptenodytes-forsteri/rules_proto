package builtin_test

import (
	"testing"

	"github.com/stackb/rules_proto/pkg/builtin"
	"github.com/stackb/rules_proto/pkg/plugintest"
)

func TestCppPlugin(t *testing.T) {
	plugintest.Cases(t, &builtin.CppPlugin{}, map[string]plugintest.Case{
		"empty file": {
			Input: "",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("test.pb.cc", "test.pb.h"),
			),
		},
		"only services": {
			Input: "service S{}",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("test.pb.cc", "test.pb.h"),
			),
		},
		"with a package": {
			Input: "package pkg;\n\nmessage M{}",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("test.pb.cc", "test.pb.h"),
			),
		},
		"in a relative directory": {
			Rel:   "rel",
			Input: "package a;\n\nmessage M{}",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("rel/test.pb.cc", "rel/test.pb.h"),
			),
		},
		"snake_case": {
			Basename: "snake_case",
			Input:    "message M{}",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("snake_case.pb.cc", "snake_case.pb.h"),
			),
		},
		"PascalCase": {
			Basename: "PascalCase",
			Input:    "message M{}",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("PascalCase.pb.cc", "PascalCase.pb.h"),
			),
		},
		"camelCase": {
			Basename: "camelCase",
			Input:    "message M{}",
			Directives: plugintest.WithDirectives(
				"proto_plugin", "cpp implementation builtin:cpp",
			),
			Configuration: plugintest.WithConfiguration(
				plugintest.WithName("cpp"),
				plugintest.WithOutputs("camelCase.pb.cc", "camelCase.pb.h"),
			),
		},
	})
}
