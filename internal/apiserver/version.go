package apiserver

import (
	pb "github.com/samsung-cnct/ims-kaas/pkg/generated/api"
	"golang.org/x/net/context"

	"github.com/samsung-cnct/ims-kaas/pkg/version"
)

func (s *Server) GetVersionInformation(ctx context.Context, in *pb.GetVersionMsg) (*pb.GetVersionReply, error) {
	versionInformation := version.Get()
	reply := &pb.GetVersionReply{
		Ok: true,
		VersionInformation: &pb.GetVersionReply_VersionInformation{
			GitCommit: versionInformation.GitCommit,
			BuildDate: versionInformation.BuildDate,
			GoVersion: versionInformation.GoVersion,
			Compiler:  versionInformation.Compiler,
			Platform:  versionInformation.Platform,
		},
	}
	return reply, nil
}
