/**
 * @fileoverview gRPC-Web generated client stub for dashboard
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as dashboard$service_pb from './dashboard-service_pb';


export class DashboardServiceClient {
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

  methodDescriptorSubscribe = new grpcWeb.MethodDescriptor(
    '/dashboard.DashboardService/Subscribe',
    grpcWeb.MethodType.SERVER_STREAMING,
    dashboard$service_pb.SubscribeRequest,
    dashboard$service_pb.SubscribeResponse,
    (request: dashboard$service_pb.SubscribeRequest) => {
      return request.serializeBinary();
    },
    dashboard$service_pb.SubscribeResponse.deserializeBinary
  );

  subscribe(
    request: dashboard$service_pb.SubscribeRequest,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<dashboard$service_pb.SubscribeResponse> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/dashboard.DashboardService/Subscribe',
      request,
      metadata || {},
      this.methodDescriptorSubscribe);
  }

}

