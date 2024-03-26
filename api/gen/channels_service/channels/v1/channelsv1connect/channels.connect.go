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
	// ChannelsServiceGetGuildChannelsProcedure is the fully-qualified name of the ChannelsService's
	// GetGuildChannels RPC.
	ChannelsServiceGetGuildChannelsProcedure = "/channels.v1.ChannelsService/GetGuildChannels"
	// ChannelsServiceGetChannelMessagesProcedure is the fully-qualified name of the ChannelsService's
	// GetChannelMessages RPC.
	ChannelsServiceGetChannelMessagesProcedure = "/channels.v1.ChannelsService/GetChannelMessages"
	// ChannelsServiceGetChannelMessageProcedure is the fully-qualified name of the ChannelsService's
	// GetChannelMessage RPC.
	ChannelsServiceGetChannelMessageProcedure = "/channels.v1.ChannelsService/GetChannelMessage"
	// ChannelsServiceCreateMessageProcedure is the fully-qualified name of the ChannelsService's
	// CreateMessage RPC.
	ChannelsServiceCreateMessageProcedure = "/channels.v1.ChannelsService/CreateMessage"
	// ChannelsServiceUpdateMessageProcedure is the fully-qualified name of the ChannelsService's
	// UpdateMessage RPC.
	ChannelsServiceUpdateMessageProcedure = "/channels.v1.ChannelsService/UpdateMessage"
	// ChannelsServiceDeleteMessageProcedure is the fully-qualified name of the ChannelsService's
	// DeleteMessage RPC.
	ChannelsServiceDeleteMessageProcedure = "/channels.v1.ChannelsService/DeleteMessage"
	// ChannelsServiceBulkDeleteMessagesProcedure is the fully-qualified name of the ChannelsService's
	// BulkDeleteMessages RPC.
	ChannelsServiceBulkDeleteMessagesProcedure = "/channels.v1.ChannelsService/BulkDeleteMessages"
	// ChannelsServiceEditChannelPermissionsProcedure is the fully-qualified name of the
	// ChannelsService's EditChannelPermissions RPC.
	ChannelsServiceEditChannelPermissionsProcedure = "/channels.v1.ChannelsService/EditChannelPermissions"
	// ChannelsServiceDeleteChannelPermissionProcedure is the fully-qualified name of the
	// ChannelsService's DeleteChannelPermission RPC.
	ChannelsServiceDeleteChannelPermissionProcedure = "/channels.v1.ChannelsService/DeleteChannelPermission"
	// ChannelsServiceTriggerTypingIndicatorProcedure is the fully-qualified name of the
	// ChannelsService's TriggerTypingIndicator RPC.
	ChannelsServiceTriggerTypingIndicatorProcedure = "/channels.v1.ChannelsService/TriggerTypingIndicator"
	// ChannelsServiceGetPinnedMessagesProcedure is the fully-qualified name of the ChannelsService's
	// GetPinnedMessages RPC.
	ChannelsServiceGetPinnedMessagesProcedure = "/channels.v1.ChannelsService/GetPinnedMessages"
	// ChannelsServiceAddPinnedMessageProcedure is the fully-qualified name of the ChannelsService's
	// AddPinnedMessage RPC.
	ChannelsServiceAddPinnedMessageProcedure = "/channels.v1.ChannelsService/AddPinnedMessage"
	// ChannelsServiceDeletePinnedMessageProcedure is the fully-qualified name of the ChannelsService's
	// DeletePinnedMessage RPC.
	ChannelsServiceDeletePinnedMessageProcedure = "/channels.v1.ChannelsService/DeletePinnedMessage"
	// ChannelsServiceGroupDMAddRecipientProcedure is the fully-qualified name of the ChannelsService's
	// GroupDMAddRecipient RPC.
	ChannelsServiceGroupDMAddRecipientProcedure = "/channels.v1.ChannelsService/GroupDMAddRecipient"
	// ChannelsServiceGroupDMRemoveRecipientProcedure is the fully-qualified name of the
	// ChannelsService's GroupDMRemoveRecipient RPC.
	ChannelsServiceGroupDMRemoveRecipientProcedure = "/channels.v1.ChannelsService/GroupDMRemoveRecipient"
)

