package protoc

func init() {
	MustRegisterProtoPlugin("cc_grpc", &CcGrpcPlugin{})
}

// CcGrpcPlugin implements ProtoPlugin for the built-in protoc python plugin.
type CcGrpcPlugin struct{}

// ShouldApply implements part of the ProtoPlugin interface.
func (p *CcGrpcPlugin) ShouldApply(rel string, cfg *ProtoPackageConfig, lib ProtoLibrary) bool {
	for _, f := range lib.Files() {
		if f.HasServices() {
			return true
		}
	}
	return false
}

// GeneratedSrcs implements part of the ProtoPlugin interface.
func (p *CcGrpcPlugin) GeneratedSrcs(rel string, cfg *ProtoPackageConfig, lib ProtoLibrary) []string {
	srcs := make([]string, 0)
	for _, f := range lib.Files() {
		if f.HasServices() {
			srcs = append(srcs, f.Name+"_pb_grpc.cc")
		}
	}
	return srcs
}