// This file was generated by Conjure and should not be manually edited.

package spec

import (
	"github.com/palantir/pkg/safejson"
	"github.com/palantir/pkg/safeyaml"
)

type ConjureDefinition struct {
	Version  int                 `json:"version"`
	Errors   []ErrorDefinition   `json:"errors"`
	Types    []TypeDefinition    `json:"types"`
	Services []ServiceDefinition `json:"services"`
}

func (o ConjureDefinition) MarshalJSON() ([]byte, error) {
	if o.Errors == nil {
		o.Errors = make([]ErrorDefinition, 0)
	}
	if o.Types == nil {
		o.Types = make([]TypeDefinition, 0)
	}
	if o.Services == nil {
		o.Services = make([]ServiceDefinition, 0)
	}
	type ConjureDefinitionAlias ConjureDefinition
	return safejson.Marshal(ConjureDefinitionAlias(o))
}

func (o *ConjureDefinition) UnmarshalJSON(data []byte) error {
	type ConjureDefinitionAlias ConjureDefinition
	var rawConjureDefinition ConjureDefinitionAlias
	if err := safejson.Unmarshal(data, &rawConjureDefinition); err != nil {
		return err
	}
	if rawConjureDefinition.Errors == nil {
		rawConjureDefinition.Errors = make([]ErrorDefinition, 0)
	}
	if rawConjureDefinition.Types == nil {
		rawConjureDefinition.Types = make([]TypeDefinition, 0)
	}
	if rawConjureDefinition.Services == nil {
		rawConjureDefinition.Services = make([]ServiceDefinition, 0)
	}
	*o = ConjureDefinition(rawConjureDefinition)
	return nil
}

func (o ConjureDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ConjureDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type TypeName struct {
	// The name of the custom Conjure type or service. It must be in UpperCamelCase. Numbers are permitted, but not at the beginning of a word. Allowed names: "FooBar", "XYCoordinate", "Build2Request". Disallowed names: "fooBar", "2BuildRequest".
	Name string `json:"name" conjure-docs:"The name of the custom Conjure type or service. It must be in UpperCamelCase. Numbers are permitted, but not at the beginning of a word. Allowed names: \"FooBar\", \"XYCoordinate\", \"Build2Request\". Disallowed names: \"fooBar\", \"2BuildRequest\".\n"`
	// A period-delimited string of package names. The package names must be lowercase. Numbers are permitted, but not at the beginning of a package name. Allowed packages: "foo", "com.palantir.bar", "com.palantir.foo.thing2". Disallowed packages: "Foo", "com.palantir.foo.2thing".
	Package string `json:"package" conjure-docs:"A period-delimited string of package names. The package names must be lowercase. Numbers are permitted, but not at the beginning of a package name. Allowed packages: \"foo\", \"com.palantir.bar\", \"com.palantir.foo.thing2\". Disallowed packages: \"Foo\", \"com.palantir.foo.2thing\".\n"`
}

func (o TypeName) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *TypeName) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ErrorDefinition struct {
	ErrorName  TypeName          `json:"errorName"`
	Docs       *Documentation    `json:"docs"`
	Namespace  ErrorNamespace    `json:"namespace"`
	Code       ErrorCode         `json:"code"`
	SafeArgs   []FieldDefinition `json:"safeArgs"`
	UnsafeArgs []FieldDefinition `json:"unsafeArgs"`
}

func (o ErrorDefinition) MarshalJSON() ([]byte, error) {
	if o.SafeArgs == nil {
		o.SafeArgs = make([]FieldDefinition, 0)
	}
	if o.UnsafeArgs == nil {
		o.UnsafeArgs = make([]FieldDefinition, 0)
	}
	type ErrorDefinitionAlias ErrorDefinition
	return safejson.Marshal(ErrorDefinitionAlias(o))
}

func (o *ErrorDefinition) UnmarshalJSON(data []byte) error {
	type ErrorDefinitionAlias ErrorDefinition
	var rawErrorDefinition ErrorDefinitionAlias
	if err := safejson.Unmarshal(data, &rawErrorDefinition); err != nil {
		return err
	}
	if rawErrorDefinition.SafeArgs == nil {
		rawErrorDefinition.SafeArgs = make([]FieldDefinition, 0)
	}
	if rawErrorDefinition.UnsafeArgs == nil {
		rawErrorDefinition.UnsafeArgs = make([]FieldDefinition, 0)
	}
	*o = ErrorDefinition(rawErrorDefinition)
	return nil
}

