# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: relayerrpc.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='relayerrpc.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x10relayerrpc.proto\x12\x02pb\x1a\x1bgoogle/protobuf/empty.proto\x1a\x0c\x63ommon.proto2O\n\nRelayerRPC\x12\x41\n\x11SubmitTransaction\x12\x14.pb.ShardTransaction\x1a\x16.google.protobuf.Emptyb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,common__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)



_RELAYERRPC = _descriptor.ServiceDescriptor(
  name='RelayerRPC',
  full_name='pb.RelayerRPC',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=67,
  serialized_end=146,
  methods=[
  _descriptor.MethodDescriptor(
    name='SubmitTransaction',
    full_name='pb.RelayerRPC.SubmitTransaction',
    index=0,
    containing_service=None,
    input_type=common__pb2._SHARDTRANSACTION,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_RELAYERRPC)

DESCRIPTOR.services_by_name['RelayerRPC'] = _RELAYERRPC

# @@protoc_insertion_point(module_scope)
