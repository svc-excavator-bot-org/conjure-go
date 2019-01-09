// This file was generated by Conjure and should not be manually edited.

package logging

import (
	"encoding/json"
	"fmt"
)

type WrappedLogV1Payload struct {
	typ             string
	serviceLogV1    *ServiceLogV1
	requestLogV2    *RequestLogV2
	traceLogV1      *TraceLogV1
	eventLogV2      *EventLogV2
	metricLogV1     *MetricLogV1
	auditLogV2      *AuditLogV2
	diagnosticLogV1 *DiagnosticLogV1
}

type wrappedLogV1PayloadDeserializer struct {
	Type            string           `json:"type" yaml:"type"`
	ServiceLogV1    *ServiceLogV1    `json:"serviceLogV1" yaml:"serviceLogV1"`
	RequestLogV2    *RequestLogV2    `json:"requestLogV2" yaml:"requestLogV2"`
	TraceLogV1      *TraceLogV1      `json:"traceLogV1" yaml:"traceLogV1"`
	EventLogV2      *EventLogV2      `json:"eventLogV2" yaml:"eventLogV2"`
	MetricLogV1     *MetricLogV1     `json:"metricLogV1" yaml:"metricLogV1"`
	AuditLogV2      *AuditLogV2      `json:"auditLogV2" yaml:"auditLogV2"`
	DiagnosticLogV1 *DiagnosticLogV1 `json:"diagnosticLogV1" yaml:"diagnosticLogV1"`
}

func (u *wrappedLogV1PayloadDeserializer) toStruct() WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: u.Type, serviceLogV1: u.ServiceLogV1, requestLogV2: u.RequestLogV2, traceLogV1: u.TraceLogV1, eventLogV2: u.EventLogV2, metricLogV1: u.MetricLogV1, auditLogV2: u.AuditLogV2, diagnosticLogV1: u.DiagnosticLogV1}
}

func (u *WrappedLogV1Payload) toSerializer() (interface{}, error) {
	switch u.typ {
	default:
		return nil, fmt.Errorf("unknown type %s", u.typ)
	case "serviceLogV1":
		return struct {
			Type         string       `json:"type" yaml:"type"`
			ServiceLogV1 ServiceLogV1 `json:"serviceLogV1" yaml:"serviceLogV1"`
		}{Type: "serviceLogV1", ServiceLogV1: *u.serviceLogV1}, nil
	case "requestLogV2":
		return struct {
			Type         string       `json:"type" yaml:"type"`
			RequestLogV2 RequestLogV2 `json:"requestLogV2" yaml:"requestLogV2"`
		}{Type: "requestLogV2", RequestLogV2: *u.requestLogV2}, nil
	case "traceLogV1":
		return struct {
			Type       string     `json:"type" yaml:"type"`
			TraceLogV1 TraceLogV1 `json:"traceLogV1" yaml:"traceLogV1"`
		}{Type: "traceLogV1", TraceLogV1: *u.traceLogV1}, nil
	case "eventLogV2":
		return struct {
			Type       string     `json:"type" yaml:"type"`
			EventLogV2 EventLogV2 `json:"eventLogV2" yaml:"eventLogV2"`
		}{Type: "eventLogV2", EventLogV2: *u.eventLogV2}, nil
	case "metricLogV1":
		return struct {
			Type        string      `json:"type" yaml:"type"`
			MetricLogV1 MetricLogV1 `json:"metricLogV1" yaml:"metricLogV1"`
		}{Type: "metricLogV1", MetricLogV1: *u.metricLogV1}, nil
	case "auditLogV2":
		return struct {
			Type       string     `json:"type" yaml:"type"`
			AuditLogV2 AuditLogV2 `json:"auditLogV2" yaml:"auditLogV2"`
		}{Type: "auditLogV2", AuditLogV2: *u.auditLogV2}, nil
	case "diagnosticLogV1":
		return struct {
			Type            string          `json:"type" yaml:"type"`
			DiagnosticLogV1 DiagnosticLogV1 `json:"diagnosticLogV1" yaml:"diagnosticLogV1"`
		}{Type: "diagnosticLogV1", DiagnosticLogV1: *u.diagnosticLogV1}, nil
	}
}

