// +build ignore

package main

import (
	"text/template"

	"src.sourcegraph.com/sourcegraph/gen"
)

func main() {
	svcs := []string{
		"../../../../go-sourcegraph/sourcegraph/sourcegraph.pb.go",
		"../../../../Godeps/_workspace/src/sourcegraph.com/sourcegraph/srclib/store/pb/srcstore.pb.go",
		"../../../../gitserver/gitpb/git_transport.pb.go",
	}
	gen.Generate("outer_middleware.go", tmpl, svcs, nil, "")
}

var tmpl = template.Must(template.New("").Delims("<<<", ">>>").Parse(`// GENERATED CODE - DO NOT EDIT!
//
// Generated by:
//
//   go run gen_middleware.go
//
// Called via:
//
//   go generate
//

package outer

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sourcegraph.com/sourcegraph/go-vcs/vcs"
	"sourcegraph.com/sourcegraph/grpccache"
	"sourcegraph.com/sourcegraph/srclib/store/pb"
	"sourcegraph.com/sourcegraph/srclib/unit"
	"sourcegraph.com/sqs/pbtypes"
	"src.sourcegraph.com/sourcegraph/gitserver/gitpb"
	"src.sourcegraph.com/sourcegraph/go-sourcegraph/sourcegraph"
	"src.sourcegraph.com/sourcegraph/pkg/inventory"
	"src.sourcegraph.com/sourcegraph/svc"
)

// Services returns a full set of services with an implementation of each service method that lets you customize the initial context.Context and map Go errors to gRPC error codes. It is similar to HTTP handler middleware, but for gRPC servers.
func Services(ctxFunc ContextFunc, services svc.Services) svc.Services {
	s := svc.Services{
		<<<range .>>><<<.Name>>>: wrapped<<<.Name>>>{ctxFunc, services},
		<<<end>>>
	}
	return s
}

<<<range .>>>
	type wrapped<<<.Name>>> struct{
		ctxFunc ContextFunc
		services svc.Services
	}

  <<<$service := .>>>
	<<<range .Methods>>>
		func (s wrapped<<<$service.Name>>>) <<<.Name>>>(ctx context.Context, v1 *<<<.ParamType>>>) (*<<<.ResultType>>>, error) {
			var cc *grpccache.CacheControl
			ctx, cc = grpccache.Internal_WithCacheControl(ctx)

			var err error
			ctx, err = initContext(ctx, s.ctxFunc, s.services)
			if err != nil {
				return nil, wrapErr(err)
			}

			innerSvc := svc.<<<$service.Name>>>OrNil(ctx)
			if innerSvc == nil {
				return nil, grpc.Errorf(codes.Unimplemented, "<<<$service.Name>>>")
			}

			rv, err := innerSvc.<<<.Name>>>(ctx, v1)
			if err != nil {
				return nil, wrapErr(err)
			}

			if !cc.IsZero() {
				if err := grpccache.Internal_SetCacheControlTrailer(ctx, *cc); err != nil {
					return nil, err
				}
			}

			return rv, nil
		}
	<<<end>>>
<<<end>>>
`))
