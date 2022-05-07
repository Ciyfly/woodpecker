// This file is generated by "./lib/proto/generate"

package proto

/*

Log

Provides access to log entries.

*/

// LogLogEntrySource enum
type LogLogEntrySource string

const (
	// LogLogEntrySourceXML enum const
	LogLogEntrySourceXML LogLogEntrySource = "xml"

	// LogLogEntrySourceJavascript enum const
	LogLogEntrySourceJavascript LogLogEntrySource = "javascript"

	// LogLogEntrySourceNetwork enum const
	LogLogEntrySourceNetwork LogLogEntrySource = "network"

	// LogLogEntrySourceStorage enum const
	LogLogEntrySourceStorage LogLogEntrySource = "storage"

	// LogLogEntrySourceAppcache enum const
	LogLogEntrySourceAppcache LogLogEntrySource = "appcache"

	// LogLogEntrySourceRendering enum const
	LogLogEntrySourceRendering LogLogEntrySource = "rendering"

	// LogLogEntrySourceSecurity enum const
	LogLogEntrySourceSecurity LogLogEntrySource = "security"

	// LogLogEntrySourceDeprecation enum const
	LogLogEntrySourceDeprecation LogLogEntrySource = "deprecation"

	// LogLogEntrySourceWorker enum const
	LogLogEntrySourceWorker LogLogEntrySource = "worker"

	// LogLogEntrySourceViolation enum const
	LogLogEntrySourceViolation LogLogEntrySource = "violation"

	// LogLogEntrySourceIntervention enum const
	LogLogEntrySourceIntervention LogLogEntrySource = "intervention"

	// LogLogEntrySourceRecommendation enum const
	LogLogEntrySourceRecommendation LogLogEntrySource = "recommendation"

	// LogLogEntrySourceOther enum const
	LogLogEntrySourceOther LogLogEntrySource = "other"
)

// LogLogEntryLevel enum
type LogLogEntryLevel string

const (
	// LogLogEntryLevelVerbose enum const
	LogLogEntryLevelVerbose LogLogEntryLevel = "verbose"

	// LogLogEntryLevelInfo enum const
	LogLogEntryLevelInfo LogLogEntryLevel = "info"

	// LogLogEntryLevelWarning enum const
	LogLogEntryLevelWarning LogLogEntryLevel = "warning"

	// LogLogEntryLevelError enum const
	LogLogEntryLevelError LogLogEntryLevel = "error"
)

// LogLogEntryCategory enum
type LogLogEntryCategory string

const (
	// LogLogEntryCategoryCors enum const
	LogLogEntryCategoryCors LogLogEntryCategory = "cors"
)

// LogLogEntry Log entry.
type LogLogEntry struct {

	// Source Log entry source.
	Source LogLogEntrySource `json:"source"`

	// Level Log entry severity.
	Level LogLogEntryLevel `json:"level"`

	// Text Logged text.
	Text string `json:"text"`

	// Category (optional) ...
	Category LogLogEntryCategory `json:"category,omitempty"`

	// Timestamp Timestamp when this entry was added.
	Timestamp RuntimeTimestamp `json:"timestamp"`

	// URL (optional) URL of the resource if known.
	URL string `json:"url,omitempty"`

	// LineNumber (optional) Line number in the resource.
	LineNumber *int `json:"lineNumber,omitempty"`

	// StackTrace (optional) JavaScript stack trace.
	StackTrace *RuntimeStackTrace `json:"stackTrace,omitempty"`

	// NetworkRequestID (optional) Identifier of the network request associated with this entry.
	NetworkRequestID NetworkRequestID `json:"networkRequestId,omitempty"`

	// WorkerID (optional) Identifier of the worker associated with this entry.
	WorkerID string `json:"workerId,omitempty"`

	// Args (optional) Call arguments.
	Args []*RuntimeRemoteObject `json:"args,omitempty"`
}

// LogViolationSettingName enum
type LogViolationSettingName string

const (
	// LogViolationSettingNameLongTask enum const
	LogViolationSettingNameLongTask LogViolationSettingName = "longTask"

	// LogViolationSettingNameLongLayout enum const
	LogViolationSettingNameLongLayout LogViolationSettingName = "longLayout"

	// LogViolationSettingNameBlockedEvent enum const
	LogViolationSettingNameBlockedEvent LogViolationSettingName = "blockedEvent"

	// LogViolationSettingNameBlockedParser enum const
	LogViolationSettingNameBlockedParser LogViolationSettingName = "blockedParser"

	// LogViolationSettingNameDiscouragedAPIUse enum const
	LogViolationSettingNameDiscouragedAPIUse LogViolationSettingName = "discouragedAPIUse"

	// LogViolationSettingNameHandler enum const
	LogViolationSettingNameHandler LogViolationSettingName = "handler"

	// LogViolationSettingNameRecurringHandler enum const
	LogViolationSettingNameRecurringHandler LogViolationSettingName = "recurringHandler"
)

// LogViolationSetting Violation configuration setting.
type LogViolationSetting struct {

	// Name Violation type.
	Name LogViolationSettingName `json:"name"`

	// Threshold Time threshold to trigger upon.
	Threshold float64 `json:"threshold"`
}

// LogClear Clears the log.
type LogClear struct {
}

// ProtoReq name
func (m LogClear) ProtoReq() string { return "Log.clear" }

// Call sends the request
func (m LogClear) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LogDisable Disables log domain, prevents further log entries from being reported to the client.
type LogDisable struct {
}

// ProtoReq name
func (m LogDisable) ProtoReq() string { return "Log.disable" }

// Call sends the request
func (m LogDisable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LogEnable Enables log domain, sends the entries collected so far to the client by means of the
// `entryAdded` notification.
type LogEnable struct {
}

// ProtoReq name
func (m LogEnable) ProtoReq() string { return "Log.enable" }

// Call sends the request
func (m LogEnable) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LogStartViolationsReport start violation reporting.
type LogStartViolationsReport struct {

	// Config Configuration for violations.
	Config []*LogViolationSetting `json:"config"`
}

// ProtoReq name
func (m LogStartViolationsReport) ProtoReq() string { return "Log.startViolationsReport" }

// Call sends the request
func (m LogStartViolationsReport) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LogStopViolationsReport Stop violation reporting.
type LogStopViolationsReport struct {
}

// ProtoReq name
func (m LogStopViolationsReport) ProtoReq() string { return "Log.stopViolationsReport" }

// Call sends the request
func (m LogStopViolationsReport) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// LogEntryAdded Issued when new message was logged.
type LogEntryAdded struct {

	// Entry The entry.
	Entry *LogLogEntry `json:"entry"`
}

// ProtoEvent name
func (evt LogEntryAdded) ProtoEvent() string {
	return "Log.entryAdded"
}
