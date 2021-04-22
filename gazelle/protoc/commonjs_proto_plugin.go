package protoc

func init() {
	MustRegisterProtoPlugin("common_js_proto", &CommonJsProtoPlugin{})
}

// CommonJsProtoPlugin implements ProtoPlugin for the built-in js plugin with
// commonjs option.
type CommonJsProtoPlugin struct{}

// ShouldApply implements part of the ProtoPlugin interface.
func (p *CommonJsProtoPlugin) ShouldApply(rel string, cfg *ProtoPackageConfig, lib ProtoLibrary) bool {
	for _, f := range lib.Files() {
		if f.HasMessages() || f.HasEnums() {
			return true
		}
	}
	return false
}

// GeneratedSrcs implements part of the ProtoPlugin interface.
func (p *CommonJsProtoPlugin) GeneratedSrcs(rel string, cfg *ProtoPackageConfig, lib ProtoLibrary) []string {
	srcs := make([]string, 0)
	for _, f := range lib.Files() {
		if f.HasMessages() || f.HasEnums() {
			srcs = append(srcs, f.Name+"_pb.js")
		}
	}
	return srcs
}