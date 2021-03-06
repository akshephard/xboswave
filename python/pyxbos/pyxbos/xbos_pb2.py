# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xbos.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import hamilton_pb2 as hamilton__pb2
from . import iot_pb2 as iot__pb2
from . import dentmeter_pb2 as dentmeter__pb2
from . import system_monitor_pb2 as system__monitor__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='xbos.proto',
  package='xbospb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\nxbos.proto\x12\x06xbospb\x1a\x0ehamilton.proto\x1a\tiot.proto\x1a\x0f\x64\x65ntmeter.proto\x1a\x14system_monitor.proto\"\xb2\x03\n\x04XBOS\x12*\n\x0cHamiltonData\x18\x32 \x01(\x0b\x32\x14.xbospb.HamiltonData\x12\x38\n\x13HamiltonBRLinkStats\x18\x33 \x01(\x0b\x32\x1b.xbospb.HamiltonBRLinkStats\x12\x34\n\x11HamiltonBRMessage\x18\x34 \x01(\x0b\x32\x19.xbospb.HamiltonBRMessage\x12\x36\n\x12XBOSIoTDeviceState\x18\x64 \x01(\x0b\x32\x1a.xbospb.XBOSIoTDeviceState\x12>\n\x16XBOSIoTDeviceActuation\x18\x65 \x01(\x0b\x32\x1e.xbospb.XBOSIoTDeviceActuation\x12.\n\x0eXBOSIoTContext\x18\x66 \x01(\x0b\x32\x16.xbospb.XBOSIoTContext\x12/\n\x0e\x44\x65ntMeterState\x18\x96\x01 \x01(\x0b\x32\x16.xbospb.DentMeterState\x12\x35\n\x11\x42\x61sicServerStatus\x18\xc8\x01 \x01(\x0b\x32\x19.xbospb.BasicServerStatusb\x06proto3')
  ,
  dependencies=[hamilton__pb2.DESCRIPTOR,iot__pb2.DESCRIPTOR,dentmeter__pb2.DESCRIPTOR,system__monitor__pb2.DESCRIPTOR,])




_XBOS = _descriptor.Descriptor(
  name='XBOS',
  full_name='xbospb.XBOS',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='HamiltonData', full_name='xbospb.XBOS.HamiltonData', index=0,
      number=50, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='HamiltonBRLinkStats', full_name='xbospb.XBOS.HamiltonBRLinkStats', index=1,
      number=51, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='HamiltonBRMessage', full_name='xbospb.XBOS.HamiltonBRMessage', index=2,
      number=52, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='XBOSIoTDeviceState', full_name='xbospb.XBOS.XBOSIoTDeviceState', index=3,
      number=100, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='XBOSIoTDeviceActuation', full_name='xbospb.XBOS.XBOSIoTDeviceActuation', index=4,
      number=101, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='XBOSIoTContext', full_name='xbospb.XBOS.XBOSIoTContext', index=5,
      number=102, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='DentMeterState', full_name='xbospb.XBOS.DentMeterState', index=6,
      number=150, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='BasicServerStatus', full_name='xbospb.XBOS.BasicServerStatus', index=7,
      number=200, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=89,
  serialized_end=523,
)

_XBOS.fields_by_name['HamiltonData'].message_type = hamilton__pb2._HAMILTONDATA
_XBOS.fields_by_name['HamiltonBRLinkStats'].message_type = hamilton__pb2._HAMILTONBRLINKSTATS
_XBOS.fields_by_name['HamiltonBRMessage'].message_type = hamilton__pb2._HAMILTONBRMESSAGE
_XBOS.fields_by_name['XBOSIoTDeviceState'].message_type = iot__pb2._XBOSIOTDEVICESTATE
_XBOS.fields_by_name['XBOSIoTDeviceActuation'].message_type = iot__pb2._XBOSIOTDEVICEACTUATION
_XBOS.fields_by_name['XBOSIoTContext'].message_type = iot__pb2._XBOSIOTCONTEXT
_XBOS.fields_by_name['DentMeterState'].message_type = dentmeter__pb2._DENTMETERSTATE
_XBOS.fields_by_name['BasicServerStatus'].message_type = system__monitor__pb2._BASICSERVERSTATUS
DESCRIPTOR.message_types_by_name['XBOS'] = _XBOS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

XBOS = _reflection.GeneratedProtocolMessageType('XBOS', (_message.Message,), dict(
  DESCRIPTOR = _XBOS,
  __module__ = 'xbos_pb2'
  # @@protoc_insertion_point(class_scope:xbospb.XBOS)
  ))
_sym_db.RegisterMessage(XBOS)


# @@protoc_insertion_point(module_scope)