func (u WrappedLogV1Payload) MarshalJSON() ([]byte, error) {
	ser, err := u.toSerializer()
	if err != nil {
		return nil, err
	}
	return json.Marshal(ser)
}

func (u *WrappedLogV1Payload) UnmarshalJSON(data []byte) error {
	var deser wrappedLogV1PayloadDeserializer
	if err := json.Unmarshal(data, &deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u WrappedLogV1Payload) MarshalYAML() (interface{}, error) {
	return u.toSerializer()
}

func (u *WrappedLogV1Payload) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var deser wrappedLogV1PayloadDeserializer
	if err := unmarshal(&deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u *WrappedLogV1Payload) Accept(v WrappedLogV1PayloadVisitor) error {
	switch u.typ {
	default:
		if u.typ == "" {
			return fmt.Errorf("invalid value in union type")
		}
		return v.VisitUnknown(u.typ)
	case "serviceLogV1":
		return v.VisitServiceLogV1(*u.serviceLogV1)
	case "requestLogV2":
		return v.VisitRequestLogV2(*u.requestLogV2)
	case "traceLogV1":
		return v.VisitTraceLogV1(*u.traceLogV1)
	case "eventLogV2":
		return v.VisitEventLogV2(*u.eventLogV2)
	case "metricLogV1":
		return v.VisitMetricLogV1(*u.metricLogV1)
	case "auditLogV2":
		return v.VisitAuditLogV2(*u.auditLogV2)
	case "diagnosticLogV1":
		return v.VisitDiagnosticLogV1(*u.diagnosticLogV1)
	}
}

type WrappedLogV1PayloadVisitor interface {
	VisitServiceLogV1(v ServiceLogV1) error
	VisitRequestLogV2(v RequestLogV2) error
	VisitTraceLogV1(v TraceLogV1) error
	VisitEventLogV2(v EventLogV2) error
	VisitMetricLogV1(v MetricLogV1) error
	VisitAuditLogV2(v AuditLogV2) error
	VisitDiagnosticLogV1(v DiagnosticLogV1) error
	VisitUnknown(typeName string) error
}

func NewWrappedLogV1PayloadFromServiceLogV1(v ServiceLogV1) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "serviceLogV1", serviceLogV1: &v}
}

func NewWrappedLogV1PayloadFromRequestLogV2(v RequestLogV2) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "requestLogV2", requestLogV2: &v}
}

func NewWrappedLogV1PayloadFromTraceLogV1(v TraceLogV1) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "traceLogV1", traceLogV1: &v}
}

func NewWrappedLogV1PayloadFromEventLogV2(v EventLogV2) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "eventLogV2", eventLogV2: &v}
}

func NewWrappedLogV1PayloadFromMetricLogV1(v MetricLogV1) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "metricLogV1", metricLogV1: &v}
}

func NewWrappedLogV1PayloadFromAuditLogV2(v AuditLogV2) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "auditLogV2", auditLogV2: &v}
}

func NewWrappedLogV1PayloadFromDiagnosticLogV1(v DiagnosticLogV1) WrappedLogV1Payload {
	return WrappedLogV1Payload{typ: "diagnosticLogV1", diagnosticLogV1: &v}
}

type RequestLog struct {
	typ string
	v1  *RequestLogV1
	v2  *RequestLogV2
}

type requestLogDeserializer struct {
	Type string        `json:"type" yaml:"type"`
	V1   *RequestLogV1 `json:"v1" yaml:"v1"`
	V2   *RequestLogV2 `json:"v2" yaml:"v2"`
}

func (u *requestLogDeserializer) toStruct() RequestLog {
	return RequestLog{typ: u.Type, v1: u.V1, v2: u.V2}
}

