// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/real_estate_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Real Estate placeholder fields.
type RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField int32

const (
	// Not specified.
	RealEstatePlaceholderFieldEnum_UNSPECIFIED RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	RealEstatePlaceholderFieldEnum_UNKNOWN RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 1
	// Data Type: STRING. Unique ID.
	RealEstatePlaceholderFieldEnum_LISTING_ID RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 2
	// Data Type: STRING. Main headline with listing name to be shown in dynamic
	// ad.
	RealEstatePlaceholderFieldEnum_LISTING_NAME RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 3
	// Data Type: STRING. City name to be shown in dynamic ad.
	RealEstatePlaceholderFieldEnum_CITY_NAME RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 4
	// Data Type: STRING. Description of listing to be shown in dynamic ad.
	RealEstatePlaceholderFieldEnum_DESCRIPTION RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 5
	// Data Type: STRING. Complete listing address, including postal code.
	RealEstatePlaceholderFieldEnum_ADDRESS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 6
	// Data Type: STRING. Price to be shown in the ad.
	// Example: "100.00 USD"
	RealEstatePlaceholderFieldEnum_PRICE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 7
	// Data Type: STRING. Formatted price to be shown in the ad.
	// Example: "Starting at $100.00 USD", "$80 - $100"
	RealEstatePlaceholderFieldEnum_FORMATTED_PRICE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 8
	// Data Type: URL. Image to be displayed in the ad.
	RealEstatePlaceholderFieldEnum_IMAGE_URL RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 9
	// Data Type: STRING. Type of property (house, condo, apartment, etc.) used
	// to group like items together for recommendation engine.
	RealEstatePlaceholderFieldEnum_PROPERTY_TYPE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 10
	// Data Type: STRING. Type of listing (resale, rental, foreclosure, etc.)
	// used to group like items together for recommendation engine.
	RealEstatePlaceholderFieldEnum_LISTING_TYPE RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 11
	// Data Type: STRING_LIST. Keywords used for product retrieval.
	RealEstatePlaceholderFieldEnum_CONTEXTUAL_KEYWORDS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 12
	// Data Type: URL_LIST. Final URLs to be used in ad when using Upgraded
	// URLs; the more specific the better (e.g. the individual URL of a specific
	// listing and its location).
	RealEstatePlaceholderFieldEnum_FINAL_URLS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 13
	// Data Type: URL_LIST. Final mobile URLs for the ad when using Upgraded
	// URLs.
	RealEstatePlaceholderFieldEnum_FINAL_MOBILE_URLS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 14
	// Data Type: URL. Tracking template for the ad when using Upgraded URLs.
	RealEstatePlaceholderFieldEnum_TRACKING_URL RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 15
	// Data Type: STRING. Android app link. Must be formatted as:
	// android-app://{package_id}/{scheme}/{host_path}.
	// The components are defined as follows:
	// package_id: app ID as specified in Google Play.
	// scheme: the scheme to pass to the application. Can be HTTP, or a custom
	//   scheme.
	// host_path: identifies the specific content within your application.
	RealEstatePlaceholderFieldEnum_ANDROID_APP_LINK RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 16
	// Data Type: STRING_LIST. List of recommended listing IDs to show together
	// with this item.
	RealEstatePlaceholderFieldEnum_SIMILAR_LISTING_IDS RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 17
	// Data Type: STRING. iOS app link.
	RealEstatePlaceholderFieldEnum_IOS_APP_LINK RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 18
	// Data Type: INT64. iOS app store ID.
	RealEstatePlaceholderFieldEnum_IOS_APP_STORE_ID RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField = 19
)

var RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "LISTING_ID",
	3:  "LISTING_NAME",
	4:  "CITY_NAME",
	5:  "DESCRIPTION",
	6:  "ADDRESS",
	7:  "PRICE",
	8:  "FORMATTED_PRICE",
	9:  "IMAGE_URL",
	10: "PROPERTY_TYPE",
	11: "LISTING_TYPE",
	12: "CONTEXTUAL_KEYWORDS",
	13: "FINAL_URLS",
	14: "FINAL_MOBILE_URLS",
	15: "TRACKING_URL",
	16: "ANDROID_APP_LINK",
	17: "SIMILAR_LISTING_IDS",
	18: "IOS_APP_LINK",
	19: "IOS_APP_STORE_ID",
}
var RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"LISTING_ID":          2,
	"LISTING_NAME":        3,
	"CITY_NAME":           4,
	"DESCRIPTION":         5,
	"ADDRESS":             6,
	"PRICE":               7,
	"FORMATTED_PRICE":     8,
	"IMAGE_URL":           9,
	"PROPERTY_TYPE":       10,
	"LISTING_TYPE":        11,
	"CONTEXTUAL_KEYWORDS": 12,
	"FINAL_URLS":          13,
	"FINAL_MOBILE_URLS":   14,
	"TRACKING_URL":        15,
	"ANDROID_APP_LINK":    16,
	"SIMILAR_LISTING_IDS": 17,
	"IOS_APP_LINK":        18,
	"IOS_APP_STORE_ID":    19,
}

