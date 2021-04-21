load(
    "@build_stack_rules_proto//rules:proto_rules.bzl",
    "proto_compile_aspect",
    "proto_compile_aspect_rule_macro",
    "proto_compile_aspect_rule",
    "proto_compile_rule",
)

_default_plugins = [
    str(Label("//plugins/gogo:gostring")),
]

_gostring_proto_compile_aspect = proto_compile_aspect(_default_plugins, "gostring_proto_compile_aspect")

_gostring_proto_compile_aspect_rule = proto_compile_aspect_rule(_gostring_proto_compile_aspect)

_gostring_proto_compile_rule = proto_compile_rule(_default_plugins)

def gostring_proto_compile(**kwargs):
    if kwargs.pop("transitive", False):
        proto_compile_aspect_rule_macro(_gostring_proto_compile_aspect_rule, **kwargs)
    else:
        _gostring_proto_compile_rule(**kwargs)