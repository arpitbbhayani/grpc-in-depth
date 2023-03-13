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
		Text: "Github is where developers code, collaborate, and ship.",
	}, nil
}
