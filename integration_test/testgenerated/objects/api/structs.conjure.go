// This file was generated by Conjure and should not be manually edited.

package api

import (
	"github.com/palantir/pkg/binary"
	"github.com/palantir/pkg/safejson"
	"github.com/palantir/pkg/safeyaml"
	"github.com/palantir/pkg/uuid"
)

type Basic struct {
	Data string `json:"data"`
}

func (o Basic) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *Basic) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type Collections struct {
	MapVar   map[string][]int   `json:"mapVar"`
	ListVar  []string           `json:"listVar"`
	MultiDim [][]map[string]int `json:"multiDim"`
}

func (o Collections) MarshalJSON() ([]byte, error) {
	if o.MapVar == nil {
		o.MapVar = make(map[string][]int, 0)
	}
	if o.ListVar == nil {
		o.ListVar = make([]string, 0)
	}
	if o.MultiDim == nil {
		o.MultiDim = make([][]map[string]int, 0)
	}
	type CollectionsAlias Collections
	return safejson.Marshal(CollectionsAlias(o))
}

func (o *Collections) UnmarshalJSON(data []byte) error {
	type CollectionsAlias Collections
	var rawCollections CollectionsAlias
	if err := safejson.Unmarshal(data, &rawCollections); err != nil {
		return err
	}
	if rawCollections.MapVar == nil {
		rawCollections.MapVar = make(map[string][]int, 0)
	}
	if rawCollections.ListVar == nil {
		rawCollections.ListVar = make([]string, 0)
	}
	if rawCollections.MultiDim == nil {
		rawCollections.MultiDim = make([][]map[string]int, 0)
	}
	*o = Collections(rawCollections)
	return nil
}

func (o Collections) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *Collections) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type Compound struct {
	Obj Collections `json:"obj"`
}

func (o Compound) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *Compound) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type ExampleUuid struct {
	Uid uuid.UUID `json:"uid"`
}

func (o ExampleUuid) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *ExampleUuid) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}

type BinaryMap struct {
	Map map[binary.Binary][]byte `json:"map"`
}

func (o BinaryMap) MarshalJSON() ([]byte, error) {
	if o.Map == nil {
		o.Map = make(map[binary.Binary][]byte, 0)
	}
	type BinaryMapAlias BinaryMap
	return safejson.Marshal(BinaryMapAlias(o))
}

func (o *BinaryMap) UnmarshalJSON(data []byte) error {
	type BinaryMapAlias BinaryMap
	var rawBinaryMap BinaryMapAlias
	if err := safejson.Unmarshal(data, &rawBinaryMap); err != nil {
		return err
	}
	if rawBinaryMap.Map == nil {
		rawBinaryMap.Map = make(map[binary.Binary][]byte, 0)
	}
	*o = BinaryMap(rawBinaryMap)
	return nil
}

func (o BinaryMap) MarshalYAML() (interface{}, error) {
	jsonBytes, err := safejson.Marshal(o)
	if err != nil {
		return nil, err
	}
	return safeyaml.JSONtoYAMLMapSlice(jsonBytes)
}

func (o *BinaryMap) UnmarshalYAML(unmarshal func(interface{}) error) error {
	jsonBytes, err := safeyaml.YAMLUnmarshalerToJSONBytes(unmarshal)
	if err != nil {
		return err
	}
	return safejson.Unmarshal(jsonBytes, *&o)
}
