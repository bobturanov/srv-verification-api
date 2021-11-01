# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/srv_verification_api/v1/srv_verification_api.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ozonmp/srv_verification_api/v1/srv_verification_api.proto',
  package='ozonmp.srv_verification_api.v1',
  syntax='proto3',
  serialized_options=b'ZTgithub.com/ozonmp/srv-verification-api/pkg/srv-verification-api;srv_verification_api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n9ozonmp/srv_verification_api/v1/srv_verification_api.proto\x12\x1eozonmp.srv_verification_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"2\n\x0cVerification\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\"Q\n\x1d\x44\x65scribeVerificationV1Request\x12\x30\n\x0fverification_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0everificationId\"d\n\x1e\x44\x65scribeVerificationV1Response\x12\x42\n\x05value\x18\x01 \x01(\x0b\x32,.ozonmp.srv_verification_api.v1.VerificationR\x05value\"S\n\x1b\x43reateVerificationV1Request\x12\x34\n\x11verification_name\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x02R\x10verificationName\"P\n\x1c\x43reateVerificationV1Response\x12\x30\n\x0fverification_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0everificationId\"\x1b\n\x19ListVerificationV1Request\"n\n\x1aListVerificationV1Response\x12P\n\x0cverification\x18\x01 \x03(\x0b\x32,.ozonmp.srv_verification_api.v1.VerificationR\x0cverification\"O\n\x1bRemoveVerificationV1Request\x12\x30\n\x0fverification_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\x0everificationId\"6\n\x1cRemoveVerificationV1Response\x12\x16\n\x06result\x18\x01 \x01(\x08R\x06result2\x91\x06\n\x19SrvVerificationApiService\x12\xc4\x01\n\x16\x44\x65scribeVerificationV1\x12=.ozonmp.srv_verification_api.v1.DescribeVerificationV1Request\x1a>.ozonmp.srv_verification_api.v1.DescribeVerificationV1Response\"+\x82\xd3\xe4\x93\x02%\x12#/v1/verifications/{verification_id}\x12\xb6\x01\n\x14\x43reateVerificationV1\x12;.ozonmp.srv_verification_api.v1.CreateVerificationV1Request\x1a<.ozonmp.srv_verification_api.v1.CreateVerificationV1Response\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/v1/verifications/create:\x01*\x12\xab\x01\n\x12ListVerificationV1\x12\x39.ozonmp.srv_verification_api.v1.ListVerificationV1Request\x1a:.ozonmp.srv_verification_api.v1.ListVerificationV1Response\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/verifications/list\x12\xc5\x01\n\x14RemoveVerificationV1\x12;.ozonmp.srv_verification_api.v1.RemoveVerificationV1Request\x1a<.ozonmp.srv_verification_api.v1.RemoveVerificationV1Response\"2\x82\xd3\xe4\x93\x02,**/v1/verifications/remove/{verification_id}BVZTgithub.com/ozonmp/srv-verification-api/pkg/srv-verification-api;srv_verification_apib\x06proto3'
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_VERIFICATION = _descriptor.Descriptor(
  name='Verification',
  full_name='ozonmp.srv_verification_api.v1.Verification',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='ozonmp.srv_verification_api.v1.Verification.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='ozonmp.srv_verification_api.v1.Verification.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='name', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=181,
  serialized_end=231,
)


_DESCRIBEVERIFICATIONV1REQUEST = _descriptor.Descriptor(
  name='DescribeVerificationV1Request',
  full_name='ozonmp.srv_verification_api.v1.DescribeVerificationV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='verification_id', full_name='ozonmp.srv_verification_api.v1.DescribeVerificationV1Request.verification_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='verificationId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=233,
  serialized_end=314,
)


_DESCRIBEVERIFICATIONV1RESPONSE = _descriptor.Descriptor(
  name='DescribeVerificationV1Response',
  full_name='ozonmp.srv_verification_api.v1.DescribeVerificationV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='ozonmp.srv_verification_api.v1.DescribeVerificationV1Response.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=316,
  serialized_end=416,
)


_CREATEVERIFICATIONV1REQUEST = _descriptor.Descriptor(
  name='CreateVerificationV1Request',
  full_name='ozonmp.srv_verification_api.v1.CreateVerificationV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='verification_name', full_name='ozonmp.srv_verification_api.v1.CreateVerificationV1Request.verification_name', index=0,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\004r\002\020\002', json_name='verificationName', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=418,
  serialized_end=501,
)


