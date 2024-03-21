// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: channels_service/channels/v1/channels.proto

package channelsv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/X3ne/ds_ms/api/gen/channels_service/channels/v1"
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
	// ChannelsServiceName is the fully-qualified name of the ChannelsService service.
	ChannelsServiceName = "channels.v1.ChannelsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ChannelsServiceCreateProcedure is the fully-qualified name of the ChannelsService's Create RPC.
	ChannelsServiceCreateProcedure = "/channels.v1.ChannelsService/Create"
	// ChannelsServiceGetByIdProcedure is the fully-qualified name of the ChannelsService's GetById RPC.
	ChannelsServiceGetByIdProcedure = "/channels.v1.ChannelsService/GetById"
	// ChannelsServiceUpdateProcedure is the fully-qualified name of the ChannelsService's Update RPC.
	ChannelsServiceUpdateProcedure = "/channels.v1.ChannelsService/Update"
	// ChannelsServiceDeleteProcedure is the fully-qualified name of the ChannelsService's Delete RPC.
	ChannelsServiceDeleteProcedure = "/channels.v1.ChannelsService/Delete"
)

// ChannelsServiceClient is a client for the channels.v1.ChannelsService service.
type ChannelsServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewChannelsServiceClient constructs a client for the channels.v1.ChannelsService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewChannelsServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ChannelsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &channelsServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+ChannelsServiceCreateProcedure,
			opts...,
		),
		getById: connect.NewClient[v1.GetByIdRequest, v1.GetByIdResponse](
			httpClient,
			baseURL+ChannelsServiceGetByIdProcedure,
			opts...,
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+ChannelsServiceUpdateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+ChannelsServiceDeleteProcedure,
			opts...,
		),
	}
}

// channelsServiceClient implements ChannelsServiceClient.
type channelsServiceClient struct {
	create  *connect.Client[v1.CreateRequest, v1.CreateResponse]
	getById *connect.Client[v1.GetByIdRequest, v1.GetByIdResponse]
	update  *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete  *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Create calls channels.v1.ChannelsService.Create.
func (c *channelsServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// GetById calls channels.v1.ChannelsService.GetById.
func (c *channelsServiceClient) GetById(ctx context.Context, req *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return c.getById.CallUnary(ctx, req)
}

// Update calls channels.v1.ChannelsService.Update.
func (c *channelsServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls channels.v1.ChannelsService.Delete.
func (c *channelsServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// ChannelsServiceHandler is an implementation of the channels.v1.ChannelsService service.
type ChannelsServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewChannelsServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewChannelsServiceHandler(svc ChannelsServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	channelsServiceCreateHandler := connect.NewUnaryHandler(
		ChannelsServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	channelsServiceGetByIdHandler := connect.NewUnaryHandler(
		ChannelsServiceGetByIdProcedure,
		svc.GetById,
		opts...,
	)
	channelsServiceUpdateHandler := connect.NewUnaryHandler(
		ChannelsServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	channelsServiceDeleteHandler := connect.NewUnaryHandler(
		ChannelsServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/channels.v1.ChannelsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ChannelsServiceCreateProcedure:
			channelsServiceCreateHandler.ServeHTTP(w, r)
		case ChannelsServiceGetByIdProcedure:
			channelsServiceGetByIdHandler.ServeHTTP(w, r)
		case ChannelsServiceUpdateProcedure:
			channelsServiceUpdateHandler.ServeHTTP(w, r)
		case ChannelsServiceDeleteProcedure:
			channelsServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedChannelsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedChannelsServiceHandler struct{}

func (UnimplementedChannelsServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.Create is not implemented"))
}

func (UnimplementedChannelsServiceHandler) GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GetById is not implemented"))
}

func (UnimplementedChannelsServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.Update is not implemented"))
}

func (UnimplementedChannelsServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.Delete is not implemented"))
}