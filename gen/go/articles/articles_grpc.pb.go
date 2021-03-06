// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package articles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ArticlesClient is the client API for Articles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArticlesClient interface {
	ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error)
}

type articlesClient struct {
	cc grpc.ClientConnInterface
}

func NewArticlesClient(cc grpc.ClientConnInterface) ArticlesClient {
	return &articlesClient{cc}
}

func (c *articlesClient) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...grpc.CallOption) (*ListArticlesResponse, error) {
	out := new(ListArticlesResponse)
	err := c.cc.Invoke(ctx, "/Articles/ListArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArticlesServer is the server API for Articles service.
// All implementations must embed UnimplementedArticlesServer
// for forward compatibility
type ArticlesServer interface {
	ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error)
	mustEmbedUnimplementedArticlesServer()
}

// UnimplementedArticlesServer must be embedded to have forward compatible implementations.
type UnimplementedArticlesServer struct {
}

func (UnimplementedArticlesServer) ListArticles(context.Context, *ListArticlesRequest) (*ListArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArticles not implemented")
}
func (UnimplementedArticlesServer) mustEmbedUnimplementedArticlesServer() {}

// UnsafeArticlesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArticlesServer will
// result in compilation errors.
type UnsafeArticlesServer interface {
	mustEmbedUnimplementedArticlesServer()
}

func RegisterArticlesServer(s grpc.ServiceRegistrar, srv ArticlesServer) {
	s.RegisterService(&Articles_ServiceDesc, srv)
}

func _Articles_ListArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArticlesServer).ListArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Articles/ListArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArticlesServer).ListArticles(ctx, req.(*ListArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Articles_ServiceDesc is the grpc.ServiceDesc for Articles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Articles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Articles",
	HandlerType: (*ArticlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListArticles",
			Handler:    _Articles_ListArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "articles/articles.proto",
}