func (x RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) String() string {
	return proto.EnumName(RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_name, int32(x))
}
func (RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_real_estate_placeholder_field_c9d2f77c9d81d833, []int{0, 0}
}

// Values for Real Estate placeholder fields.
// For more information about dynamic remarketing feeds, see
// https://support.google.com/google-ads/answer/6053288.
type RealEstatePlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RealEstatePlaceholderFieldEnum) Reset()         { *m = RealEstatePlaceholderFieldEnum{} }
func (m *RealEstatePlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*RealEstatePlaceholderFieldEnum) ProtoMessage()    {}
func (*RealEstatePlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_real_estate_placeholder_field_c9d2f77c9d81d833, []int{0}
}
func (m *RealEstatePlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealEstatePlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *RealEstatePlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealEstatePlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *RealEstatePlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealEstatePlaceholderFieldEnum.Merge(dst, src)
}
func (m *RealEstatePlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_RealEstatePlaceholderFieldEnum.Size(m)
}
func (m *RealEstatePlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RealEstatePlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RealEstatePlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RealEstatePlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.RealEstatePlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField", RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_name, RealEstatePlaceholderFieldEnum_RealEstatePlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/real_estate_placeholder_field.proto", fileDescriptor_real_estate_placeholder_field_c9d2f77c9d81d833)
}

