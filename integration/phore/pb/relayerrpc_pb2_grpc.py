# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import common_pb2 as common__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class RelayerRPCStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.SubmitTransaction = channel.unary_unary(
        '/pb.RelayerRPC/SubmitTransaction',
        request_serializer=common__pb2.ShardTransaction.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )
    self.GetListeningAddresses = channel.unary_unary(
        '/pb.RelayerRPC/GetListeningAddresses',
        request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
        response_deserializer=common__pb2.ListeningAddressesResponse.FromString,
        )
    self.Connect = channel.unary_unary(
        '/pb.RelayerRPC/Connect',
        request_serializer=common__pb2.ConnectMessage.SerializeToString,
        response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
        )


class RelayerRPCServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def SubmitTransaction(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GetListeningAddresses(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def Connect(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_RelayerRPCServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'SubmitTransaction': grpc.unary_unary_rpc_method_handler(
          servicer.SubmitTransaction,
          request_deserializer=common__pb2.ShardTransaction.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
      'GetListeningAddresses': grpc.unary_unary_rpc_method_handler(
          servicer.GetListeningAddresses,
          request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
          response_serializer=common__pb2.ListeningAddressesResponse.SerializeToString,
      ),
      'Connect': grpc.unary_unary_rpc_method_handler(
          servicer.Connect,
          request_deserializer=common__pb2.ConnectMessage.FromString,
          response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'pb.RelayerRPC', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
