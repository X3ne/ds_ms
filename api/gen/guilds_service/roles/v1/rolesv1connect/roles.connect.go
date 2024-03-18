// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: guilds_service/roles/v1/roles.proto

package rolesv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/X3ne/ds_ms/api/gen/guilds_service/roles/v1"
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
	// RolesServiceName is the fully-qualified name of the RolesService service.
	RolesServiceName = "roles.v1.RolesService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RolesServiceCreateProcedure is the fully-qualified name of the RolesService's Create RPC.
	RolesServiceCreateProcedure = "/roles.v1.RolesService/Create"
	// RolesServiceGetByIdProcedure is the fully-qualified name of the RolesService's GetById RPC.
	RolesServiceGetByIdProcedure = "/roles.v1.RolesService/GetById"
	// RolesServiceUpdateProcedure is the fully-qualified name of the RolesService's Update RPC.
	RolesServiceUpdateProcedure = "/roles.v1.RolesService/Update"
	// RolesServiceDeleteProcedure is the fully-qualified name of the RolesService's Delete RPC.
	RolesServiceDeleteProcedure = "/roles.v1.RolesService/Delete"
)

// RolesServiceClient is a client for the roles.v1.RolesService service.
type RolesServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewRolesServiceClient constructs a client for the roles.v1.RolesService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRolesServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RolesServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rolesServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+RolesServiceCreateProcedure,
			opts...,
		),
		getById: connect.NewClient[v1.GetByIdRequest, v1.GetByIdResponse](
			httpClient,
			baseURL+RolesServiceGetByIdProcedure,
			opts...,
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+RolesServiceUpdateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+RolesServiceDeleteProcedure,
			opts...,
		),
	}
}

// rolesServiceClient implements RolesServiceClient.
type rolesServiceClient struct {
	create  *connect.Client[v1.CreateRequest, v1.CreateResponse]
	getById *connect.Client[v1.GetByIdRequest, v1.GetByIdResponse]
	update  *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete  *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Create calls roles.v1.RolesService.Create.
func (c *rolesServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// GetById calls roles.v1.RolesService.GetById.
func (c *rolesServiceClient) GetById(ctx context.Context, req *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return c.getById.CallUnary(ctx, req)
}

// Update calls roles.v1.RolesService.Update.
func (c *rolesServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls roles.v1.RolesService.Delete.
func (c *rolesServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// RolesServiceHandler is an implementation of the roles.v1.RolesService service.
type RolesServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewRolesServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRolesServiceHandler(svc RolesServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	rolesServiceCreateHandler := connect.NewUnaryHandler(
		RolesServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	rolesServiceGetByIdHandler := connect.NewUnaryHandler(
		RolesServiceGetByIdProcedure,
		svc.GetById,
		opts...,
	)
	rolesServiceUpdateHandler := connect.NewUnaryHandler(
		RolesServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	rolesServiceDeleteHandler := connect.NewUnaryHandler(
		RolesServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/roles.v1.RolesService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RolesServiceCreateProcedure:
			rolesServiceCreateHandler.ServeHTTP(w, r)
		case RolesServiceGetByIdProcedure:
			rolesServiceGetByIdHandler.ServeHTTP(w, r)
		case RolesServiceUpdateProcedure:
			rolesServiceUpdateHandler.ServeHTTP(w, r)
		case RolesServiceDeleteProcedure:
			rolesServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRolesServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedRolesServiceHandler struct{}

func (UnimplementedRolesServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("roles.v1.RolesService.Create is not implemented"))
}

func (UnimplementedRolesServiceHandler) GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("roles.v1.RolesService.GetById is not implemented"))
}

func (UnimplementedRolesServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("roles.v1.RolesService.Update is not implemented"))
}

func (UnimplementedRolesServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("roles.v1.RolesService.Delete is not implemented"))
}