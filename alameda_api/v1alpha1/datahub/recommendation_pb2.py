# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/recommendation.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from alameda_api.v1alpha1.datahub import metadata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2
from alameda_api.v1alpha1.datahub import pod_assignment_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_pod__assignment__pb2
from alameda_api.v1alpha1.datahub import metric_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2
from alameda_api.v1alpha1.datahub import resource_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_resource__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/recommendation.proto',
  package='containers_ai.alameda.v1alpha1.datahub',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n1alameda_api/v1alpha1/datahub/recommendation.proto\x12&containers_ai.alameda.v1alpha1.datahub\x1a\x1fgoogle/protobuf/timestamp.proto\x1a+alameda_api/v1alpha1/datahub/metadata.proto\x1a\x31\x61lameda_api/v1alpha1/datahub/pod_assignment.proto\x1a)alameda_api/v1alpha1/datahub/metric.proto\x1a+alameda_api/v1alpha1/datahub/resource.proto\"\x87\x03\n\x17\x43ontainerRecommendation\x12\x0c\n\x04name\x18\x01 \x01(\t\x12Q\n\x15limit_recommendations\x18\x02 \x03(\x0b\x32\x32.containers_ai.alameda.v1alpha1.datahub.MetricData\x12S\n\x17request_recommendations\x18\x03 \x03(\x0b\x32\x32.containers_ai.alameda.v1alpha1.datahub.MetricData\x12Y\n\x1dinitial_limit_recommendations\x18\x04 \x03(\x0b\x32\x32.containers_ai.alameda.v1alpha1.datahub.MetricData\x12[\n\x1finitial_request_recommendations\x18\x05 \x03(\x0b\x32\x32.containers_ai.alameda.v1alpha1.datahub.MetricData\"\xf4\x01\n\x0f\x41ssignPodPolicy\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12M\n\rnode_priority\x18\x02 \x01(\x0b\x32\x34.containers_ai.alameda.v1alpha1.datahub.NodePriorityH\x00\x12I\n\rnode_selector\x18\x03 \x01(\x0b\x32\x30.containers_ai.alameda.v1alpha1.datahub.SelectorH\x00\x12\x13\n\tnode_name\x18\x04 \x01(\tH\x00\x42\x08\n\x06policy\"\x86\x04\n\x11PodRecommendation\x12O\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32\x36.containers_ai.alameda.v1alpha1.datahub.NamespacedName\x12 \n\x18\x61pply_recommendation_now\x18\x02 \x01(\x08\x12R\n\x11\x61ssign_pod_policy\x18\x03 \x01(\x0b\x32\x37.containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy\x12\x62\n\x19\x63ontainer_recommendations\x18\x04 \x03(\x0b\x32?.containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation\x12.\n\nstart_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12M\n\x0etop_controller\x18\x07 \x01(\x0b\x32\x35.containers_ai.alameda.v1alpha1.datahub.TopController\x12\x19\n\x11recommendation_id\x18\x08 \x01(\tb\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_pod__assignment__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_resource__pb2.DESCRIPTOR,])




_CONTAINERRECOMMENDATION = _descriptor.Descriptor(
  name='ContainerRecommendation',
  full_name='containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit_recommendations', full_name='containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation.limit_recommendations', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='request_recommendations', full_name='containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation.request_recommendations', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='initial_limit_recommendations', full_name='containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation.initial_limit_recommendations', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='initial_request_recommendations', full_name='containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation.initial_request_recommendations', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=311,
  serialized_end=702,
)