// ChannelsServiceClient is a client for the channels.v1.ChannelsService service.
type ChannelsServiceClient interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
	GetGuildChannels(context.Context, *connect.Request[v1.GetGuildChannelsRequest]) (*connect.Response[v1.GetGuildChannelsResponse], error)
	GetChannelMessages(context.Context, *connect.Request[v1.GetChannelMessagesRequest]) (*connect.Response[v1.GetChannelMessagesResponse], error)
	GetChannelMessage(context.Context, *connect.Request[v1.GetChannelMessageRequest]) (*connect.Response[v1.GetChannelMessageResponse], error)
	CreateMessage(context.Context, *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error)
	UpdateMessage(context.Context, *connect.Request[v1.UpdateMessageRequest]) (*connect.Response[v1.UpdateMessageResponse], error)
	DeleteMessage(context.Context, *connect.Request[v1.DeleteMessageRequest]) (*connect.Response[v1.DeleteMessageResponse], error)
	BulkDeleteMessages(context.Context, *connect.Request[v1.BulkDeleteMessagesRequest]) (*connect.Response[v1.BulkDeleteMessagesResponse], error)
	EditChannelPermissions(context.Context, *connect.Request[v1.EditChannelPermissionsRequest]) (*connect.Response[v1.EditChannelPermissionsResponse], error)
	DeleteChannelPermission(context.Context, *connect.Request[v1.DeleteChannelPermissionRequest]) (*connect.Response[v1.DeleteChannelPermissionResponse], error)
	TriggerTypingIndicator(context.Context, *connect.Request[v1.TriggerTypingIndicatorRequest]) (*connect.Response[v1.TriggerTypingIndicatorResponse], error)
	GetPinnedMessages(context.Context, *connect.Request[v1.GetPinnedMessagesRequest]) (*connect.Response[v1.GetPinnedMessagesResponse], error)
	AddPinnedMessage(context.Context, *connect.Request[v1.AddPinnedMessageRequest]) (*connect.Response[v1.AddPinnedMessageResponse], error)
	DeletePinnedMessage(context.Context, *connect.Request[v1.DeletePinnedMessageRequest]) (*connect.Response[v1.DeletePinnedMessageResponse], error)
	GroupDMAddRecipient(context.Context, *connect.Request[v1.GroupDMAddRecipientRequest]) (*connect.Response[v1.GroupDMAddRecipientResponse], error)
	GroupDMRemoveRecipient(context.Context, *connect.Request[v1.GroupDMRemoveRecipientRequest]) (*connect.Response[v1.GroupDMRemoveRecipientResponse], error)
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
		getGuildChannels: connect.NewClient[v1.GetGuildChannelsRequest, v1.GetGuildChannelsResponse](
			httpClient,
			baseURL+ChannelsServiceGetGuildChannelsProcedure,
			opts...,
		),
		getChannelMessages: connect.NewClient[v1.GetChannelMessagesRequest, v1.GetChannelMessagesResponse](
			httpClient,
			baseURL+ChannelsServiceGetChannelMessagesProcedure,
			opts...,
		),
		getChannelMessage: connect.NewClient[v1.GetChannelMessageRequest, v1.GetChannelMessageResponse](
			httpClient,
			baseURL+ChannelsServiceGetChannelMessageProcedure,
			opts...,
		),
		createMessage: connect.NewClient[v1.CreateMessageRequest, v1.CreateMessageResponse](
			httpClient,
			baseURL+ChannelsServiceCreateMessageProcedure,
			opts...,
		),
		updateMessage: connect.NewClient[v1.UpdateMessageRequest, v1.UpdateMessageResponse](
			httpClient,
			baseURL+ChannelsServiceUpdateMessageProcedure,
			opts...,
		),
		deleteMessage: connect.NewClient[v1.DeleteMessageRequest, v1.DeleteMessageResponse](
			httpClient,
			baseURL+ChannelsServiceDeleteMessageProcedure,
			opts...,
		),
		bulkDeleteMessages: connect.NewClient[v1.BulkDeleteMessagesRequest, v1.BulkDeleteMessagesResponse](
			httpClient,
			baseURL+ChannelsServiceBulkDeleteMessagesProcedure,
			opts...,
		),
		editChannelPermissions: connect.NewClient[v1.EditChannelPermissionsRequest, v1.EditChannelPermissionsResponse](
			httpClient,
			baseURL+ChannelsServiceEditChannelPermissionsProcedure,
			opts...,
		),
		deleteChannelPermission: connect.NewClient[v1.DeleteChannelPermissionRequest, v1.DeleteChannelPermissionResponse](
			httpClient,
			baseURL+ChannelsServiceDeleteChannelPermissionProcedure,
			opts...,
		),
		triggerTypingIndicator: connect.NewClient[v1.TriggerTypingIndicatorRequest, v1.TriggerTypingIndicatorResponse](
			httpClient,
			baseURL+ChannelsServiceTriggerTypingIndicatorProcedure,
			opts...,
		),
		getPinnedMessages: connect.NewClient[v1.GetPinnedMessagesRequest, v1.GetPinnedMessagesResponse](
			httpClient,
			baseURL+ChannelsServiceGetPinnedMessagesProcedure,
			opts...,
		),
		addPinnedMessage: connect.NewClient[v1.AddPinnedMessageRequest, v1.AddPinnedMessageResponse](
			httpClient,
			baseURL+ChannelsServiceAddPinnedMessageProcedure,
			opts...,
		),
		deletePinnedMessage: connect.NewClient[v1.DeletePinnedMessageRequest, v1.DeletePinnedMessageResponse](
			httpClient,
			baseURL+ChannelsServiceDeletePinnedMessageProcedure,
			opts...,
		),
		groupDMAddRecipient: connect.NewClient[v1.GroupDMAddRecipientRequest, v1.GroupDMAddRecipientResponse](
			httpClient,
			baseURL+ChannelsServiceGroupDMAddRecipientProcedure,
			opts...,
		),
		groupDMRemoveRecipient: connect.NewClient[v1.GroupDMRemoveRecipientRequest, v1.GroupDMRemoveRecipientResponse](
			httpClient,
			baseURL+ChannelsServiceGroupDMRemoveRecipientProcedure,
			opts...,
		),
	}
}

