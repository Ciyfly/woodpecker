// This file is generated by "./lib/proto/generate"

package proto

/*

Browser

The Browser domain defines methods and events for browser managing.

*/

// BrowserBrowserContextID (experimental) ...
type BrowserBrowserContextID string

// BrowserWindowID (experimental) ...
type BrowserWindowID int

// BrowserWindowState (experimental) The state of the browser window.
type BrowserWindowState string

const (
	// BrowserWindowStateNormal enum const
	BrowserWindowStateNormal BrowserWindowState = "normal"

	// BrowserWindowStateMinimized enum const
	BrowserWindowStateMinimized BrowserWindowState = "minimized"

	// BrowserWindowStateMaximized enum const
	BrowserWindowStateMaximized BrowserWindowState = "maximized"

	// BrowserWindowStateFullscreen enum const
	BrowserWindowStateFullscreen BrowserWindowState = "fullscreen"
)

// BrowserBounds (experimental) Browser window bounds information
type BrowserBounds struct {

	// Left (optional) The offset from the left edge of the screen to the window in pixels.
	Left *int `json:"left,omitempty"`

	// Top (optional) The offset from the top edge of the screen to the window in pixels.
	Top *int `json:"top,omitempty"`

	// Width (optional) The window width in pixels.
	Width *int `json:"width,omitempty"`

	// Height (optional) The window height in pixels.
	Height *int `json:"height,omitempty"`

	// WindowState (optional) The window state. Default to normal.
	WindowState BrowserWindowState `json:"windowState,omitempty"`
}

// BrowserPermissionType (experimental) ...
type BrowserPermissionType string

const (
	// BrowserPermissionTypeAccessibilityEvents enum const
	BrowserPermissionTypeAccessibilityEvents BrowserPermissionType = "accessibilityEvents"

	// BrowserPermissionTypeAudioCapture enum const
	BrowserPermissionTypeAudioCapture BrowserPermissionType = "audioCapture"

	// BrowserPermissionTypeBackgroundSync enum const
	BrowserPermissionTypeBackgroundSync BrowserPermissionType = "backgroundSync"

	// BrowserPermissionTypeBackgroundFetch enum const
	BrowserPermissionTypeBackgroundFetch BrowserPermissionType = "backgroundFetch"

	// BrowserPermissionTypeClipboardReadWrite enum const
	BrowserPermissionTypeClipboardReadWrite BrowserPermissionType = "clipboardReadWrite"

	// BrowserPermissionTypeClipboardSanitizedWrite enum const
	BrowserPermissionTypeClipboardSanitizedWrite BrowserPermissionType = "clipboardSanitizedWrite"

	// BrowserPermissionTypeDisplayCapture enum const
	BrowserPermissionTypeDisplayCapture BrowserPermissionType = "displayCapture"

	// BrowserPermissionTypeDurableStorage enum const
	BrowserPermissionTypeDurableStorage BrowserPermissionType = "durableStorage"

	// BrowserPermissionTypeFlash enum const
	BrowserPermissionTypeFlash BrowserPermissionType = "flash"

	// BrowserPermissionTypeGeolocation enum const
	BrowserPermissionTypeGeolocation BrowserPermissionType = "geolocation"

	// BrowserPermissionTypeMidi enum const
	BrowserPermissionTypeMidi BrowserPermissionType = "midi"

	// BrowserPermissionTypeMidiSysex enum const
	BrowserPermissionTypeMidiSysex BrowserPermissionType = "midiSysex"

	// BrowserPermissionTypeNfc enum const
	BrowserPermissionTypeNfc BrowserPermissionType = "nfc"

	// BrowserPermissionTypeNotifications enum const
	BrowserPermissionTypeNotifications BrowserPermissionType = "notifications"

	// BrowserPermissionTypePaymentHandler enum const
	BrowserPermissionTypePaymentHandler BrowserPermissionType = "paymentHandler"

	// BrowserPermissionTypePeriodicBackgroundSync enum const
	BrowserPermissionTypePeriodicBackgroundSync BrowserPermissionType = "periodicBackgroundSync"

	// BrowserPermissionTypeProtectedMediaIdentifier enum const
	BrowserPermissionTypeProtectedMediaIdentifier BrowserPermissionType = "protectedMediaIdentifier"

	// BrowserPermissionTypeSensors enum const
	BrowserPermissionTypeSensors BrowserPermissionType = "sensors"

	// BrowserPermissionTypeVideoCapture enum const
	BrowserPermissionTypeVideoCapture BrowserPermissionType = "videoCapture"

	// BrowserPermissionTypeVideoCapturePanTiltZoom enum const
	BrowserPermissionTypeVideoCapturePanTiltZoom BrowserPermissionType = "videoCapturePanTiltZoom"

	// BrowserPermissionTypeIdleDetection enum const
	BrowserPermissionTypeIdleDetection BrowserPermissionType = "idleDetection"

	// BrowserPermissionTypeWakeLockScreen enum const
	BrowserPermissionTypeWakeLockScreen BrowserPermissionType = "wakeLockScreen"

	// BrowserPermissionTypeWakeLockSystem enum const
	BrowserPermissionTypeWakeLockSystem BrowserPermissionType = "wakeLockSystem"
)

