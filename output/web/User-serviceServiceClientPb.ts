/**
 * @fileoverview gRPC-Web generated client stub for user
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as user$service_pb from './user-service_pb';


export class UserServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorAddUser = new grpcWeb.MethodDescriptor(
    '/user.UserService/AddUser',
    grpcWeb.MethodType.UNARY,
    user$service_pb.AddUserRequest,
    user$service_pb.AddUserResponse,
    (request: user$service_pb.AddUserRequest) => {
      return request.serializeBinary();
    },
    user$service_pb.AddUserResponse.deserializeBinary
  );

  addUser(
    request: user$service_pb.AddUserRequest,
    metadata: grpcWeb.Metadata | null): Promise<user$service_pb.AddUserResponse>;

  addUser(
    request: user$service_pb.AddUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: user$service_pb.AddUserResponse) => void): grpcWeb.ClientReadableStream<user$service_pb.AddUserResponse>;

  addUser(
    request: user$service_pb.AddUserRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: user$service_pb.AddUserResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/user.UserService/AddUser',
        request,
        metadata || {},
        this.methodDescriptorAddUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/user.UserService/AddUser',
    request,
    metadata || {},
    this.methodDescriptorAddUser);
  }

}