_CREATEVERIFICATIONV1RESPONSE = _descriptor.Descriptor(
  name='CreateVerificationV1Response',
  full_name='ozonmp.srv_verification_api.v1.CreateVerificationV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='verification_id', full_name='ozonmp.srv_verification_api.v1.CreateVerificationV1Response.verification_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='verificationId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=503,
  serialized_end=583,
)


_LISTVERIFICATIONV1REQUEST = _descriptor.Descriptor(
  name='ListVerificationV1Request',
  full_name='ozonmp.srv_verification_api.v1.ListVerificationV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=585,
  serialized_end=612,
)


_LISTVERIFICATIONV1RESPONSE = _descriptor.Descriptor(
  name='ListVerificationV1Response',
  full_name='ozonmp.srv_verification_api.v1.ListVerificationV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='verification', full_name='ozonmp.srv_verification_api.v1.ListVerificationV1Response.verification', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='verification', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=614,
  serialized_end=724,
)


_REMOVEVERIFICATIONV1REQUEST = _descriptor.Descriptor(
  name='RemoveVerificationV1Request',
  full_name='ozonmp.srv_verification_api.v1.RemoveVerificationV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='verification_id', full_name='ozonmp.srv_verification_api.v1.RemoveVerificationV1Request.verification_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='verificationId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=726,
  serialized_end=805,
)


_REMOVEVERIFICATIONV1RESPONSE = _descriptor.Descriptor(
  name='RemoveVerificationV1Response',
  full_name='ozonmp.srv_verification_api.v1.RemoveVerificationV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='ozonmp.srv_verification_api.v1.RemoveVerificationV1Response.result', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='result', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=807,
  serialized_end=861,
)

_DESCRIBEVERIFICATIONV1RESPONSE.fields_by_name['value'].message_type = _VERIFICATION
_LISTVERIFICATIONV1RESPONSE.fields_by_name['verification'].message_type = _VERIFICATION
DESCRIPTOR.message_types_by_name['Verification'] = _VERIFICATION
DESCRIPTOR.message_types_by_name['DescribeVerificationV1Request'] = _DESCRIBEVERIFICATIONV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeVerificationV1Response'] = _DESCRIBEVERIFICATIONV1RESPONSE
DESCRIPTOR.message_types_by_name['CreateVerificationV1Request'] = _CREATEVERIFICATIONV1REQUEST
DESCRIPTOR.message_types_by_name['CreateVerificationV1Response'] = _CREATEVERIFICATIONV1RESPONSE
DESCRIPTOR.message_types_by_name['ListVerificationV1Request'] = _LISTVERIFICATIONV1REQUEST
DESCRIPTOR.message_types_by_name['ListVerificationV1Response'] = _LISTVERIFICATIONV1RESPONSE
DESCRIPTOR.message_types_by_name['RemoveVerificationV1Request'] = _REMOVEVERIFICATIONV1REQUEST
DESCRIPTOR.message_types_by_name['RemoveVerificationV1Response'] = _REMOVEVERIFICATIONV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Verification = _reflection.GeneratedProtocolMessageType('Verification', (_message.Message,), {
  'DESCRIPTOR' : _VERIFICATION,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.Verification)
  })
_sym_db.RegisterMessage(Verification)

DescribeVerificationV1Request = _reflection.GeneratedProtocolMessageType('DescribeVerificationV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEVERIFICATIONV1REQUEST,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.DescribeVerificationV1Request)
  })
_sym_db.RegisterMessage(DescribeVerificationV1Request)

DescribeVerificationV1Response = _reflection.GeneratedProtocolMessageType('DescribeVerificationV1Response', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEVERIFICATIONV1RESPONSE,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.DescribeVerificationV1Response)
  })
_sym_db.RegisterMessage(DescribeVerificationV1Response)

CreateVerificationV1Request = _reflection.GeneratedProtocolMessageType('CreateVerificationV1Request', (_message.Message,), {
  'DESCRIPTOR' : _CREATEVERIFICATIONV1REQUEST,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.CreateVerificationV1Request)
  })