// BrowserPermissionSetting (experimental) ...
type BrowserPermissionSetting string

const (
	// BrowserPermissionSettingGranted enum const
	BrowserPermissionSettingGranted BrowserPermissionSetting = "granted"

	// BrowserPermissionSettingDenied enum const
	BrowserPermissionSettingDenied BrowserPermissionSetting = "denied"

	// BrowserPermissionSettingPrompt enum const
	BrowserPermissionSettingPrompt BrowserPermissionSetting = "prompt"
)

// BrowserPermissionDescriptor (experimental) Definition of PermissionDescriptor defined in the Permissions API:
// https://w3c.github.io/permissions/#dictdef-permissiondescriptor.
type BrowserPermissionDescriptor struct {

	// Name Name of permission.
	// See https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Name string `json:"name"`

	// Sysex (optional) For "midi" permission, may also specify sysex control.
	Sysex bool `json:"sysex,omitempty"`

	// UserVisibleOnly (optional) For "push" permission, may specify userVisibleOnly.
	// Note that userVisibleOnly = true is the only currently supported type.
	UserVisibleOnly bool `json:"userVisibleOnly,omitempty"`

	// AllowWithoutSanitization (optional) For "clipboard" permission, may specify allowWithoutSanitization.
	AllowWithoutSanitization bool `json:"allowWithoutSanitization,omitempty"`

	// PanTiltZoom (optional) For "camera" permission, may specify panTiltZoom.
	PanTiltZoom bool `json:"panTiltZoom,omitempty"`
}

// BrowserBrowserCommandID (experimental) Browser command ids used by executeBrowserCommand.
type BrowserBrowserCommandID string

const (
	// BrowserBrowserCommandIDOpenTabSearch enum const
	BrowserBrowserCommandIDOpenTabSearch BrowserBrowserCommandID = "openTabSearch"

	// BrowserBrowserCommandIDCloseTabSearch enum const
	BrowserBrowserCommandIDCloseTabSearch BrowserBrowserCommandID = "closeTabSearch"
)

// BrowserBucket (experimental) Chrome histogram bucket.
type BrowserBucket struct {

	// Low Minimum value (inclusive).
	Low int `json:"low"`

	// High Maximum value (exclusive).
	High int `json:"high"`

	// Count Number of samples.
	Count int `json:"count"`
}

// BrowserHistogram (experimental) Chrome histogram.
type BrowserHistogram struct {

	// Name Name.
	Name string `json:"name"`

	// Sum Sum of sample values.
	Sum int `json:"sum"`

	// Count Total number of samples.
	Count int `json:"count"`

	// Buckets Buckets.
	Buckets []*BrowserBucket `json:"buckets"`
}

