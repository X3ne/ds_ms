// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: guilds_service/guilds/v1/guilds.proto

package guildsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/X3ne/ds_ms/api/guilds_service/guilds/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// GuildsServiceName is the fully-qualified name of the GuildsService service.
	GuildsServiceName = "guilds.v1.GuildsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GuildsServiceCreateProcedure is the fully-qualified name of the GuildsService's Create RPC.
	GuildsServiceCreateProcedure = "/guilds.v1.GuildsService/Create"
	// GuildsServiceGetByIdProcedure is the fully-qualified name of the GuildsService's GetById RPC.
	GuildsServiceGetByIdProcedure = "/guilds.v1.GuildsService/GetById"
	// GuildsServiceUpdateProcedure is the fully-qualified name of the GuildsService's Update RPC.
	GuildsServiceUpdateProcedure = "/guilds.v1.GuildsService/Update"
	// GuildsServiceDeleteProcedure is the fully-qualified name of the GuildsService's Delete RPC.
	GuildsServiceDeleteProcedure = "/guilds.v1.GuildsService/Delete"
)

// GuildsServiceClient is a client for the guilds.v1.GuildsService service.
type GuildsServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewGuildsServiceClient constructs a client for the guilds.v1.GuildsService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGuildsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GuildsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &guildsServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+GuildsServiceCreateProcedure,
			opts...,
		),
		getById: connect.NewClient[v1.GetByIdRequest, v1.GetByIdResponse](
			httpClient,
			baseURL+GuildsServiceGetByIdProcedure,
			opts...,
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+GuildsServiceUpdateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+GuildsServiceDeleteProcedure,
			opts...,
		),
	}
}

// guildsServiceClient implements GuildsServiceClient.
type guildsServiceClient struct {
	create  *connect.Client[v1.CreateRequest, v1.CreateResponse]
	getById *connect.Client[v1.GetByIdRequest, v1.GetByIdResponse]
	update  *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete  *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Create calls guilds.v1.GuildsService.Create.
func (c *guildsServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// GetById calls guilds.v1.GuildsService.GetById.
func (c *guildsServiceClient) GetById(ctx context.Context, req *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return c.getById.CallUnary(ctx, req)
}

// Update calls guilds.v1.GuildsService.Update.
func (c *guildsServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls guilds.v1.GuildsService.Delete.
func (c *guildsServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// GuildsServiceHandler is an implementation of the guilds.v1.GuildsService service.
type GuildsServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewGuildsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGuildsServiceHandler(svc GuildsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	guildsServiceCreateHandler := connect.NewUnaryHandler(
		GuildsServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	guildsServiceGetByIdHandler := connect.NewUnaryHandler(
		GuildsServiceGetByIdProcedure,
		svc.GetById,
		opts...,
	)
	guildsServiceUpdateHandler := connect.NewUnaryHandler(
		GuildsServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	guildsServiceDeleteHandler := connect.NewUnaryHandler(
		GuildsServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/guilds.v1.GuildsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GuildsServiceCreateProcedure:
			guildsServiceCreateHandler.ServeHTTP(w, r)
		case GuildsServiceGetByIdProcedure:
			guildsServiceGetByIdHandler.ServeHTTP(w, r)
		case GuildsServiceUpdateProcedure:
			guildsServiceUpdateHandler.ServeHTTP(w, r)
		case GuildsServiceDeleteProcedure:
			guildsServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGuildsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGuildsServiceHandler struct{}

func (UnimplementedGuildsServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("guilds.v1.GuildsService.Create is not implemented"))
}

func (UnimplementedGuildsServiceHandler) GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("guilds.v1.GuildsService.GetById is not implemented"))
}

func (UnimplementedGuildsServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("guilds.v1.GuildsService.Update is not implemented"))
}

func (UnimplementedGuildsServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("guilds.v1.GuildsService.Delete is not implemented"))
}
