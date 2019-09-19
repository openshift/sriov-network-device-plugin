// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/model.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Deployment state of the model.
type Model_DeploymentState int32

const (
	// Should not be used, an un-set enum has this value by default.
	Model_DEPLOYMENT_STATE_UNSPECIFIED Model_DeploymentState = 0
	// Model is deployed.
	Model_DEPLOYED Model_DeploymentState = 1
	// Model is not deployed.
	Model_UNDEPLOYED Model_DeploymentState = 2
)

var Model_DeploymentState_name = map[int32]string{
	0: "DEPLOYMENT_STATE_UNSPECIFIED",
	1: "DEPLOYED",
	2: "UNDEPLOYED",
}

var Model_DeploymentState_value = map[string]int32{
	"DEPLOYMENT_STATE_UNSPECIFIED": 0,
	"DEPLOYED":                     1,
	"UNDEPLOYED":                   2,
}

func (x Model_DeploymentState) String() string {
	return proto.EnumName(Model_DeploymentState_name, int32(x))
}

func (Model_DeploymentState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_452845e4ed6fce9d, []int{0, 0}
}

// API proto representing a trained machine learning model.
type Model struct {
	// Required.
	// The model metadata that is specific to the problem type.
	// Must match the metadata type of the dataset used to train the model.
	//
	// Types that are valid to be assigned to ModelMetadata:
	//	*Model_TranslationModelMetadata
	ModelMetadata isModel_ModelMetadata `protobuf_oneof:"model_metadata"`
	// Output only. Resource name of the model.
	// Format: `projects/{project_id}/locations/{location_id}/models/{model_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the model to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9. It must start with a letter.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Required. The resource ID of the dataset used to create the model. The dataset must
	// come from the same ancestor project and location.
	DatasetId string `protobuf:"bytes,3,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Output only. Timestamp when the model training finished  and can be used for prediction.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this model was last updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,11,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. Deployment state of the model. A model can only serve
	// prediction requests after it gets deployed.
	DeploymentState Model_DeploymentState `protobuf:"varint,8,opt,name=deployment_state,json=deploymentState,proto3,enum=google.cloud.automl.v1.Model_DeploymentState" json:"deployment_state,omitempty"`
	// Optional. The labels with user-defined metadata to organize your model.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// Label values are optional. Label keys must start with a letter.
	//
	// See https://goo.gl/xmQnxf for more information on and examples of labels.
	Labels               map[string]string `protobuf:"bytes,34,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Model) Reset()         { *m = Model{} }
func (m *Model) String() string { return proto.CompactTextString(m) }
func (*Model) ProtoMessage()    {}
func (*Model) Descriptor() ([]byte, []int) {
	return fileDescriptor_452845e4ed6fce9d, []int{0}
}

func (m *Model) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Model.Unmarshal(m, b)
}
func (m *Model) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Model.Marshal(b, m, deterministic)
}
func (m *Model) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Model.Merge(m, src)
}
func (m *Model) XXX_Size() int {
	return xxx_messageInfo_Model.Size(m)
}
func (m *Model) XXX_DiscardUnknown() {
	xxx_messageInfo_Model.DiscardUnknown(m)
}

var xxx_messageInfo_Model proto.InternalMessageInfo

type isModel_ModelMetadata interface {
	isModel_ModelMetadata()
}

type Model_TranslationModelMetadata struct {
	TranslationModelMetadata *TranslationModelMetadata `protobuf:"bytes,15,opt,name=translation_model_metadata,json=translationModelMetadata,proto3,oneof"`
}

func (*Model_TranslationModelMetadata) isModel_ModelMetadata() {}

func (m *Model) GetModelMetadata() isModel_ModelMetadata {
	if m != nil {
		return m.ModelMetadata
	}
	return nil
}

func (m *Model) GetTranslationModelMetadata() *TranslationModelMetadata {
	if x, ok := m.GetModelMetadata().(*Model_TranslationModelMetadata); ok {
		return x.TranslationModelMetadata
	}
	return nil
}

