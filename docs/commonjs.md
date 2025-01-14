---
layout: default
title: commonjs
permalink: examples/commonjs
parent: Examples
---


# commonjs example

`bazel test //example/golden:commonjs_test`


## `BUILD.bazel` (after gazelle)

~~~python
load("@rules_proto//proto:defs.bzl", "proto_library")
load("@build_stack_rules_proto//rules:proto_compile.bzl", "proto_compile")
load("@rules_proto//proto:defs.bzl", "proto_library")

# "proto_rule" instantiates the proto_compile rule
# gazelle:proto_rule proto_compile implementation stackb:rules_proto:proto_compile

# "proto_plugin" instantiates the builtin commonjs plugin
# gazelle:proto_plugin commonjs implementation builtin:js:common

# "proto_language" binds the rule(s) and plugin(s) together
# gazelle:proto_language commonjs rule proto_compile
# gazelle:proto_language commonjs plugin commonjs

proto_library(
    name = "example_proto",
    srcs = ["example.proto"],
    visibility = ["//visibility:public"],
)

proto_compile(
    name = "example_commonjs_compile",
    options = {"@build_stack_rules_proto//plugin/builtin:commonjs": ["import_style=commonjs"]},
    outputs = ["example_pb.js"],
    plugins = ["@build_stack_rules_proto//plugin/builtin:commonjs"],
    proto = "example_proto",
)
~~~


## `BUILD.bazel` (before gazelle)

~~~python
# "proto_rule" instantiates the proto_compile rule
# gazelle:proto_rule proto_compile implementation stackb:rules_proto:proto_compile

# "proto_plugin" instantiates the builtin commonjs plugin
# gazelle:proto_plugin commonjs implementation builtin:js:common

# "proto_language" binds the rule(s) and plugin(s) together
# gazelle:proto_language commonjs rule proto_compile
# gazelle:proto_language commonjs plugin commonjs
~~~


## `WORKSPACE`

~~~python
~~~