func (o ErrorDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ErrorDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ExternalReference struct {
	// An identifier for a non-Conjure type which is already defined in a different language (e.g. Java).
	ExternalReference TypeName `json:"externalReference" conjure-docs:"An identifier for a non-Conjure type which is already defined in a different language (e.g. Java)."`
	// Other language generators may use the provided fallback if the non-Conjure type is not available. The ANY PrimitiveType is permissible for all external types, but a more specific definition is preferrable.
	Fallback Type `json:"fallback" conjure-docs:"Other language generators may use the provided fallback if the non-Conjure type is not available. The ANY PrimitiveType is permissible for all external types, but a more specific definition is preferrable.\n"`
}

func (o ExternalReference) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ExternalReference) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type OptionalType struct {
	ItemType Type `json:"itemType"`
}

func (o OptionalType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *OptionalType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ListType struct {
	ItemType Type `json:"itemType"`
}

func (o ListType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ListType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type SetType struct {
	ItemType Type `json:"itemType"`
}

func (o SetType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *SetType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type MapType struct {
	KeyType   Type `json:"keyType"`
	ValueType Type `json:"valueType"`
}

func (o MapType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *MapType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type AliasDefinition struct {
	TypeName TypeName       `json:"typeName"`
	Alias    Type           `json:"alias"`
	Docs     *Documentation `json:"docs"`
}

func (o AliasDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *AliasDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type EnumDefinition struct {
	TypeName TypeName              `json:"typeName"`
	Values   []EnumValueDefinition `json:"values"`
	Docs     *Documentation        `json:"docs"`
}

func (o EnumDefinition) MarshalJSON() ([]byte, error) {
	if o.Values == nil {
		o.Values = make([]EnumValueDefinition, 0)
	}
	type EnumDefinitionAlias EnumDefinition
	return safejson.Marshal(EnumDefinitionAlias(o))
}

func (o *EnumDefinition) UnmarshalJSON(data []byte) error {
	type EnumDefinitionAlias EnumDefinition
	var rawEnumDefinition EnumDefinitionAlias
	if err := safejson.Unmarshal(data, &rawEnumDefinition); err != nil {
		return err
	}
	if rawEnumDefinition.Values == nil {
		rawEnumDefinition.Values = make([]EnumValueDefinition, 0)
	}
	*o = EnumDefinition(rawEnumDefinition)
	return nil
}

func (o EnumDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *EnumDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ObjectDefinition struct {
	TypeName TypeName          `json:"typeName"`
	Fields   []FieldDefinition `json:"fields"`
	Docs     *Documentation    `json:"docs"`
}

func (o ObjectDefinition) MarshalJSON() ([]byte, error) {
	if o.Fields == nil {
		o.Fields = make([]FieldDefinition, 0)
	}
	type ObjectDefinitionAlias ObjectDefinition
	return safejson.Marshal(ObjectDefinitionAlias(o))
}

func (o *ObjectDefinition) UnmarshalJSON(data []byte) error {
	type ObjectDefinitionAlias ObjectDefinition
	var rawObjectDefinition ObjectDefinitionAlias
	if err := safejson.Unmarshal(data, &rawObjectDefinition); err != nil {
		return err
	}
	if rawObjectDefinition.Fields == nil {
		rawObjectDefinition.Fields = make([]FieldDefinition, 0)
	}
	*o = ObjectDefinition(rawObjectDefinition)
	return nil
}

func (o ObjectDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ObjectDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type UnionDefinition struct {
	TypeName TypeName          `json:"typeName"`
	Union    []FieldDefinition `json:"union"`
	Docs     *Documentation    `json:"docs"`
}

func (o UnionDefinition) MarshalJSON() ([]byte, error) {
	if o.Union == nil {
		o.Union = make([]FieldDefinition, 0)
	}
	type UnionDefinitionAlias UnionDefinition
	return safejson.Marshal(UnionDefinitionAlias(o))
}

func (o *UnionDefinition) UnmarshalJSON(data []byte) error {
	type UnionDefinitionAlias UnionDefinition
	var rawUnionDefinition UnionDefinitionAlias
	if err := safejson.Unmarshal(data, &rawUnionDefinition); err != nil {
		return err
	}
	if rawUnionDefinition.Union == nil {
		rawUnionDefinition.Union = make([]FieldDefinition, 0)
	}
	*o = UnionDefinition(rawUnionDefinition)
	return nil
}

func (o UnionDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *UnionDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type EnumValueDefinition struct {
	Value string         `json:"value"`
	Docs  *Documentation `json:"docs"`
}

func (o EnumValueDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *EnumValueDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type FieldDefinition struct {
	FieldName FieldName      `json:"fieldName"`
	Type      Type           `json:"type"`
	Docs      *Documentation `json:"docs"`
}

func (o FieldDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *FieldDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ServiceDefinition struct {
	ServiceName TypeName             `json:"serviceName"`
	Endpoints   []EndpointDefinition `json:"endpoints"`
	Docs        *Documentation       `json:"docs"`
}

func (o ServiceDefinition) MarshalJSON() ([]byte, error) {
	if o.Endpoints == nil {
		o.Endpoints = make([]EndpointDefinition, 0)
	}
	type ServiceDefinitionAlias ServiceDefinition
	return safejson.Marshal(ServiceDefinitionAlias(o))
}

func (o *ServiceDefinition) UnmarshalJSON(data []byte) error {
	type ServiceDefinitionAlias ServiceDefinition
	var rawServiceDefinition ServiceDefinitionAlias
	if err := safejson.Unmarshal(data, &rawServiceDefinition); err != nil {
		return err
	}
	if rawServiceDefinition.Endpoints == nil {
		rawServiceDefinition.Endpoints = make([]EndpointDefinition, 0)
	}
	*o = ServiceDefinition(rawServiceDefinition)
	return nil
}

func (o ServiceDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ServiceDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type EndpointDefinition struct {
	EndpointName EndpointName         `json:"endpointName"`
	HttpMethod   HttpMethod           `json:"httpMethod"`
	HttpPath     HttpPath             `json:"httpPath"`
	Auth         *AuthType            `json:"auth"`
	Args         []ArgumentDefinition `json:"args"`
	Returns      *Type                `json:"returns"`
	Docs         *Documentation       `json:"docs"`
	Deprecated   *Documentation       `json:"deprecated"`
	Markers      []Type               `json:"markers"`
}

func (o EndpointDefinition) MarshalJSON() ([]byte, error) {
	if o.Args == nil {
		o.Args = make([]ArgumentDefinition, 0)
	}
	if o.Markers == nil {
		o.Markers = make([]Type, 0)
	}
	type EndpointDefinitionAlias EndpointDefinition
	return safejson.Marshal(EndpointDefinitionAlias(o))
}

func (o *EndpointDefinition) UnmarshalJSON(data []byte) error {
	type EndpointDefinitionAlias EndpointDefinition
	var rawEndpointDefinition EndpointDefinitionAlias
	if err := safejson.Unmarshal(data, &rawEndpointDefinition); err != nil {
		return err
	}
	if rawEndpointDefinition.Args == nil {
		rawEndpointDefinition.Args = make([]ArgumentDefinition, 0)
	}
	if rawEndpointDefinition.Markers == nil {
		rawEndpointDefinition.Markers = make([]Type, 0)
	}
	*o = EndpointDefinition(rawEndpointDefinition)
	return nil
}

func (o EndpointDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *EndpointDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type HeaderAuthType struct {
}

func (o HeaderAuthType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *HeaderAuthType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type CookieAuthType struct {
	CookieName string `json:"cookieName"`
}

func (o CookieAuthType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *CookieAuthType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ArgumentDefinition struct {
	ArgName   ArgumentName   `json:"argName"`
	Type      Type           `json:"type"`
	ParamType ParameterType  `json:"paramType"`
	Docs      *Documentation `json:"docs"`
	Markers   []Type         `json:"markers"`
}

func (o ArgumentDefinition) MarshalJSON() ([]byte, error) {
	if o.Markers == nil {
		o.Markers = make([]Type, 0)
	}
	type ArgumentDefinitionAlias ArgumentDefinition
	return safejson.Marshal(ArgumentDefinitionAlias(o))
}

func (o *ArgumentDefinition) UnmarshalJSON(data []byte) error {
	type ArgumentDefinitionAlias ArgumentDefinition
	var rawArgumentDefinition ArgumentDefinitionAlias
	if err := safejson.Unmarshal(data, &rawArgumentDefinition); err != nil {
		return err
	}
	if rawArgumentDefinition.Markers == nil {
		rawArgumentDefinition.Markers = make([]Type, 0)
	}
	*o = ArgumentDefinition(rawArgumentDefinition)
	return nil
}

func (o ArgumentDefinition) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ArgumentDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type BodyParameterType struct {
}

func (o BodyParameterType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *BodyParameterType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type HeaderParameterType struct {
	ParamId ParameterId `json:"paramId"`
}

func (o HeaderParameterType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *HeaderParameterType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type PathParameterType struct {
}

func (o PathParameterType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *PathParameterType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type QueryParameterType struct {
	ParamId ParameterId `json:"paramId"`
}

func (o QueryParameterType) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *QueryParameterType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}
