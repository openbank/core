// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: bnk.to/core/api/v1/groups/all.proto

package groups

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GroupsServiceClient is the client API for GroupsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupsServiceClient interface {
	// CreateGroup creates a new group.
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// ListGroups lists groups.
	ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error)
	// GetGroup retrieves a group.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// UpdateGroup updates the specified group.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error)
	// DeleteGroup deletes the specified group.
	DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// AddMember adds the specified groups to a client.
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListMembers lists the members in a group.
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
	// GetMember retrieves the specified member in a group.
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*Member, error)
	// RemoveMember removes the specified groups from the client.
	RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// CreateRole creates a new group role.
	CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// ListGroupRoles retrieves all of the roles of the specified group.
	ListGroupRoles(ctx context.Context, in *ListGroupRolesRequest, opts ...grpc.CallOption) (*ListGroupRolesResponse, error)
	// GetRole retrieves group role.
	GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// UpdateRole updates the specified group role.
	UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*Role, error)
	// DeleteRole deletes the specified group role.
	DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// AddMemberRole adds a role to a member.
	AddMemberRole(ctx context.Context, in *AddMemberRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// ListMemberRoles lists the roles for a member.
	ListMemberRoles(ctx context.Context, in *ListMemberRolesRequest, opts ...grpc.CallOption) (*ListMemberRolesResponse, error)
	// RemoveMemberRole removes a role from a member.
	RemoveMemberRole(ctx context.Context, in *RemoveMemberRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type groupsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupsServiceClient(cc grpc.ClientConnInterface) GroupsServiceClient {
	return &groupsServiceClient{cc}
}

func (c *groupsServiceClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) ListGroups(ctx context.Context, in *ListGroupsRequest, opts ...grpc.CallOption) (*ListGroupsResponse, error) {
	out := new(ListGroupsResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) DeleteGroup(ctx context.Context, in *DeleteGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/DeleteGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/AddMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*Member, error) {
	out := new(Member)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/RemoveMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) CreateRole(ctx context.Context, in *CreateRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/CreateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) ListGroupRoles(ctx context.Context, in *ListGroupRolesRequest, opts ...grpc.CallOption) (*ListGroupRolesResponse, error) {
	out := new(ListGroupRolesResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/ListGroupRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) GetRole(ctx context.Context, in *GetRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/GetRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) UpdateRole(ctx context.Context, in *UpdateRoleRequest, opts ...grpc.CallOption) (*Role, error) {
	out := new(Role)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/UpdateRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) DeleteRole(ctx context.Context, in *DeleteRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/DeleteRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) AddMemberRole(ctx context.Context, in *AddMemberRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/AddMemberRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) ListMemberRoles(ctx context.Context, in *ListMemberRolesRequest, opts ...grpc.CallOption) (*ListMemberRolesResponse, error) {
	out := new(ListMemberRolesResponse)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/ListMemberRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsServiceClient) RemoveMemberRole(ctx context.Context, in *RemoveMemberRoleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/openbank.core.v1.groups.GroupsService/RemoveMemberRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsServiceServer is the server API for GroupsService service.
// All implementations must embed UnimplementedGroupsServiceServer
// for forward compatibility
type GroupsServiceServer interface {
	// CreateGroup creates a new group.
	CreateGroup(context.Context, *CreateGroupRequest) (*Group, error)
	// ListGroups lists groups.
	ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error)
	// GetGroup retrieves a group.
	GetGroup(context.Context, *GetGroupRequest) (*Group, error)
	// UpdateGroup updates the specified group.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error)
	// DeleteGroup deletes the specified group.
	DeleteGroup(context.Context, *DeleteGroupRequest) (*emptypb.Empty, error)
	// AddMember adds the specified groups to a client.
	AddMember(context.Context, *AddMemberRequest) (*emptypb.Empty, error)
	// ListMembers lists the members in a group.
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	// GetMember retrieves the specified member in a group.
	GetMember(context.Context, *GetMemberRequest) (*Member, error)
	// RemoveMember removes the specified groups from the client.
	RemoveMember(context.Context, *RemoveMemberRequest) (*emptypb.Empty, error)
	// CreateRole creates a new group role.
	CreateRole(context.Context, *CreateRoleRequest) (*Role, error)
	// ListGroupRoles retrieves all of the roles of the specified group.
	ListGroupRoles(context.Context, *ListGroupRolesRequest) (*ListGroupRolesResponse, error)
	// GetRole retrieves group role.
	GetRole(context.Context, *GetRoleRequest) (*Role, error)
	// UpdateRole updates the specified group role.
	UpdateRole(context.Context, *UpdateRoleRequest) (*Role, error)
	// DeleteRole deletes the specified group role.
	DeleteRole(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error)
	// AddMemberRole adds a role to a member.
	AddMemberRole(context.Context, *AddMemberRoleRequest) (*emptypb.Empty, error)
	// ListMemberRoles lists the roles for a member.
	ListMemberRoles(context.Context, *ListMemberRolesRequest) (*ListMemberRolesResponse, error)
	// RemoveMemberRole removes a role from a member.
	RemoveMemberRole(context.Context, *RemoveMemberRoleRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedGroupsServiceServer()
}

// UnimplementedGroupsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGroupsServiceServer struct{}

func (UnimplementedGroupsServiceServer) CreateGroup(context.Context, *CreateGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}

func (UnimplementedGroupsServiceServer) ListGroups(context.Context, *ListGroupsRequest) (*ListGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}

func (UnimplementedGroupsServiceServer) GetGroup(context.Context, *GetGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}

func (UnimplementedGroupsServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}

func (UnimplementedGroupsServiceServer) DeleteGroup(context.Context, *DeleteGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}

func (UnimplementedGroupsServiceServer) AddMember(context.Context, *AddMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}

func (UnimplementedGroupsServiceServer) ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}

func (UnimplementedGroupsServiceServer) GetMember(context.Context, *GetMemberRequest) (*Member, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}

func (UnimplementedGroupsServiceServer) RemoveMember(context.Context, *RemoveMemberRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMember not implemented")
}

func (UnimplementedGroupsServiceServer) CreateRole(context.Context, *CreateRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRole not implemented")
}

func (UnimplementedGroupsServiceServer) ListGroupRoles(context.Context, *ListGroupRolesRequest) (*ListGroupRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroupRoles not implemented")
}

func (UnimplementedGroupsServiceServer) GetRole(context.Context, *GetRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRole not implemented")
}

func (UnimplementedGroupsServiceServer) UpdateRole(context.Context, *UpdateRoleRequest) (*Role, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRole not implemented")
}

func (UnimplementedGroupsServiceServer) DeleteRole(context.Context, *DeleteRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRole not implemented")
}

func (UnimplementedGroupsServiceServer) AddMemberRole(context.Context, *AddMemberRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMemberRole not implemented")
}

func (UnimplementedGroupsServiceServer) ListMemberRoles(context.Context, *ListMemberRolesRequest) (*ListMemberRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMemberRoles not implemented")
}

func (UnimplementedGroupsServiceServer) RemoveMemberRole(context.Context, *RemoveMemberRoleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMemberRole not implemented")
}
func (UnimplementedGroupsServiceServer) mustEmbedUnimplementedGroupsServiceServer() {}

// UnsafeGroupsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupsServiceServer will
// result in compilation errors.
type UnsafeGroupsServiceServer interface {
	mustEmbedUnimplementedGroupsServiceServer()
}

func RegisterGroupsServiceServer(s grpc.ServiceRegistrar, srv GroupsServiceServer) {
	s.RegisterService(&GroupsService_ServiceDesc, srv)
}

func _GroupsService_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).ListGroups(ctx, req.(*ListGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/DeleteGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).DeleteGroup(ctx, req.(*DeleteGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/AddMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_RemoveMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).RemoveMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/RemoveMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).RemoveMember(ctx, req.(*RemoveMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_CreateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).CreateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/CreateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).CreateRole(ctx, req.(*CreateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_ListGroupRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).ListGroupRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/ListGroupRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).ListGroupRoles(ctx, req.(*ListGroupRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_GetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).GetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/GetRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).GetRole(ctx, req.(*GetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_UpdateRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).UpdateRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/UpdateRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).UpdateRole(ctx, req.(*UpdateRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_DeleteRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).DeleteRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/DeleteRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).DeleteRole(ctx, req.(*DeleteRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_AddMemberRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).AddMemberRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/AddMemberRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).AddMemberRole(ctx, req.(*AddMemberRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_ListMemberRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMemberRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).ListMemberRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/ListMemberRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).ListMemberRoles(ctx, req.(*ListMemberRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupsService_RemoveMemberRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServiceServer).RemoveMemberRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openbank.core.v1.groups.GroupsService/RemoveMemberRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServiceServer).RemoveMemberRole(ctx, req.(*RemoveMemberRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupsService_ServiceDesc is the grpc.ServiceDesc for GroupsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openbank.core.v1.groups.GroupsService",
	HandlerType: (*GroupsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGroup",
			Handler:    _GroupsService_CreateGroup_Handler,
		},
		{
			MethodName: "ListGroups",
			Handler:    _GroupsService_ListGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _GroupsService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _GroupsService_UpdateGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _GroupsService_DeleteGroup_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _GroupsService_AddMember_Handler,
		},
		{
			MethodName: "ListMembers",
			Handler:    _GroupsService_ListMembers_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _GroupsService_GetMember_Handler,
		},
		{
			MethodName: "RemoveMember",
			Handler:    _GroupsService_RemoveMember_Handler,
		},
		{
			MethodName: "CreateRole",
			Handler:    _GroupsService_CreateRole_Handler,
		},
		{
			MethodName: "ListGroupRoles",
			Handler:    _GroupsService_ListGroupRoles_Handler,
		},
		{
			MethodName: "GetRole",
			Handler:    _GroupsService_GetRole_Handler,
		},
		{
			MethodName: "UpdateRole",
			Handler:    _GroupsService_UpdateRole_Handler,
		},
		{
			MethodName: "DeleteRole",
			Handler:    _GroupsService_DeleteRole_Handler,
		},
		{
			MethodName: "AddMemberRole",
			Handler:    _GroupsService_AddMemberRole_Handler,
		},
		{
			MethodName: "ListMemberRoles",
			Handler:    _GroupsService_ListMemberRoles_Handler,
		},
		{
			MethodName: "RemoveMemberRole",
			Handler:    _GroupsService_RemoveMemberRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bnk.to/core/api/v1/groups/all.proto",
}