func (u *RequestLog) toSerializer() (interface{}, error) {
	switch u.typ {
	default:
		return nil, fmt.Errorf("unknown type %s", u.typ)
	case "v1":
		return struct {
			Type string       `json:"type" yaml:"type"`
			V1   RequestLogV1 `json:"v1" yaml:"v1"`
		}{Type: "v1", V1: *u.v1}, nil
	case "v2":
		return struct {
			Type string       `json:"type" yaml:"type"`
			V2   RequestLogV2 `json:"v2" yaml:"v2"`
		}{Type: "v2", V2: *u.v2}, nil
	}
}

func (u RequestLog) MarshalJSON() ([]byte, error) {
	ser, err := u.toSerializer()
	if err != nil {
		return nil, err
	}
	return json.Marshal(ser)
}

func (u *RequestLog) UnmarshalJSON(data []byte) error {
	var deser requestLogDeserializer
	if err := json.Unmarshal(data, &deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u RequestLog) MarshalYAML() (interface{}, error) {
	return u.toSerializer()
}

func (u *RequestLog) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var deser requestLogDeserializer
	if err := unmarshal(&deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u *RequestLog) Accept(v RequestLogVisitor) error {
	switch u.typ {
	default:
		if u.typ == "" {
			return fmt.Errorf("invalid value in union type")
		}
		return v.VisitUnknown(u.typ)
	case "v1":
		return v.VisitV1(*u.v1)
	case "v2":
		return v.VisitV2(*u.v2)
	}
}

type RequestLogVisitor interface {
	VisitV1(v RequestLogV1) error
	VisitV2(v RequestLogV2) error
	VisitUnknown(typeName string) error
}

func NewRequestLogFromV1(v RequestLogV1) RequestLog {
	return RequestLog{typ: "v1", v1: &v}
}

func NewRequestLogFromV2(v RequestLogV2) RequestLog {
	return RequestLog{typ: "v2", v2: &v}
}

type Diagnostic struct {
	typ        string
	generic    *GenericDiagnostic
	threadDump *ThreadDumpV1
}

type diagnosticDeserializer struct {
	Type       string             `json:"type" yaml:"type"`
	Generic    *GenericDiagnostic `json:"generic" yaml:"generic"`
	ThreadDump *ThreadDumpV1      `json:"threadDump" yaml:"threadDump"`
}

func (u *diagnosticDeserializer) toStruct() Diagnostic {
	return Diagnostic{typ: u.Type, generic: u.Generic, threadDump: u.ThreadDump}
}

func (u *Diagnostic) toSerializer() (interface{}, error) {
	switch u.typ {
	default:
		return nil, fmt.Errorf("unknown type %s", u.typ)
	case "generic":
		return struct {
			Type    string            `json:"type" yaml:"type"`
			Generic GenericDiagnostic `json:"generic" yaml:"generic"`
		}{Type: "generic", Generic: *u.generic}, nil
	case "threadDump":
		return struct {
			Type       string       `json:"type" yaml:"type"`
			ThreadDump ThreadDumpV1 `json:"threadDump" yaml:"threadDump"`
		}{Type: "threadDump", ThreadDump: *u.threadDump}, nil
	}
}

func (u Diagnostic) MarshalJSON() ([]byte, error) {
	ser, err := u.toSerializer()
	if err != nil {
		return nil, err
	}
	return json.Marshal(ser)
}

func (u *Diagnostic) UnmarshalJSON(data []byte) error {
	var deser diagnosticDeserializer
	if err := json.Unmarshal(data, &deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u Diagnostic) MarshalYAML() (interface{}, error) {
	return u.toSerializer()
}

func (u *Diagnostic) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var deser diagnosticDeserializer
	if err := unmarshal(&deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u *Diagnostic) Accept(v DiagnosticVisitor) error {
	switch u.typ {
	default:
		if u.typ == "" {
			return fmt.Errorf("invalid value in union type")
		}
		return v.VisitUnknown(u.typ)
	case "generic":
		return v.VisitGeneric(*u.generic)
	case "threadDump":
		return v.VisitThreadDump(*u.threadDump)
	}
}

type DiagnosticVisitor interface {
	VisitGeneric(v GenericDiagnostic) error
	VisitThreadDump(v ThreadDumpV1) error
	VisitUnknown(typeName string) error
}

func NewDiagnosticFromGeneric(v GenericDiagnostic) Diagnostic {
	return Diagnostic{typ: "generic", generic: &v}
}

func NewDiagnosticFromThreadDump(v ThreadDumpV1) Diagnostic {
	return Diagnostic{typ: "threadDump", threadDump: &v}
}

// Union type containing log types that are logged to event.log.
type UnionEventLog struct {
	typ        string
	eventLog   *EventLogV1
	eventLogV2 *EventLogV2
	beaconLog  *BeaconLogV1
}

type unionEventLogDeserializer struct {
	Type       string       `json:"type" yaml:"type"`
	EventLog   *EventLogV1  `json:"eventLog" yaml:"eventLog"`
	EventLogV2 *EventLogV2  `json:"eventLogV2" yaml:"eventLogV2"`
	BeaconLog  *BeaconLogV1 `json:"beaconLog" yaml:"beaconLog"`
}

func (u *unionEventLogDeserializer) toStruct() UnionEventLog {
	return UnionEventLog{typ: u.Type, eventLog: u.EventLog, eventLogV2: u.EventLogV2, beaconLog: u.BeaconLog}
}

func (u *UnionEventLog) toSerializer() (interface{}, error) {
	switch u.typ {
	default:
		return nil, fmt.Errorf("unknown type %s", u.typ)
	case "eventLog":
		return struct {
			Type     string     `json:"type" yaml:"type"`
			EventLog EventLogV1 `json:"eventLog" yaml:"eventLog"`
		}{Type: "eventLog", EventLog: *u.eventLog}, nil
	case "eventLogV2":
		return struct {
			Type       string     `json:"type" yaml:"type"`
			EventLogV2 EventLogV2 `json:"eventLogV2" yaml:"eventLogV2"`
		}{Type: "eventLogV2", EventLogV2: *u.eventLogV2}, nil
	case "beaconLog":
		return struct {
			Type      string      `json:"type" yaml:"type"`
			BeaconLog BeaconLogV1 `json:"beaconLog" yaml:"beaconLog"`
		}{Type: "beaconLog", BeaconLog: *u.beaconLog}, nil
	}
}

func (u UnionEventLog) MarshalJSON() ([]byte, error) {
	ser, err := u.toSerializer()
	if err != nil {
		return nil, err
	}
	return json.Marshal(ser)
}

func (u *UnionEventLog) UnmarshalJSON(data []byte) error {
	var deser unionEventLogDeserializer
	if err := json.Unmarshal(data, &deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u UnionEventLog) MarshalYAML() (interface{}, error) {
	return u.toSerializer()
}

func (u *UnionEventLog) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var deser unionEventLogDeserializer
	if err := unmarshal(&deser); err != nil {
		return err
	}
	*u = deser.toStruct()
	return nil
}

func (u *UnionEventLog) Accept(v UnionEventLogVisitor) error {
	switch u.typ {
	default:
		if u.typ == "" {
			return fmt.Errorf("invalid value in union type")
		}
		return v.VisitUnknown(u.typ)
	case "eventLog":
		return v.VisitEventLog(*u.eventLog)
	case "eventLogV2":
		return v.VisitEventLogV2(*u.eventLogV2)
	case "beaconLog":
		return v.VisitBeaconLog(*u.beaconLog)
	}
}

type UnionEventLogVisitor interface {
	VisitEventLog(v EventLogV1) error
	VisitEventLogV2(v EventLogV2) error
	VisitBeaconLog(v BeaconLogV1) error
	VisitUnknown(typeName string) error
}

func NewUnionEventLogFromEventLog(v EventLogV1) UnionEventLog {
	return UnionEventLog{typ: "eventLog", eventLog: &v}
}

func NewUnionEventLogFromEventLogV2(v EventLogV2) UnionEventLog {
	return UnionEventLog{typ: "eventLogV2", eventLogV2: &v}
}

func NewUnionEventLogFromBeaconLog(v BeaconLogV1) UnionEventLog {
	return UnionEventLog{typ: "beaconLog", beaconLog: &v}
}