func (m *Model) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Model) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Model) GetDatasetId() string {
	if m != nil {
		return m.DatasetId
	}
	return ""
}

func (m *Model) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Model) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Model) GetDeploymentState() Model_DeploymentState {
	if m != nil {
		return m.DeploymentState
	}
	return Model_DEPLOYMENT_STATE_UNSPECIFIED
}

func (m *Model) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Model) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Model_TranslationModelMetadata)(nil),
	}
}

func init() {
	proto.RegisterEnum("google.cloud.automl.v1.Model_DeploymentState", Model_DeploymentState_name, Model_DeploymentState_value)
	proto.RegisterType((*Model)(nil), "google.cloud.automl.v1.Model")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.automl.v1.Model.LabelsEntry")
}

func init() { proto.RegisterFile("google/cloud/automl/v1/model.proto", fileDescriptor_452845e4ed6fce9d) }

var fileDescriptor_452845e4ed6fce9d = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xd1, 0x8e, 0x93, 0x40,
	0x18, 0x85, 0x97, 0xd6, 0xae, 0xbb, 0x3f, 0x9b, 0xb6, 0x99, 0x98, 0x0d, 0x92, 0x35, 0xd6, 0x5e,
	0xe1, 0x85, 0x60, 0xeb, 0x8d, 0xb2, 0xde, 0x74, 0xb7, 0xa8, 0x4d, 0xda, 0x5a, 0x29, 0xbb, 0x51,
	0xd3, 0x84, 0x4c, 0xcb, 0x48, 0x88, 0x03, 0x43, 0x60, 0x68, 0xd2, 0x17, 0xf0, 0xde, 0xd7, 0xf2,
	0x51, 0x7c, 0x0a, 0xc3, 0x0c, 0xad, 0xbb, 0x1b, 0xeb, 0xde, 0xcd, 0xfc, 0xe7, 0x3b, 0x7f, 0x0e,
	0x07, 0x80, 0x6e, 0xc8, 0x58, 0x48, 0x89, 0xb5, 0xa2, 0xac, 0x08, 0x2c, 0x5c, 0x70, 0x16, 0x53,
	0x6b, 0xdd, 0xb3, 0x62, 0x16, 0x10, 0x6a, 0xa6, 0x19, 0xe3, 0x0c, 0x9d, 0x4a, 0xc6, 0x14, 0x8c,
	0x29, 0x19, 0x73, 0xdd, 0xd3, 0x8d, 0x3d, 0x5e, 0x9e, 0xe1, 0x24, 0xa7, 0x98, 0x47, 0x2c, 0x91,
	0x1b, 0xf4, 0xa7, 0x15, 0x29, 0x6e, 0xcb, 0xe2, 0x9b, 0xc5, 0xa3, 0x98, 0xe4, 0x1c, 0xc7, 0x69,
	0x05, 0x9c, 0x55, 0x00, 0x4e, 0x23, 0x0b, 0x27, 0x09, 0xe3, 0xc2, 0x9d, 0x4b, 0xb5, 0xfb, 0xa3,
	0x01, 0x8d, 0x49, 0x19, 0x08, 0xa5, 0xa0, 0xdf, 0xd8, 0xee, 0x8b, 0x94, 0x7e, 0x4c, 0x38, 0x0e,
	0x30, 0xc7, 0x5a, 0xab, 0xa3, 0x18, 0x6a, 0xff, 0xa5, 0xf9, 0xef, 0xbc, 0xa6, 0xf7, 0xd7, 0x29,
	0xb6, 0x4d, 0x2a, 0xdf, 0x87, 0x03, 0x57, 0xe3, 0x7b, 0x34, 0x84, 0xe0, 0x41, 0x82, 0x63, 0xa2,
	0x29, 0x1d, 0xc5, 0x38, 0x76, 0xc5, 0x19, 0x3d, 0x83, 0x93, 0x20, 0xca, 0x53, 0x8a, 0x37, 0xbe,
	0xd0, 0x6a, 0x42, 0x53, 0xab, 0xd9, 0xb4, 0x44, 0x9e, 0x00, 0x94, 0xf6, 0x9c, 0x70, 0x3f, 0x0a,
	0xb4, 0xba, 0x00, 0x8e, 0xab, 0xc9, 0x28, 0x40, 0xe7, 0xa0, 0xae, 0x32, 0x82, 0x39, 0xf1, 0xcb,
	0x26, 0xb4, 0x87, 0x22, 0xb8, 0xbe, 0x0d, 0xbe, 0xad, 0xc9, 0xf4, 0xb6, 0x35, 0xb9, 0x20, 0xf1,
	0x72, 0x50, 0x9a, 0x8b, 0x34, 0xd8, 0x99, 0xd5, 0xfb, 0xcd, 0x12, 0x17, 0xe6, 0xcf, 0xd0, 0x0e,
	0x48, 0x4a, 0xd9, 0x26, 0x26, 0x09, 0xf7, 0x73, 0x8e, 0x39, 0xd1, 0x8e, 0x3a, 0x8a, 0xd1, 0xec,
	0xbf, 0xd8, 0xd7, 0x9b, 0x28, 0xc4, 0x1c, 0xee, 0x5c, 0xf3, 0xd2, 0xe4, 0xb6, 0x82, 0xdb, 0x03,
	0x34, 0x80, 0x43, 0x8a, 0x97, 0x84, 0xe6, 0x5a, 0xb7, 0x53, 0x37, 0xd4, 0xfe, 0xf3, 0xff, 0xef,
	0x1b, 0x0b, 0xd6, 0x49, 0x78, 0xb6, 0x71, 0x2b, 0xa3, 0xfe, 0x06, 0xd4, 0x1b, 0x63, 0xd4, 0x86,
	0xfa, 0x77, 0xb2, 0xa9, 0xaa, 0x2f, 0x8f, 0xe8, 0x11, 0x34, 0xd6, 0x98, 0x16, 0xdb, 0xca, 0xe5,
	0xc5, 0xae, 0xbd, 0x56, 0xba, 0x9f, 0xa0, 0x75, 0x27, 0x21, 0xea, 0xc0, 0xd9, 0xd0, 0x99, 0x8d,
	0x3f, 0x7e, 0x99, 0x38, 0x53, 0xcf, 0x9f, 0x7b, 0x03, 0xcf, 0xf1, 0xaf, 0xa6, 0xf3, 0x99, 0x73,
	0x39, 0x7a, 0x37, 0x72, 0x86, 0xed, 0x03, 0x74, 0x02, 0x47, 0x92, 0x70, 0x86, 0x6d, 0x05, 0x35,
	0x01, 0xae, 0xa6, 0xbb, 0x7b, 0xed, 0xa2, 0x0d, 0xcd, 0xdb, 0x1f, 0xd8, 0xc5, 0x4f, 0x05, 0xf4,
	0x15, 0x8b, 0xf7, 0x3c, 0xd8, 0x4c, 0xf9, 0xfa, 0xb6, 0x52, 0x42, 0x46, 0x71, 0x12, 0x9a, 0x2c,
	0x0b, 0xad, 0x90, 0x24, 0xe2, 0x95, 0x58, 0x52, 0xc2, 0x69, 0x94, 0xdf, 0xfd, 0x63, 0xce, 0xe5,
	0xe9, 0x57, 0xed, 0xf4, 0xbd, 0x60, 0x16, 0x97, 0xa5, 0xbe, 0x18, 0x14, 0x9c, 0x4d, 0xe8, 0xe2,
	0xba, 0xf7, 0xbb, 0xf6, 0x58, 0x0a, 0xb6, 0x2d, 0x14, 0xdb, 0x16, 0xd2, 0xd8, 0xb6, 0xaf, 0x7b,
	0xcb, 0x43, 0xb1, 0xfd, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xd5, 0xa7, 0x63, 0xca,
	0x03, 0x00, 0x00,
}