// channelsServiceClient implements ChannelsServiceClient.
type channelsServiceClient struct {
	create                  *connect.Client[v1.CreateRequest, v1.CreateResponse]
	getById                 *connect.Client[v1.GetByIdRequest, v1.GetByIdResponse]
	update                  *connect.Client[v1.UpdateRequest, v1.UpdateResponse]
	delete                  *connect.Client[v1.DeleteRequest, v1.DeleteResponse]
	getGuildChannels        *connect.Client[v1.GetGuildChannelsRequest, v1.GetGuildChannelsResponse]
	getChannelMessages      *connect.Client[v1.GetChannelMessagesRequest, v1.GetChannelMessagesResponse]
	getChannelMessage       *connect.Client[v1.GetChannelMessageRequest, v1.GetChannelMessageResponse]
	createMessage           *connect.Client[v1.CreateMessageRequest, v1.CreateMessageResponse]
	updateMessage           *connect.Client[v1.UpdateMessageRequest, v1.UpdateMessageResponse]
	deleteMessage           *connect.Client[v1.DeleteMessageRequest, v1.DeleteMessageResponse]
	bulkDeleteMessages      *connect.Client[v1.BulkDeleteMessagesRequest, v1.BulkDeleteMessagesResponse]
	editChannelPermissions  *connect.Client[v1.EditChannelPermissionsRequest, v1.EditChannelPermissionsResponse]
	deleteChannelPermission *connect.Client[v1.DeleteChannelPermissionRequest, v1.DeleteChannelPermissionResponse]
	triggerTypingIndicator  *connect.Client[v1.TriggerTypingIndicatorRequest, v1.TriggerTypingIndicatorResponse]
	getPinnedMessages       *connect.Client[v1.GetPinnedMessagesRequest, v1.GetPinnedMessagesResponse]
	addPinnedMessage        *connect.Client[v1.AddPinnedMessageRequest, v1.AddPinnedMessageResponse]
	deletePinnedMessage     *connect.Client[v1.DeletePinnedMessageRequest, v1.DeletePinnedMessageResponse]
	groupDMAddRecipient     *connect.Client[v1.GroupDMAddRecipientRequest, v1.GroupDMAddRecipientResponse]
	groupDMRemoveRecipient  *connect.Client[v1.GroupDMRemoveRecipientRequest, v1.GroupDMRemoveRecipientResponse]
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

// GetGuildChannels calls channels.v1.ChannelsService.GetGuildChannels.
func (c *channelsServiceClient) GetGuildChannels(ctx context.Context, req *connect.Request[v1.GetGuildChannelsRequest]) (*connect.Response[v1.GetGuildChannelsResponse], error) {
	return c.getGuildChannels.CallUnary(ctx, req)
}

// GetChannelMessages calls channels.v1.ChannelsService.GetChannelMessages.
func (c *channelsServiceClient) GetChannelMessages(ctx context.Context, req *connect.Request[v1.GetChannelMessagesRequest]) (*connect.Response[v1.GetChannelMessagesResponse], error) {
	return c.getChannelMessages.CallUnary(ctx, req)
}

// GetChannelMessage calls channels.v1.ChannelsService.GetChannelMessage.
func (c *channelsServiceClient) GetChannelMessage(ctx context.Context, req *connect.Request[v1.GetChannelMessageRequest]) (*connect.Response[v1.GetChannelMessageResponse], error) {
	return c.getChannelMessage.CallUnary(ctx, req)
}

// CreateMessage calls channels.v1.ChannelsService.CreateMessage.
func (c *channelsServiceClient) CreateMessage(ctx context.Context, req *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error) {
	return c.createMessage.CallUnary(ctx, req)
}

// UpdateMessage calls channels.v1.ChannelsService.UpdateMessage.
func (c *channelsServiceClient) UpdateMessage(ctx context.Context, req *connect.Request[v1.UpdateMessageRequest]) (*connect.Response[v1.UpdateMessageResponse], error) {
	return c.updateMessage.CallUnary(ctx, req)
}

// DeleteMessage calls channels.v1.ChannelsService.DeleteMessage.
func (c *channelsServiceClient) DeleteMessage(ctx context.Context, req *connect.Request[v1.DeleteMessageRequest]) (*connect.Response[v1.DeleteMessageResponse], error) {
	return c.deleteMessage.CallUnary(ctx, req)
}

// BulkDeleteMessages calls channels.v1.ChannelsService.BulkDeleteMessages.
func (c *channelsServiceClient) BulkDeleteMessages(ctx context.Context, req *connect.Request[v1.BulkDeleteMessagesRequest]) (*connect.Response[v1.BulkDeleteMessagesResponse], error) {
	return c.bulkDeleteMessages.CallUnary(ctx, req)
}

// EditChannelPermissions calls channels.v1.ChannelsService.EditChannelPermissions.
func (c *channelsServiceClient) EditChannelPermissions(ctx context.Context, req *connect.Request[v1.EditChannelPermissionsRequest]) (*connect.Response[v1.EditChannelPermissionsResponse], error) {
	return c.editChannelPermissions.CallUnary(ctx, req)
}

// DeleteChannelPermission calls channels.v1.ChannelsService.DeleteChannelPermission.
func (c *channelsServiceClient) DeleteChannelPermission(ctx context.Context, req *connect.Request[v1.DeleteChannelPermissionRequest]) (*connect.Response[v1.DeleteChannelPermissionResponse], error) {
	return c.deleteChannelPermission.CallUnary(ctx, req)
}

// TriggerTypingIndicator calls channels.v1.ChannelsService.TriggerTypingIndicator.
func (c *channelsServiceClient) TriggerTypingIndicator(ctx context.Context, req *connect.Request[v1.TriggerTypingIndicatorRequest]) (*connect.Response[v1.TriggerTypingIndicatorResponse], error) {
	return c.triggerTypingIndicator.CallUnary(ctx, req)
}

// GetPinnedMessages calls channels.v1.ChannelsService.GetPinnedMessages.
func (c *channelsServiceClient) GetPinnedMessages(ctx context.Context, req *connect.Request[v1.GetPinnedMessagesRequest]) (*connect.Response[v1.GetPinnedMessagesResponse], error) {
	return c.getPinnedMessages.CallUnary(ctx, req)
}

// AddPinnedMessage calls channels.v1.ChannelsService.AddPinnedMessage.
func (c *channelsServiceClient) AddPinnedMessage(ctx context.Context, req *connect.Request[v1.AddPinnedMessageRequest]) (*connect.Response[v1.AddPinnedMessageResponse], error) {
	return c.addPinnedMessage.CallUnary(ctx, req)
}

// DeletePinnedMessage calls channels.v1.ChannelsService.DeletePinnedMessage.
func (c *channelsServiceClient) DeletePinnedMessage(ctx context.Context, req *connect.Request[v1.DeletePinnedMessageRequest]) (*connect.Response[v1.DeletePinnedMessageResponse], error) {
	return c.deletePinnedMessage.CallUnary(ctx, req)
}

// GroupDMAddRecipient calls channels.v1.ChannelsService.GroupDMAddRecipient.
func (c *channelsServiceClient) GroupDMAddRecipient(ctx context.Context, req *connect.Request[v1.GroupDMAddRecipientRequest]) (*connect.Response[v1.GroupDMAddRecipientResponse], error) {
	return c.groupDMAddRecipient.CallUnary(ctx, req)
}

// GroupDMRemoveRecipient calls channels.v1.ChannelsService.GroupDMRemoveRecipient.
func (c *channelsServiceClient) GroupDMRemoveRecipient(ctx context.Context, req *connect.Request[v1.GroupDMRemoveRecipientRequest]) (*connect.Response[v1.GroupDMRemoveRecipientResponse], error) {
	return c.groupDMRemoveRecipient.CallUnary(ctx, req)
}

// ChannelsServiceHandler is an implementation of the channels.v1.ChannelsService service.
type ChannelsServiceHandler interface {
	Create(context.Context, *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error)
	GetById(context.Context, *connect.Request[v1.GetByIdRequest]) (*connect.Response[v1.GetByIdResponse], error)
	Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error)
	Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error)
	GetGuildChannels(context.Context, *connect.Request[v1.GetGuildChannelsRequest]) (*connect.Response[v1.GetGuildChannelsResponse], error)
	GetChannelMessages(context.Context, *connect.Request[v1.GetChannelMessagesRequest]) (*connect.Response[v1.GetChannelMessagesResponse], error)
	GetChannelMessage(context.Context, *connect.Request[v1.GetChannelMessageRequest]) (*connect.Response[v1.GetChannelMessageResponse], error)
	CreateMessage(context.Context, *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error)
	UpdateMessage(context.Context, *connect.Request[v1.UpdateMessageRequest]) (*connect.Response[v1.UpdateMessageResponse], error)
	DeleteMessage(context.Context, *connect.Request[v1.DeleteMessageRequest]) (*connect.Response[v1.DeleteMessageResponse], error)
	BulkDeleteMessages(context.Context, *connect.Request[v1.BulkDeleteMessagesRequest]) (*connect.Response[v1.BulkDeleteMessagesResponse], error)
	EditChannelPermissions(context.Context, *connect.Request[v1.EditChannelPermissionsRequest]) (*connect.Response[v1.EditChannelPermissionsResponse], error)
	DeleteChannelPermission(context.Context, *connect.Request[v1.DeleteChannelPermissionRequest]) (*connect.Response[v1.DeleteChannelPermissionResponse], error)
	TriggerTypingIndicator(context.Context, *connect.Request[v1.TriggerTypingIndicatorRequest]) (*connect.Response[v1.TriggerTypingIndicatorResponse], error)
	GetPinnedMessages(context.Context, *connect.Request[v1.GetPinnedMessagesRequest]) (*connect.Response[v1.GetPinnedMessagesResponse], error)
	AddPinnedMessage(context.Context, *connect.Request[v1.AddPinnedMessageRequest]) (*connect.Response[v1.AddPinnedMessageResponse], error)
	DeletePinnedMessage(context.Context, *connect.Request[v1.DeletePinnedMessageRequest]) (*connect.Response[v1.DeletePinnedMessageResponse], error)
	GroupDMAddRecipient(context.Context, *connect.Request[v1.GroupDMAddRecipientRequest]) (*connect.Response[v1.GroupDMAddRecipientResponse], error)
	GroupDMRemoveRecipient(context.Context, *connect.Request[v1.GroupDMRemoveRecipientRequest]) (*connect.Response[v1.GroupDMRemoveRecipientResponse], error)
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
	channelsServiceGetGuildChannelsHandler := connect.NewUnaryHandler(
		ChannelsServiceGetGuildChannelsProcedure,
		svc.GetGuildChannels,
		opts...,
	)
	channelsServiceGetChannelMessagesHandler := connect.NewUnaryHandler(
		ChannelsServiceGetChannelMessagesProcedure,
		svc.GetChannelMessages,
		opts...,
	)
	channelsServiceGetChannelMessageHandler := connect.NewUnaryHandler(
		ChannelsServiceGetChannelMessageProcedure,
		svc.GetChannelMessage,
		opts...,
	)
	channelsServiceCreateMessageHandler := connect.NewUnaryHandler(
		ChannelsServiceCreateMessageProcedure,
		svc.CreateMessage,
		opts...,
	)
	channelsServiceUpdateMessageHandler := connect.NewUnaryHandler(
		ChannelsServiceUpdateMessageProcedure,
		svc.UpdateMessage,
		opts...,
	)
	channelsServiceDeleteMessageHandler := connect.NewUnaryHandler(
		ChannelsServiceDeleteMessageProcedure,
		svc.DeleteMessage,
		opts...,
	)
	channelsServiceBulkDeleteMessagesHandler := connect.NewUnaryHandler(
		ChannelsServiceBulkDeleteMessagesProcedure,
		svc.BulkDeleteMessages,
		opts...,
	)
	channelsServiceEditChannelPermissionsHandler := connect.NewUnaryHandler(
		ChannelsServiceEditChannelPermissionsProcedure,
		svc.EditChannelPermissions,
		opts...,
	)
	channelsServiceDeleteChannelPermissionHandler := connect.NewUnaryHandler(
		ChannelsServiceDeleteChannelPermissionProcedure,
		svc.DeleteChannelPermission,
		opts...,
	)
	channelsServiceTriggerTypingIndicatorHandler := connect.NewUnaryHandler(
		ChannelsServiceTriggerTypingIndicatorProcedure,
		svc.TriggerTypingIndicator,
		opts...,
	)
	channelsServiceGetPinnedMessagesHandler := connect.NewUnaryHandler(
		ChannelsServiceGetPinnedMessagesProcedure,
		svc.GetPinnedMessages,
		opts...,
	)
	channelsServiceAddPinnedMessageHandler := connect.NewUnaryHandler(
		ChannelsServiceAddPinnedMessageProcedure,
		svc.AddPinnedMessage,
		opts...,
	)
	channelsServiceDeletePinnedMessageHandler := connect.NewUnaryHandler(
		ChannelsServiceDeletePinnedMessageProcedure,
		svc.DeletePinnedMessage,
		opts...,
	)
	channelsServiceGroupDMAddRecipientHandler := connect.NewUnaryHandler(
		ChannelsServiceGroupDMAddRecipientProcedure,
		svc.GroupDMAddRecipient,
		opts...,
	)
	channelsServiceGroupDMRemoveRecipientHandler := connect.NewUnaryHandler(
		ChannelsServiceGroupDMRemoveRecipientProcedure,
		svc.GroupDMRemoveRecipient,
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
		case ChannelsServiceGetGuildChannelsProcedure:
			channelsServiceGetGuildChannelsHandler.ServeHTTP(w, r)
		case ChannelsServiceGetChannelMessagesProcedure:
			channelsServiceGetChannelMessagesHandler.ServeHTTP(w, r)
		case ChannelsServiceGetChannelMessageProcedure:
			channelsServiceGetChannelMessageHandler.ServeHTTP(w, r)
		case ChannelsServiceCreateMessageProcedure:
			channelsServiceCreateMessageHandler.ServeHTTP(w, r)
		case ChannelsServiceUpdateMessageProcedure:
			channelsServiceUpdateMessageHandler.ServeHTTP(w, r)
		case ChannelsServiceDeleteMessageProcedure:
			channelsServiceDeleteMessageHandler.ServeHTTP(w, r)
		case ChannelsServiceBulkDeleteMessagesProcedure:
			channelsServiceBulkDeleteMessagesHandler.ServeHTTP(w, r)
		case ChannelsServiceEditChannelPermissionsProcedure:
			channelsServiceEditChannelPermissionsHandler.ServeHTTP(w, r)
		case ChannelsServiceDeleteChannelPermissionProcedure:
			channelsServiceDeleteChannelPermissionHandler.ServeHTTP(w, r)
		case ChannelsServiceTriggerTypingIndicatorProcedure:
			channelsServiceTriggerTypingIndicatorHandler.ServeHTTP(w, r)
		case ChannelsServiceGetPinnedMessagesProcedure:
			channelsServiceGetPinnedMessagesHandler.ServeHTTP(w, r)
		case ChannelsServiceAddPinnedMessageProcedure:
			channelsServiceAddPinnedMessageHandler.ServeHTTP(w, r)
		case ChannelsServiceDeletePinnedMessageProcedure:
			channelsServiceDeletePinnedMessageHandler.ServeHTTP(w, r)
		case ChannelsServiceGroupDMAddRecipientProcedure:
			channelsServiceGroupDMAddRecipientHandler.ServeHTTP(w, r)
		case ChannelsServiceGroupDMRemoveRecipientProcedure:
			channelsServiceGroupDMRemoveRecipientHandler.ServeHTTP(w, r)
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

func (UnimplementedChannelsServiceHandler) GetGuildChannels(context.Context, *connect.Request[v1.GetGuildChannelsRequest]) (*connect.Response[v1.GetGuildChannelsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GetGuildChannels is not implemented"))
}

func (UnimplementedChannelsServiceHandler) GetChannelMessages(context.Context, *connect.Request[v1.GetChannelMessagesRequest]) (*connect.Response[v1.GetChannelMessagesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GetChannelMessages is not implemented"))
}

func (UnimplementedChannelsServiceHandler) GetChannelMessage(context.Context, *connect.Request[v1.GetChannelMessageRequest]) (*connect.Response[v1.GetChannelMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GetChannelMessage is not implemented"))
}

func (UnimplementedChannelsServiceHandler) CreateMessage(context.Context, *connect.Request[v1.CreateMessageRequest]) (*connect.Response[v1.CreateMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.CreateMessage is not implemented"))
}

func (UnimplementedChannelsServiceHandler) UpdateMessage(context.Context, *connect.Request[v1.UpdateMessageRequest]) (*connect.Response[v1.UpdateMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.UpdateMessage is not implemented"))
}

func (UnimplementedChannelsServiceHandler) DeleteMessage(context.Context, *connect.Request[v1.DeleteMessageRequest]) (*connect.Response[v1.DeleteMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.DeleteMessage is not implemented"))
}

func (UnimplementedChannelsServiceHandler) BulkDeleteMessages(context.Context, *connect.Request[v1.BulkDeleteMessagesRequest]) (*connect.Response[v1.BulkDeleteMessagesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.BulkDeleteMessages is not implemented"))
}

func (UnimplementedChannelsServiceHandler) EditChannelPermissions(context.Context, *connect.Request[v1.EditChannelPermissionsRequest]) (*connect.Response[v1.EditChannelPermissionsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.EditChannelPermissions is not implemented"))
}

func (UnimplementedChannelsServiceHandler) DeleteChannelPermission(context.Context, *connect.Request[v1.DeleteChannelPermissionRequest]) (*connect.Response[v1.DeleteChannelPermissionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.DeleteChannelPermission is not implemented"))
}

func (UnimplementedChannelsServiceHandler) TriggerTypingIndicator(context.Context, *connect.Request[v1.TriggerTypingIndicatorRequest]) (*connect.Response[v1.TriggerTypingIndicatorResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.TriggerTypingIndicator is not implemented"))
}

func (UnimplementedChannelsServiceHandler) GetPinnedMessages(context.Context, *connect.Request[v1.GetPinnedMessagesRequest]) (*connect.Response[v1.GetPinnedMessagesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GetPinnedMessages is not implemented"))
}

func (UnimplementedChannelsServiceHandler) AddPinnedMessage(context.Context, *connect.Request[v1.AddPinnedMessageRequest]) (*connect.Response[v1.AddPinnedMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.AddPinnedMessage is not implemented"))
}

func (UnimplementedChannelsServiceHandler) DeletePinnedMessage(context.Context, *connect.Request[v1.DeletePinnedMessageRequest]) (*connect.Response[v1.DeletePinnedMessageResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.DeletePinnedMessage is not implemented"))
}

func (UnimplementedChannelsServiceHandler) GroupDMAddRecipient(context.Context, *connect.Request[v1.GroupDMAddRecipientRequest]) (*connect.Response[v1.GroupDMAddRecipientResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GroupDMAddRecipient is not implemented"))
}

func (UnimplementedChannelsServiceHandler) GroupDMRemoveRecipient(context.Context, *connect.Request[v1.GroupDMRemoveRecipientRequest]) (*connect.Response[v1.GroupDMRemoveRecipientResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("channels.v1.ChannelsService.GroupDMRemoveRecipient is not implemented"))
}