_sym_db.RegisterMessage(CreateVerificationV1Request)

CreateVerificationV1Response = _reflection.GeneratedProtocolMessageType('CreateVerificationV1Response', (_message.Message,), {
  'DESCRIPTOR' : _CREATEVERIFICATIONV1RESPONSE,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.CreateVerificationV1Response)
  })
_sym_db.RegisterMessage(CreateVerificationV1Response)

ListVerificationV1Request = _reflection.GeneratedProtocolMessageType('ListVerificationV1Request', (_message.Message,), {
  'DESCRIPTOR' : _LISTVERIFICATIONV1REQUEST,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.ListVerificationV1Request)
  })
_sym_db.RegisterMessage(ListVerificationV1Request)

ListVerificationV1Response = _reflection.GeneratedProtocolMessageType('ListVerificationV1Response', (_message.Message,), {
  'DESCRIPTOR' : _LISTVERIFICATIONV1RESPONSE,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.ListVerificationV1Response)
  })
_sym_db.RegisterMessage(ListVerificationV1Response)

RemoveVerificationV1Request = _reflection.GeneratedProtocolMessageType('RemoveVerificationV1Request', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEVERIFICATIONV1REQUEST,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.RemoveVerificationV1Request)
  })
_sym_db.RegisterMessage(RemoveVerificationV1Request)

RemoveVerificationV1Response = _reflection.GeneratedProtocolMessageType('RemoveVerificationV1Response', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEVERIFICATIONV1RESPONSE,
  '__module__' : 'ozonmp.srv_verification_api.v1.srv_verification_api_pb2'
  # @@protoc_insertion_point(class_scope:ozonmp.srv_verification_api.v1.RemoveVerificationV1Response)
  })
_sym_db.RegisterMessage(RemoveVerificationV1Response)


DESCRIPTOR._options = None
_DESCRIBEVERIFICATIONV1REQUEST.fields_by_name['verification_id']._options = None
_CREATEVERIFICATIONV1REQUEST.fields_by_name['verification_name']._options = None
_CREATEVERIFICATIONV1RESPONSE.fields_by_name['verification_id']._options = None
_REMOVEVERIFICATIONV1REQUEST.fields_by_name['verification_id']._options = None

_SRVVERIFICATIONAPISERVICE = _descriptor.ServiceDescriptor(
  name='SrvVerificationApiService',
  full_name='ozonmp.srv_verification_api.v1.SrvVerificationApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=864,
  serialized_end=1649,
  methods=[
  _descriptor.MethodDescriptor(
    name='DescribeVerificationV1',
    full_name='ozonmp.srv_verification_api.v1.SrvVerificationApiService.DescribeVerificationV1',
    index=0,
    containing_service=None,
    input_type=_DESCRIBEVERIFICATIONV1REQUEST,
    output_type=_DESCRIBEVERIFICATIONV1RESPONSE,
    serialized_options=b'\202\323\344\223\002%\022#/v1/verifications/{verification_id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateVerificationV1',
    full_name='ozonmp.srv_verification_api.v1.SrvVerificationApiService.CreateVerificationV1',
    index=1,
    containing_service=None,
    input_type=_CREATEVERIFICATIONV1REQUEST,
    output_type=_CREATEVERIFICATIONV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\035\"\030/v1/verifications/create:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ListVerificationV1',
    full_name='ozonmp.srv_verification_api.v1.SrvVerificationApiService.ListVerificationV1',
    index=2,
    containing_service=None,
    input_type=_LISTVERIFICATIONV1REQUEST,
    output_type=_LISTVERIFICATIONV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\030\022\026/v1/verifications/list',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveVerificationV1',
    full_name='ozonmp.srv_verification_api.v1.SrvVerificationApiService.RemoveVerificationV1',
    index=3,
    containing_service=None,
    input_type=_REMOVEVERIFICATIONV1REQUEST,
    output_type=_REMOVEVERIFICATIONV1RESPONSE,
    serialized_options=b'\202\323\344\223\002,**/v1/verifications/remove/{verification_id}',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_SRVVERIFICATIONAPISERVICE)

DESCRIPTOR.services_by_name['SrvVerificationApiService'] = _SRVVERIFICATIONAPISERVICE

# @@protoc_insertion_point(module_scope)