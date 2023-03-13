package infosvc_server

import (
	context "context"

	infosvc_proto "github.com/arpitbbhayani/grpc-in-depth/infosvc/proto"
)

type InfoSvcServerV1 struct {
	infosvc_proto.UnimplementedInfoSvcServer
}

func (InfoSvcServerV1) WhatIsGithub(context.Context, *infosvc_proto.WhatIsGithubRequest) (*infosvc_proto.WhatIsGithubResponse, error) {
	return &infosvc_proto.WhatIsGithubResponse{
		Text: "GitHub is a platform for developers to store, manage, and share their code repositories.",
	}, nil
}
