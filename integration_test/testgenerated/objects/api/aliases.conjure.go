// This file was generated by Conjure and should not be manually edited.

package api

import (
	"github.com/palantir/pkg/safejson"
	"github.com/palantir/pkg/safeyaml"
	"github.com/palantir/pkg/uuid"
)

type OptionalUuidAlias struct {
	Value *uuid.UUID
}

func (a OptionalUuidAlias) MarshalText() ([]byte, error) {
	if a.Value == nil {
		return nil, nil
	}
	return a.Value.MarshalText()
}

func (a *OptionalUuidAlias) UnmarshalText(data []byte) error {
	if a.Value == nil {
		a.Value = new(uuid.UUID)
	}
	return a.Value.UnmarshalText(data)
}

func (a OptionalUuidAlias) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(a)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (a *OptionalUuidAlias) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&a)
}

type UuidAlias uuid.UUID

func (a UuidAlias) MarshalText() ([]byte, error) {
	return uuid.UUID(a).MarshalText()
}

func (a *UuidAlias) UnmarshalText(data []byte) error {
	var rawUuidAlias uuid.UUID
	if err := rawUuidAlias.UnmarshalText(data); err != nil {
		return err
	}
	*a = UuidAlias(rawUuidAlias)
	return nil
}

func (a UuidAlias) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(a)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (a *UuidAlias) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&a)
}

type UuidAlias2 Compound

func (a UuidAlias2) MarshalJSON() ([]byte, error) {
	return safejson.Marshal(Compound(a))
}

func (a *UuidAlias2) UnmarshalJSON(data []byte) error {
	var rawUuidAlias2 Compound
	if err := safejson.Unmarshal(data, &rawUuidAlias2); err != nil {
		return err
	}
	*a = UuidAlias2(rawUuidAlias2)
	return nil
}

func (a UuidAlias2) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(a)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (a *UuidAlias2) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&a)
}

type BinaryAlias []byte