// BrowserSetPermission (experimental) Set permission settings for given origin.
type BrowserSetPermission struct {

	// Permission Descriptor of permission to override.
	Permission *BrowserPermissionDescriptor `json:"permission"`

	// Setting Setting of the permission.
	Setting BrowserPermissionSetting `json:"setting"`

	// Origin (optional) Origin the permission applies to, all origins if not specified.
	Origin string `json:"origin,omitempty"`

	// BrowserContextID (optional) Context to override. When omitted, default browser context is used.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m BrowserSetPermission) ProtoReq() string { return "Browser.setPermission" }

// Call sends the request
func (m BrowserSetPermission) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserGrantPermissions (experimental) Grant specific permissions to the given origin and reject all others.
type BrowserGrantPermissions struct {

	// Permissions ...
	Permissions []BrowserPermissionType `json:"permissions"`

	// Origin (optional) Origin the permission applies to, all origins if not specified.
	Origin string `json:"origin,omitempty"`

	// BrowserContextID (optional) BrowserContext to override permissions. When omitted, default browser context is used.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m BrowserGrantPermissions) ProtoReq() string { return "Browser.grantPermissions" }

// Call sends the request
func (m BrowserGrantPermissions) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserResetPermissions (experimental) Reset all permission management for all origins.
type BrowserResetPermissions struct {

	// BrowserContextID (optional) BrowserContext to reset permissions. When omitted, default browser context is used.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m BrowserResetPermissions) ProtoReq() string { return "Browser.resetPermissions" }

// Call sends the request
func (m BrowserResetPermissions) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserSetDownloadBehaviorBehavior enum
type BrowserSetDownloadBehaviorBehavior string

const (
	// BrowserSetDownloadBehaviorBehaviorDeny enum const
	BrowserSetDownloadBehaviorBehaviorDeny BrowserSetDownloadBehaviorBehavior = "deny"

	// BrowserSetDownloadBehaviorBehaviorAllow enum const
	BrowserSetDownloadBehaviorBehaviorAllow BrowserSetDownloadBehaviorBehavior = "allow"

	// BrowserSetDownloadBehaviorBehaviorAllowAndName enum const
	BrowserSetDownloadBehaviorBehaviorAllowAndName BrowserSetDownloadBehaviorBehavior = "allowAndName"

	// BrowserSetDownloadBehaviorBehaviorDefault enum const
	BrowserSetDownloadBehaviorBehaviorDefault BrowserSetDownloadBehaviorBehavior = "default"
)

// BrowserSetDownloadBehavior (experimental) Set the behavior when downloading a file.
type BrowserSetDownloadBehavior struct {

	// Behavior Whether to allow all or deny all download requests, or use default Chrome behavior if
	// available (otherwise deny). |allowAndName| allows download and names files according to
	// their dowmload guids.
	Behavior BrowserSetDownloadBehaviorBehavior `json:"behavior"`

	// BrowserContextID (optional) BrowserContext to set download behavior. When omitted, default browser context is used.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`

	// DownloadPath (optional) The default path to save downloaded files to. This is required if behavior is set to 'allow'
	// or 'allowAndName'.
	DownloadPath string `json:"downloadPath,omitempty"`

	// EventsEnabled (optional) Whether to emit download events (defaults to false).
	EventsEnabled bool `json:"eventsEnabled,omitempty"`
}

// ProtoReq name
func (m BrowserSetDownloadBehavior) ProtoReq() string { return "Browser.setDownloadBehavior" }

// Call sends the request
func (m BrowserSetDownloadBehavior) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserCancelDownload (experimental) Cancel a download if in progress
type BrowserCancelDownload struct {

	// GUID Global unique identifier of the download.
	GUID string `json:"guid"`

	// BrowserContextID (optional) BrowserContext to perform the action in. When omitted, default browser context is used.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// ProtoReq name
func (m BrowserCancelDownload) ProtoReq() string { return "Browser.cancelDownload" }

// Call sends the request
func (m BrowserCancelDownload) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserClose Close browser gracefully.
type BrowserClose struct {
}

// ProtoReq name
func (m BrowserClose) ProtoReq() string { return "Browser.close" }

// Call sends the request
func (m BrowserClose) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserCrash (experimental) Crashes browser on the main thread.
type BrowserCrash struct {
}

// ProtoReq name
func (m BrowserCrash) ProtoReq() string { return "Browser.crash" }

// Call sends the request
func (m BrowserCrash) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserCrashGpuProcess (experimental) Crashes GPU process.
type BrowserCrashGpuProcess struct {
}

// ProtoReq name
func (m BrowserCrashGpuProcess) ProtoReq() string { return "Browser.crashGpuProcess" }

// Call sends the request
func (m BrowserCrashGpuProcess) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserGetVersion Returns version information.
type BrowserGetVersion struct {
}

// ProtoReq name
func (m BrowserGetVersion) ProtoReq() string { return "Browser.getVersion" }

// Call the request
func (m BrowserGetVersion) Call(c Client) (*BrowserGetVersionResult, error) {
	var res BrowserGetVersionResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// BrowserGetVersionResult ...
type BrowserGetVersionResult struct {

	// ProtocolVersion Protocol version.
	ProtocolVersion string `json:"protocolVersion"`

	// Product Product name.
	Product string `json:"product"`

	// Revision Product revision.
	Revision string `json:"revision"`

	// UserAgent User-Agent.
	UserAgent string `json:"userAgent"`

	// JsVersion V8 version.
	JsVersion string `json:"jsVersion"`
}

// BrowserGetBrowserCommandLine (experimental) Returns the command line switches for the browser process if, and only if
// --enable-automation is on the commandline.
type BrowserGetBrowserCommandLine struct {
}

// ProtoReq name
func (m BrowserGetBrowserCommandLine) ProtoReq() string { return "Browser.getBrowserCommandLine" }

// Call the request
func (m BrowserGetBrowserCommandLine) Call(c Client) (*BrowserGetBrowserCommandLineResult, error) {
	var res BrowserGetBrowserCommandLineResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// BrowserGetBrowserCommandLineResult (experimental) ...
type BrowserGetBrowserCommandLineResult struct {

	// Arguments Commandline parameters
	Arguments []string `json:"arguments"`
}

// BrowserGetHistograms (experimental) Get Chrome histograms.
type BrowserGetHistograms struct {

	// Query (optional) Requested substring in name. Only histograms which have query as a
	// substring in their name are extracted. An empty or absent query returns
	// all histograms.
	Query string `json:"query,omitempty"`

	// Delta (optional) If true, retrieve delta since last call.
	Delta bool `json:"delta,omitempty"`
}

// ProtoReq name
func (m BrowserGetHistograms) ProtoReq() string { return "Browser.getHistograms" }

// Call the request
func (m BrowserGetHistograms) Call(c Client) (*BrowserGetHistogramsResult, error) {
	var res BrowserGetHistogramsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// BrowserGetHistogramsResult (experimental) ...
type BrowserGetHistogramsResult struct {

	// Histograms Histograms.
	Histograms []*BrowserHistogram `json:"histograms"`
}

// BrowserGetHistogram (experimental) Get a Chrome histogram by name.
type BrowserGetHistogram struct {

	// Name Requested histogram name.
	Name string `json:"name"`

	// Delta (optional) If true, retrieve delta since last call.
	Delta bool `json:"delta,omitempty"`
}

// ProtoReq name
func (m BrowserGetHistogram) ProtoReq() string { return "Browser.getHistogram" }

// Call the request
func (m BrowserGetHistogram) Call(c Client) (*BrowserGetHistogramResult, error) {
	var res BrowserGetHistogramResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// BrowserGetHistogramResult (experimental) ...
type BrowserGetHistogramResult struct {

	// Histogram Histogram.
	Histogram *BrowserHistogram `json:"histogram"`
}

// BrowserGetWindowBounds (experimental) Get position and size of the browser window.
type BrowserGetWindowBounds struct {

	// WindowID Browser window id.
	WindowID BrowserWindowID `json:"windowId"`
}

// ProtoReq name
func (m BrowserGetWindowBounds) ProtoReq() string { return "Browser.getWindowBounds" }

// Call the request
func (m BrowserGetWindowBounds) Call(c Client) (*BrowserGetWindowBoundsResult, error) {
	var res BrowserGetWindowBoundsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// BrowserGetWindowBoundsResult (experimental) ...
type BrowserGetWindowBoundsResult struct {

	// Bounds Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds *BrowserBounds `json:"bounds"`
}

// BrowserGetWindowForTarget (experimental) Get the browser window that contains the devtools target.
type BrowserGetWindowForTarget struct {

	// TargetID (optional) Devtools agent host id. If called as a part of the session, associated targetId is used.
	TargetID TargetTargetID `json:"targetId,omitempty"`
}

// ProtoReq name
func (m BrowserGetWindowForTarget) ProtoReq() string { return "Browser.getWindowForTarget" }

// Call the request
func (m BrowserGetWindowForTarget) Call(c Client) (*BrowserGetWindowForTargetResult, error) {
	var res BrowserGetWindowForTargetResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// BrowserGetWindowForTargetResult (experimental) ...
type BrowserGetWindowForTargetResult struct {

	// WindowID Browser window id.
	WindowID BrowserWindowID `json:"windowId"`

	// Bounds Bounds information of the window. When window state is 'minimized', the restored window
	// position and size are returned.
	Bounds *BrowserBounds `json:"bounds"`
}

// BrowserSetWindowBounds (experimental) Set position and/or size of the browser window.
type BrowserSetWindowBounds struct {

	// WindowID Browser window id.
	WindowID BrowserWindowID `json:"windowId"`

	// Bounds New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined
	// with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds *BrowserBounds `json:"bounds"`
}

// ProtoReq name
func (m BrowserSetWindowBounds) ProtoReq() string { return "Browser.setWindowBounds" }

// Call sends the request
func (m BrowserSetWindowBounds) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserSetDockTile (experimental) Set dock tile details, platform-specific.
type BrowserSetDockTile struct {

	// BadgeLabel (optional) ...
	BadgeLabel string `json:"badgeLabel,omitempty"`

	// Image (optional) Png encoded image.
	Image []byte `json:"image,omitempty"`
}

// ProtoReq name
func (m BrowserSetDockTile) ProtoReq() string { return "Browser.setDockTile" }

// Call sends the request
func (m BrowserSetDockTile) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserExecuteBrowserCommand (experimental) Invoke custom browser commands used by telemetry.
type BrowserExecuteBrowserCommand struct {

	// CommandID ...
	CommandID BrowserBrowserCommandID `json:"commandId"`
}

// ProtoReq name
func (m BrowserExecuteBrowserCommand) ProtoReq() string { return "Browser.executeBrowserCommand" }

// Call sends the request
func (m BrowserExecuteBrowserCommand) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// BrowserDownloadWillBegin (experimental) Fired when page is about to start a download.
type BrowserDownloadWillBegin struct {

	// FrameID Id of the frame that caused the download to begin.
	FrameID PageFrameID `json:"frameId"`

	// GUID Global unique identifier of the download.
	GUID string `json:"guid"`

	// URL URL of the resource being downloaded.
	URL string `json:"url"`

	// SuggestedFilename Suggested file name of the resource (the actual name of the file saved on disk may differ).
	SuggestedFilename string `json:"suggestedFilename"`
}

// ProtoEvent name
func (evt BrowserDownloadWillBegin) ProtoEvent() string {
	return "Browser.downloadWillBegin"
}

// BrowserDownloadProgressState enum
type BrowserDownloadProgressState string

const (
	// BrowserDownloadProgressStateInProgress enum const
	BrowserDownloadProgressStateInProgress BrowserDownloadProgressState = "inProgress"

	// BrowserDownloadProgressStateCompleted enum const
	BrowserDownloadProgressStateCompleted BrowserDownloadProgressState = "completed"

	// BrowserDownloadProgressStateCanceled enum const
	BrowserDownloadProgressStateCanceled BrowserDownloadProgressState = "canceled"
)

// BrowserDownloadProgress (experimental) Fired when download makes progress. Last call has |done| == true.
type BrowserDownloadProgress struct {

	// GUID Global unique identifier of the download.
	GUID string `json:"guid"`

	// TotalBytes Total expected bytes to download.
	TotalBytes float64 `json:"totalBytes"`

	// ReceivedBytes Total bytes received.
	ReceivedBytes float64 `json:"receivedBytes"`

	// State Download status.
	State BrowserDownloadProgressState `json:"state"`
}

// ProtoEvent name
func (evt BrowserDownloadProgress) ProtoEvent() string {
	return "Browser.downloadProgress"
}
