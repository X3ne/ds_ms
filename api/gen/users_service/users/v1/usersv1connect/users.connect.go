// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: users_service/users/v1/users.proto

package usersv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/X3ne/ds_ms/api/gen/users_service/users/v1"
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
	// UsersServiceName is the fully-qualified name of the UsersService service.
	UsersServiceName = "users.v1.UsersService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UsersServiceCreateProcedure is the fully-qualified name of the UsersService's Create RPC.
	UsersServiceCreateProcedure = "/users.v1.UsersService/Create"
	// UsersServiceGetByIdProcedure is the fully-qualified name of the UsersService's GetById RPC.
	UsersServiceGetByIdProcedure = "/users.v1.UsersService/GetById"
	// UsersServiceGetByEmailProcedure is the fully-qualified name of the UsersService's GetByEmail RPC.
	UsersServiceGetByEmailProcedure = "/users.v1.UsersService/GetByEmail"
	// UsersServiceUpdateProcedure is the fully-qualified name of the UsersService's Update RPC.
	UsersServiceUpdateProcedure = "/users.v1.UsersService/Update"
	// UsersServiceDeleteProcedure is the fully-qualified name of the UsersService's Delete RPC.
	UsersServiceDeleteProcedure = "/users.v1.UsersService/Delete"
)

// UsersServiceClient is a client for the users.v1.UsersService service.
type UsersServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	GetByEmail(context.Context, *connect.Request[v1.GetByEmailRequest]) (*connect.Response[v1.GetByEmailResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewUsersServiceClient constructs a client for the users.v1.UsersService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUsersServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UsersServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &usersServiceClient{
		create: connect.NewClient[v1.CreateRequest, v1.CreateResponse](
			httpClient,
			baseURL+UsersServiceCreateProcedure,
			opts...,
		),
		getById: connect.NewClient[v1.GetByIdRequest, v1.GetByIdResponse](
			httpClient,
			baseURL+UsersServiceGetByIdProcedure,
			opts...,
		),
		getByEmail: connect.NewClient[v1.GetByEmailRequest, v1.GetByEmailResponse](
			httpClient,
			baseURL+UsersServiceGetByEmailProcedure,
			opts...,
		),
		update: connect.NewClient[v1.UpdateRequest, v1.UpdateResponse](
			httpClient,
			baseURL+UsersServiceUpdateProcedure,
			opts...,
		),
		delete: connect.NewClient[v1.DeleteRequest, v1.DeleteResponse](
			httpClient,
			baseURL+UsersServiceDeleteProcedure,
			opts...,
		),
	}
}

// usersServiceClient implements UsersServiceClient.
type usersServiceClient struct {
	create     *connect.Client[v1.CreateRequest, v1.CreateResponse]
	getById    *connect.Client[v1.GetByIdRequest, v1.GetByIdResponse]
	getByEmail *connect.Client[v1.GetByEmailRequest, v1.GetByEmailResponse]
	update     *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete     *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
}

// Create calls users.v1.UsersService.Create.
func (c *usersServiceClient) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// GetById calls users.v1.UsersService.GetById.
func (c *usersServiceClient) GetById(ctx context.Context, req *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return c.getById.CallUnary(ctx, req)
}

// GetByEmail calls users.v1.UsersService.GetByEmail.
func (c *usersServiceClient) GetByEmail(ctx context.Context, req *connect.Request[v1.GetByEmailRequest]) (*connect.Response[v1.GetByEmailResponse], error) {
	return c.getByEmail.CallUnary(ctx, req)
}

// Update calls users.v1.UsersService.Update.
func (c *usersServiceClient) Update(ctx context.Context, req *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return c.update.CallUnary(ctx, req)
}

// Delete calls users.v1.UsersService.Delete.
func (c *usersServiceClient) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// UsersServiceHandler is an implementation of the users.v1.UsersService service.
type UsersServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	GetByEmail(context.Context, *connect.Request[v1.GetByEmailRequest]) (*connect.Response[v1.GetByEmailResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
}

// NewUsersServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUsersServiceHandler(svc UsersServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	usersServiceCreateHandler := connect.NewUnaryHandler(
		UsersServiceCreateProcedure,
		svc.Create,
		opts...,
	)
	usersServiceGetByIdHandler := connect.NewUnaryHandler(
		UsersServiceGetByIdProcedure,
		svc.GetById,
		opts...,
	)
	usersServiceGetByEmailHandler := connect.NewUnaryHandler(
		UsersServiceGetByEmailProcedure,
		svc.GetByEmail,
		opts...,
	)
	usersServiceUpdateHandler := connect.NewUnaryHandler(
		UsersServiceUpdateProcedure,
		svc.Update,
		opts...,
	)
	usersServiceDeleteHandler := connect.NewUnaryHandler(
		UsersServiceDeleteProcedure,
		svc.Delete,
		opts...,
	)
	return "/users.v1.UsersService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UsersServiceCreateProcedure:
			usersServiceCreateHandler.ServeHTTP(w, r)
		case UsersServiceGetByIdProcedure:
			usersServiceGetByIdHandler.ServeHTTP(w, r)
		case UsersServiceGetByEmailProcedure:
			usersServiceGetByEmailHandler.ServeHTTP(w, r)
		case UsersServiceUpdateProcedure:
			usersServiceUpdateHandler.ServeHTTP(w, r)
		case UsersServiceDeleteProcedure:
			usersServiceDeleteHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUsersServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUsersServiceHandler struct{}

func (UnimplementedUsersServiceHandler) Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("users.v1.UsersService.Create is not implemented"))
}

func (UnimplementedUsersServiceHandler) GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("users.v1.UsersService.GetById is not implemented"))
}

func (UnimplementedUsersServiceHandler) GetByEmail(context.Context, *connect.Request[v1.GetByEmailRequest]) (*connect.Response[v1.GetByEmailResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("users.v1.UsersService.GetByEmail is not implemented"))
}

func (UnimplementedUsersServiceHandler) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("users.v1.UsersService.Update is not implemented"))
}

func (UnimplementedUsersServiceHandler) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("users.v1.UsersService.Delete is not implemented"))
}