var fileDescriptor_real_estate_placeholder_field_c9d2f77c9d81d833 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xd1, 0x6e, 0xda, 0x3c,
	0x14, 0xc7, 0x3f, 0xe0, 0x6b, 0x3b, 0x4c, 0x29, 0xc6, 0x6c, 0x9a, 0x54, 0xad, 0x9b, 0xda, 0x07,
	0x48, 0x84, 0x76, 0x97, 0x5d, 0x99, 0xc4, 0x20, 0x8b, 0xe0, 0x58, 0xb6, 0xa1, 0x63, 0x42, 0xb2,
	0xb2, 0x26, 0xcb, 0x90, 0x42, 0x82, 0x08, 0xed, 0x23, 0xec, 0x41, 0x76, 0xb7, 0x3d, 0xca, 0x1e,
	0x65, 0xd7, 0x7b, 0x80, 0xc9, 0xc9, 0x02, 0xbb, 0x61, 0x37, 0xd1, 0x39, 0xff, 0x73, 0xf2, 0xfb,
	0x1f, 0xf9, 0x1c, 0x80, 0x93, 0x3c, 0x4f, 0xd2, 0xd8, 0x0e, 0xa3, 0xc2, 0xae, 0x42, 0x13, 0x3d,
	0x0d, 0xed, 0x38, 0x7b, 0xdc, 0x14, 0xf6, 0x2e, 0x0e, 0x53, 0x1d, 0x17, 0xfb, 0x70, 0x1f, 0xeb,
	0x6d, 0x1a, 0x3e, 0xc4, 0x9f, 0xf3, 0x34, 0x8a, 0x77, 0xfa, 0xd3, 0x3a, 0x4e, 0x23, 0x6b, 0xbb,
	0xcb, 0xf7, 0x39, 0xba, 0xa9, 0xfe, 0xb3, 0xc2, 0xa8, 0xb0, 0x0e, 0x08, 0xeb, 0x69, 0x68, 0x95,
	0x88, 0xeb, 0x57, 0xb5, 0xc3, 0x76, 0x6d, 0x87, 0x59, 0x96, 0xef, 0xc3, 0xfd, 0x3a, 0xcf, 0x8a,
	0xea, 0xe7, 0xbb, 0x6f, 0x2d, 0xf0, 0x5a, 0xc4, 0x61, 0x4a, 0x4a, 0x0f, 0x7e, 0xb4, 0x18, 0x1b,
	0x07, 0x92, 0x3d, 0x6e, 0xee, 0xbe, 0xb4, 0xc0, 0xf5, 0xe9, 0x16, 0xd4, 0x03, 0x9d, 0x39, 0x93,
	0x9c, 0xb8, 0x74, 0x4c, 0x89, 0x07, 0xff, 0x43, 0x1d, 0x70, 0x31, 0x67, 0x53, 0x16, 0xdc, 0x33,
	0xd8, 0x40, 0x57, 0x00, 0xf8, 0x54, 0x2a, 0xca, 0x26, 0x9a, 0x7a, 0xb0, 0x89, 0x20, 0xb8, 0xac,
	0x73, 0x86, 0x67, 0x04, 0xb6, 0x50, 0x17, 0xb4, 0x5d, 0xaa, 0x96, 0x55, 0xfa, 0xbf, 0xc1, 0x79,
	0x44, 0xba, 0x82, 0x72, 0x45, 0x03, 0x06, 0xcf, 0x0c, 0x0e, 0x7b, 0x9e, 0x20, 0x52, 0xc2, 0x73,
	0xd4, 0x06, 0x67, 0x5c, 0x50, 0x97, 0xc0, 0x0b, 0x34, 0x00, 0xbd, 0x71, 0x20, 0x66, 0x58, 0x29,
	0xe2, 0xe9, 0x4a, 0x7c, 0x66, 0x60, 0x74, 0x86, 0x27, 0x44, 0xcf, 0x85, 0x0f, 0xdb, 0xa8, 0x0f,
	0xba, 0x5c, 0x04, 0x9c, 0x08, 0xb5, 0xd4, 0x6a, 0xc9, 0x09, 0x04, 0x7f, 0x0f, 0x50, 0x2a, 0x1d,
	0xf4, 0x12, 0x0c, 0xdc, 0x80, 0x29, 0xf2, 0x5e, 0xcd, 0xb1, 0xaf, 0xa7, 0x64, 0x79, 0x1f, 0x08,
	0x4f, 0xc2, 0x4b, 0x33, 0xfb, 0x98, 0x32, 0xec, 0x1b, 0x98, 0x84, 0x5d, 0xf4, 0x02, 0xf4, 0xab,
	0x7c, 0x16, 0x8c, 0xa8, 0x4f, 0x2a, 0xf9, 0xca, 0x10, 0x95, 0xc0, 0xee, 0xd4, 0x20, 0x8d, 0x6d,
	0x0f, 0x3d, 0x07, 0x10, 0x33, 0x4f, 0x04, 0xd4, 0xd3, 0x98, 0x73, 0xed, 0x53, 0x36, 0x85, 0xd0,
	0xf8, 0x48, 0x3a, 0xa3, 0x3e, 0x16, 0xfa, 0xf8, 0x24, 0x12, 0xf6, 0x0d, 0x80, 0x06, 0xf2, 0xd8,
	0x8a, 0x0c, 0xa0, 0x56, 0xa4, 0x0a, 0x04, 0x31, 0x6f, 0x37, 0x18, 0xfd, 0x6a, 0x80, 0xdb, 0x87,
	0x7c, 0x63, 0xfd, 0x73, 0xdf, 0xa3, 0x37, 0xa7, 0x77, 0xc5, 0xcd, 0xca, 0x79, 0xe3, 0xc3, 0xe8,
	0x0f, 0x21, 0xc9, 0xd3, 0x30, 0x4b, 0xac, 0x7c, 0x97, 0xd8, 0x49, 0x9c, 0x95, 0x07, 0x51, 0x1f,
	0xe1, 0x76, 0x5d, 0x9c, 0xb8, 0xc9, 0x77, 0xe5, 0xf7, 0x6b, 0xb3, 0x35, 0xc1, 0xf8, 0x7b, 0xf3,
	0x66, 0x52, 0xa1, 0x70, 0x54, 0x58, 0x55, 0x68, 0xa2, 0xc5, 0xd0, 0x32, 0xa7, 0x53, 0xfc, 0xa8,
	0xeb, 0x2b, 0x1c, 0x15, 0xab, 0x43, 0x7d, 0xb5, 0x18, 0xae, 0xca, 0xfa, 0xcf, 0xe6, 0x6d, 0x25,
	0x3a, 0x0e, 0x8e, 0x0a, 0xc7, 0x39, 0x74, 0x38, 0xce, 0x62, 0xe8, 0x38, 0x65, 0xcf, 0xc7, 0xf3,
	0x72, 0xb0, 0xb7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xfe, 0xe3, 0x31, 0x2b, 0x03, 0x00,
	0x00,
}
