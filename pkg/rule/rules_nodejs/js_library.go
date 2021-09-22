package rules_nodejs

import (
	"sort"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/bazelbuild/bazel-gazelle/rule"

	"github.com/stackb/rules_proto/pkg/protoc"
)

var jsLibraryKindInfo = rule.KindInfo{
	MergeableAttrs: map[string]bool{
		"srcs":       true,
		"deps":       true,
		"visibility": true,
	},
}

// JsLibrary implements RuleProvider for 'cc_library'-derived rules.
type JsLibrary struct {
	KindName       string
	RuleNameSuffix string
	Outputs        []string
	Config         *protoc.ProtocConfiguration
	RuleConfig     *protoc.LanguageRuleConfig
	Resolver       protoc.DepsResolver
}

// Kind implements part of the ruleProvider interface.
func (s *JsLibrary) Kind() string {
	return s.KindName
}

// Name implements part of the ruleProvider interface.
func (s *JsLibrary) Name() string {
	return s.Config.Library.BaseName() + s.RuleNameSuffix
}

// Srcs computes the srcs list for the rule.
func (s *JsLibrary) Srcs() []string {
	srcs := make([]string, 0)
	for _, output := range s.Outputs {
		if strings.HasSuffix(output, ".js") {
			srcs = append(srcs, protoc.StripRel(s.Config.Rel, output))
		}
	}
	return srcs
}

// Deps computes the deps list for the rule.
func (s *JsLibrary) Deps() []string {
	return s.RuleConfig.GetDeps()
}

// Visibility implements part of the ruleProvider interface.
func (s *JsLibrary) Visibility() []string {
	visibility := make([]string, 0)
	for k, want := range s.RuleConfig.Visibility {
		if !want {
			continue
		}
		visibility = append(visibility, k)
	}
	sort.Strings(visibility)
	return visibility
}

// Rule implements part of the ruleProvider interface.
func (s *JsLibrary) Rule() *rule.Rule {
	newRule := rule.NewRule(s.Kind(), s.Name())

	newRule.SetAttr("srcs", s.Srcs())

	visibility := s.Visibility()
	if len(visibility) > 0 {
		newRule.SetAttr("visibility", visibility)
	}

	return newRule
}

// Resolve implements part of the RuleProvider interface.
func (s *JsLibrary) Resolve(c *config.Config, r *rule.Rule, importsRaw interface{}, from label.Label) {
	if s.Resolver == nil {
		return
	}
	s.Resolver(s, s.Config, c, r, importsRaw, from)
}