_ASSIGNPODPOLICY = _descriptor.Descriptor(
  name='AssignPodPolicy',
  full_name='containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='node_priority', full_name='containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy.node_priority', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='node_selector', full_name='containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy.node_selector', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='node_name', full_name='containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy.node_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
    _descriptor.OneofDescriptor(
      name='policy', full_name='containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy.policy',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=705,
  serialized_end=949,
)


_PODRECOMMENDATION = _descriptor.Descriptor(
  name='PodRecommendation',
  full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='apply_recommendation_now', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.apply_recommendation_now', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='assign_pod_policy', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.assign_pod_policy', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container_recommendations', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.container_recommendations', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.start_time', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.end_time', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='top_controller', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.top_controller', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='recommendation_id', full_name='containers_ai.alameda.v1alpha1.datahub.PodRecommendation.recommendation_id', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=952,
  serialized_end=1470,
)

_CONTAINERRECOMMENDATION.fields_by_name['limit_recommendations'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2._METRICDATA
_CONTAINERRECOMMENDATION.fields_by_name['request_recommendations'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2._METRICDATA
_CONTAINERRECOMMENDATION.fields_by_name['initial_limit_recommendations'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2._METRICDATA
_CONTAINERRECOMMENDATION.fields_by_name['initial_request_recommendations'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2._METRICDATA
_ASSIGNPODPOLICY.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_ASSIGNPODPOLICY.fields_by_name['node_priority'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_pod__assignment__pb2._NODEPRIORITY
_ASSIGNPODPOLICY.fields_by_name['node_selector'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_pod__assignment__pb2._SELECTOR
_ASSIGNPODPOLICY.oneofs_by_name['policy'].fields.append(
  _ASSIGNPODPOLICY.fields_by_name['node_priority'])
_ASSIGNPODPOLICY.fields_by_name['node_priority'].containing_oneof = _ASSIGNPODPOLICY.oneofs_by_name['policy']
_ASSIGNPODPOLICY.oneofs_by_name['policy'].fields.append(
  _ASSIGNPODPOLICY.fields_by_name['node_selector'])
_ASSIGNPODPOLICY.fields_by_name['node_selector'].containing_oneof = _ASSIGNPODPOLICY.oneofs_by_name['policy']
_ASSIGNPODPOLICY.oneofs_by_name['policy'].fields.append(
  _ASSIGNPODPOLICY.fields_by_name['node_name'])
_ASSIGNPODPOLICY.fields_by_name['node_name'].containing_oneof = _ASSIGNPODPOLICY.oneofs_by_name['policy']
_PODRECOMMENDATION.fields_by_name['namespaced_name'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2._NAMESPACEDNAME
_PODRECOMMENDATION.fields_by_name['assign_pod_policy'].message_type = _ASSIGNPODPOLICY
_PODRECOMMENDATION.fields_by_name['container_recommendations'].message_type = _CONTAINERRECOMMENDATION
_PODRECOMMENDATION.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_PODRECOMMENDATION.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_PODRECOMMENDATION.fields_by_name['top_controller'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_resource__pb2._TOPCONTROLLER
DESCRIPTOR.message_types_by_name['ContainerRecommendation'] = _CONTAINERRECOMMENDATION
DESCRIPTOR.message_types_by_name['AssignPodPolicy'] = _ASSIGNPODPOLICY
DESCRIPTOR.message_types_by_name['PodRecommendation'] = _PODRECOMMENDATION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerRecommendation = _reflection.GeneratedProtocolMessageType('ContainerRecommendation', (_message.Message,), dict(
  DESCRIPTOR = _CONTAINERRECOMMENDATION,
  __module__ = 'alameda_api.v1alpha1.datahub.recommendation_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.ContainerRecommendation)
  ))
_sym_db.RegisterMessage(ContainerRecommendation)

AssignPodPolicy = _reflection.GeneratedProtocolMessageType('AssignPodPolicy', (_message.Message,), dict(
  DESCRIPTOR = _ASSIGNPODPOLICY,
  __module__ = 'alameda_api.v1alpha1.datahub.recommendation_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.AssignPodPolicy)
  ))
_sym_db.RegisterMessage(AssignPodPolicy)

PodRecommendation = _reflection.GeneratedProtocolMessageType('PodRecommendation', (_message.Message,), dict(
  DESCRIPTOR = _PODRECOMMENDATION,
  __module__ = 'alameda_api.v1alpha1.datahub.recommendation_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.PodRecommendation)
  ))
_sym_db.RegisterMessage(PodRecommendation)


# @@protoc_insertion_point(module_scope)
