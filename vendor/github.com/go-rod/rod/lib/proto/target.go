// This file is generated by "./lib/proto/generate"

package proto

/*

Target

Supports additional targets discovery and allows to attach to them.

*/

// TargetTargetID ...
type TargetTargetID string

// TargetSessionID Unique identifier of attached debugging session.
type TargetSessionID string

// TargetTargetInfoType enum
type TargetTargetInfoType string

const (
	// TargetTargetInfoTypePage enum const
	TargetTargetInfoTypePage TargetTargetInfoType = "page"

	// TargetTargetInfoTypeBackgroundPage enum const
	TargetTargetInfoTypeBackgroundPage TargetTargetInfoType = "background_page"

	// TargetTargetInfoTypeServiceWorker enum const
	TargetTargetInfoTypeServiceWorker TargetTargetInfoType = "service_worker"

	// TargetTargetInfoTypeSharedWorker enum const
	TargetTargetInfoTypeSharedWorker TargetTargetInfoType = "shared_worker"

	// TargetTargetInfoTypeBrowser enum const
	TargetTargetInfoTypeBrowser TargetTargetInfoType = "browser"

	// TargetTargetInfoTypeOther enum const
	TargetTargetInfoTypeOther TargetTargetInfoType = "other"
)

// TargetTargetInfo ...
type TargetTargetInfo struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`

	// Type ...
	Type TargetTargetInfoType `json:"type"`

	// Title ...
	Title string `json:"title"`

	// URL ...
	URL string `json:"url"`

	// Attached Whether the target has an attached client.
	Attached bool `json:"attached"`

	// OpenerID (optional) Opener target Id
	OpenerID TargetTargetID `json:"openerId,omitempty"`

	// CanAccessOpener (experimental) Whether the target has access to the originating window.
	CanAccessOpener bool `json:"canAccessOpener"`

	// OpenerFrameID (experimental) (optional) Frame id of originating window (is only set if target has an opener).
	OpenerFrameID PageFrameID `json:"openerFrameId,omitempty"`

	// BrowserContextID (experimental) (optional) ...
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`
}

// TargetRemoteLocation (experimental) ...
type TargetRemoteLocation struct {

	// Host ...
	Host string `json:"host"`

	// Port ...
	Port int `json:"port"`
}

// TargetActivateTarget Activates (focuses) the target.
type TargetActivateTarget struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`
}

// ProtoReq name
func (m TargetActivateTarget) ProtoReq() string { return "Target.activateTarget" }

// Call sends the request
func (m TargetActivateTarget) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetAttachToTarget Attaches to the target with given id.
type TargetAttachToTarget struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`

	// Flatten (optional) Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// ProtoReq name
func (m TargetAttachToTarget) ProtoReq() string { return "Target.attachToTarget" }

// Call the request
func (m TargetAttachToTarget) Call(c Client) (*TargetAttachToTargetResult, error) {
	var res TargetAttachToTargetResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetAttachToTargetResult ...
type TargetAttachToTargetResult struct {

	// SessionID Id assigned to the session.
	SessionID TargetSessionID `json:"sessionId"`
}

// TargetAttachToBrowserTarget (experimental) Attaches to the browser target, only uses flat sessionId mode.
type TargetAttachToBrowserTarget struct {
}

// ProtoReq name
func (m TargetAttachToBrowserTarget) ProtoReq() string { return "Target.attachToBrowserTarget" }

// Call the request
func (m TargetAttachToBrowserTarget) Call(c Client) (*TargetAttachToBrowserTargetResult, error) {
	var res TargetAttachToBrowserTargetResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetAttachToBrowserTargetResult (experimental) ...
type TargetAttachToBrowserTargetResult struct {

	// SessionID Id assigned to the session.
	SessionID TargetSessionID `json:"sessionId"`
}

// TargetCloseTarget Closes the target. If the target is a page that gets closed too.
type TargetCloseTarget struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`
}

// ProtoReq name
func (m TargetCloseTarget) ProtoReq() string { return "Target.closeTarget" }

// Call the request
func (m TargetCloseTarget) Call(c Client) (*TargetCloseTargetResult, error) {
	var res TargetCloseTargetResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetCloseTargetResult ...
type TargetCloseTargetResult struct {

	// Success (deprecated) Always set to true. If an error occurs, the response indicates protocol error.
	Success bool `json:"success"`
}

// TargetExposeDevToolsProtocol (experimental) Inject object to the target's main frame that provides a communication
// channel with browser target.
//
// Injected object will be available as `window[bindingName]`.
//
// The object has the follwing API:
// - `binding.send(json)` - a method to send messages over the remote debugging protocol
// - `binding.onmessage = json => handleMessage(json)` - a callback that will be called for the protocol notifications and command responses.
type TargetExposeDevToolsProtocol struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`

	// BindingName (optional) Binding name, 'cdp' if not specified.
	BindingName string `json:"bindingName,omitempty"`
}

// ProtoReq name
func (m TargetExposeDevToolsProtocol) ProtoReq() string { return "Target.exposeDevToolsProtocol" }

// Call sends the request
func (m TargetExposeDevToolsProtocol) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetCreateBrowserContext (experimental) Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than
// one.
type TargetCreateBrowserContext struct {

	// DisposeOnDetach (optional) If specified, disposes this context when debugging session disconnects.
	DisposeOnDetach bool `json:"disposeOnDetach,omitempty"`

	// ProxyServer (optional) Proxy server, similar to the one passed to --proxy-server
	ProxyServer string `json:"proxyServer,omitempty"`

	// ProxyBypassList (optional) Proxy bypass list, similar to the one passed to --proxy-bypass-list
	ProxyBypassList string `json:"proxyBypassList,omitempty"`

	// OriginsWithUniversalNetworkAccess (optional) An optional list of origins to grant unlimited cross-origin access to.
	// Parts of the URL other than those constituting origin are ignored.
	OriginsWithUniversalNetworkAccess []string `json:"originsWithUniversalNetworkAccess,omitempty"`
}

// ProtoReq name
func (m TargetCreateBrowserContext) ProtoReq() string { return "Target.createBrowserContext" }

// Call the request
func (m TargetCreateBrowserContext) Call(c Client) (*TargetCreateBrowserContextResult, error) {
	var res TargetCreateBrowserContextResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetCreateBrowserContextResult (experimental) ...
type TargetCreateBrowserContextResult struct {

	// BrowserContextID The id of the context created.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId"`
}

// TargetGetBrowserContexts (experimental) Returns all browser contexts created with `Target.createBrowserContext` method.
type TargetGetBrowserContexts struct {
}

// ProtoReq name
func (m TargetGetBrowserContexts) ProtoReq() string { return "Target.getBrowserContexts" }

// Call the request
func (m TargetGetBrowserContexts) Call(c Client) (*TargetGetBrowserContextsResult, error) {
	var res TargetGetBrowserContextsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetGetBrowserContextsResult (experimental) ...
type TargetGetBrowserContextsResult struct {

	// BrowserContextIds An array of browser context ids.
	BrowserContextIds []BrowserBrowserContextID `json:"browserContextIds"`
}

// TargetCreateTarget Creates a new page.
type TargetCreateTarget struct {

	// URL The initial URL the page will be navigated to. An empty string indicates about:blank.
	URL string `json:"url"`

	// Width (optional) Frame width in DIP (headless chrome only).
	Width *int `json:"width,omitempty"`

	// Height (optional) Frame height in DIP (headless chrome only).
	Height *int `json:"height,omitempty"`

	// BrowserContextID (experimental) (optional) The browser context to create the page in.
	BrowserContextID BrowserBrowserContextID `json:"browserContextId,omitempty"`

	// EnableBeginFrameControl (experimental) (optional) Whether BeginFrames for this target will be controlled via DevTools (headless chrome only,
	// not supported on MacOS yet, false by default).
	EnableBeginFrameControl bool `json:"enableBeginFrameControl,omitempty"`

	// NewWindow (optional) Whether to create a new Window or Tab (chrome-only, false by default).
	NewWindow bool `json:"newWindow,omitempty"`

	// Background (optional) Whether to create the target in background or foreground (chrome-only,
	// false by default).
	Background bool `json:"background,omitempty"`
}

// ProtoReq name
func (m TargetCreateTarget) ProtoReq() string { return "Target.createTarget" }

// Call the request
func (m TargetCreateTarget) Call(c Client) (*TargetCreateTargetResult, error) {
	var res TargetCreateTargetResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetCreateTargetResult ...
type TargetCreateTargetResult struct {

	// TargetID The id of the page opened.
	TargetID TargetTargetID `json:"targetId"`
}

// TargetDetachFromTarget Detaches session with given id.
type TargetDetachFromTarget struct {

	// SessionID (optional) Session to detach.
	SessionID TargetSessionID `json:"sessionId,omitempty"`

	// TargetID (deprecated) (optional) Deprecated.
	TargetID TargetTargetID `json:"targetId,omitempty"`
}

// ProtoReq name
func (m TargetDetachFromTarget) ProtoReq() string { return "Target.detachFromTarget" }

// Call sends the request
func (m TargetDetachFromTarget) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetDisposeBrowserContext (experimental) Deletes a BrowserContext. All the belonging pages will be closed without calling their
// beforeunload hooks.
type TargetDisposeBrowserContext struct {

	// BrowserContextID ...
	BrowserContextID BrowserBrowserContextID `json:"browserContextId"`
}

// ProtoReq name
func (m TargetDisposeBrowserContext) ProtoReq() string { return "Target.disposeBrowserContext" }

// Call sends the request
func (m TargetDisposeBrowserContext) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetGetTargetInfo (experimental) Returns information about a target.
type TargetGetTargetInfo struct {

	// TargetID (optional) ...
	TargetID TargetTargetID `json:"targetId,omitempty"`
}

// ProtoReq name
func (m TargetGetTargetInfo) ProtoReq() string { return "Target.getTargetInfo" }

// Call the request
func (m TargetGetTargetInfo) Call(c Client) (*TargetGetTargetInfoResult, error) {
	var res TargetGetTargetInfoResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetGetTargetInfoResult (experimental) ...
type TargetGetTargetInfoResult struct {

	// TargetInfo ...
	TargetInfo *TargetTargetInfo `json:"targetInfo"`
}

// TargetGetTargets Retrieves a list of available targets.
type TargetGetTargets struct {
}

// ProtoReq name
func (m TargetGetTargets) ProtoReq() string { return "Target.getTargets" }

// Call the request
func (m TargetGetTargets) Call(c Client) (*TargetGetTargetsResult, error) {
	var res TargetGetTargetsResult
	return &res, call(m.ProtoReq(), m, &res, c)
}

// TargetGetTargetsResult ...
type TargetGetTargetsResult struct {

	// TargetInfos The list of targets.
	TargetInfos []*TargetTargetInfo `json:"targetInfos"`
}

// TargetSendMessageToTarget (deprecated) Sends protocol message over session with given id.
// Consider using flat mode instead; see commands attachToTarget, setAutoAttach,
// and crbug.com/991325.
type TargetSendMessageToTarget struct {

	// Message ...
	Message string `json:"message"`

	// SessionID (optional) Identifier of the session.
	SessionID TargetSessionID `json:"sessionId,omitempty"`

	// TargetID (deprecated) (optional) Deprecated.
	TargetID TargetTargetID `json:"targetId,omitempty"`
}

// ProtoReq name
func (m TargetSendMessageToTarget) ProtoReq() string { return "Target.sendMessageToTarget" }

// Call sends the request
func (m TargetSendMessageToTarget) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetSetAutoAttach (experimental) Controls whether to automatically attach to new targets which are considered to be related to
// this one. When turned on, attaches to all existing related targets as well. When turned off,
// automatically detaches from all currently attached targets.
// This also clears all targets added by `autoAttachRelated` from the list of targets to watch
// for creation of related targets.
type TargetSetAutoAttach struct {

	// AutoAttach Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`

	// WaitForDebuggerOnStart Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`

	// Flatten (optional) Enables "flat" access to the session via specifying sessionId attribute in the commands.
	// We plan to make this the default, deprecate non-flattened mode,
	// and eventually retire it. See crbug.com/991325.
	Flatten bool `json:"flatten,omitempty"`
}

// ProtoReq name
func (m TargetSetAutoAttach) ProtoReq() string { return "Target.setAutoAttach" }

// Call sends the request
func (m TargetSetAutoAttach) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetAutoAttachRelated (experimental) Adds the specified target to the list of targets that will be monitored for any related target
// creation (such as child frames, child workers and new versions of service worker) and reported
// through `attachedToTarget`. The specified target is also auto-attached.
// This cancels the effect of any previous `setAutoAttach` and is also cancelled by subsequent
// `setAutoAttach`. Only available at the Browser target.
type TargetAutoAttachRelated struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`

	// WaitForDebuggerOnStart Whether to pause new targets when attaching to them. Use `Runtime.runIfWaitingForDebugger`
	// to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

// ProtoReq name
func (m TargetAutoAttachRelated) ProtoReq() string { return "Target.autoAttachRelated" }

// Call sends the request
func (m TargetAutoAttachRelated) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetSetDiscoverTargets Controls whether to discover available targets and notify via
// `targetCreated/targetInfoChanged/targetDestroyed` events.
type TargetSetDiscoverTargets struct {

	// Discover Whether to discover available targets.
	Discover bool `json:"discover"`
}

// ProtoReq name
func (m TargetSetDiscoverTargets) ProtoReq() string { return "Target.setDiscoverTargets" }

// Call sends the request
func (m TargetSetDiscoverTargets) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetSetRemoteLocations (experimental) Enables target discovery for the specified locations, when `setDiscoverTargets` was set to
// `true`.
type TargetSetRemoteLocations struct {

	// Locations List of remote locations.
	Locations []*TargetRemoteLocation `json:"locations"`
}

// ProtoReq name
func (m TargetSetRemoteLocations) ProtoReq() string { return "Target.setRemoteLocations" }

// Call sends the request
func (m TargetSetRemoteLocations) Call(c Client) error {
	return call(m.ProtoReq(), m, nil, c)
}

// TargetAttachedToTarget (experimental) Issued when attached to target because of auto-attach or `attachToTarget` command.
type TargetAttachedToTarget struct {

	// SessionID Identifier assigned to the session used to send/receive messages.
	SessionID TargetSessionID `json:"sessionId"`

	// TargetInfo ...
	TargetInfo *TargetTargetInfo `json:"targetInfo"`

	// WaitingForDebugger ...
	WaitingForDebugger bool `json:"waitingForDebugger"`
}

// ProtoEvent name
func (evt TargetAttachedToTarget) ProtoEvent() string {
	return "Target.attachedToTarget"
}

// TargetDetachedFromTarget (experimental) Issued when detached from target for any reason (including `detachFromTarget` command). Can be
// issued multiple times per target if multiple sessions have been attached to it.
type TargetDetachedFromTarget struct {

	// SessionID Detached session identifier.
	SessionID TargetSessionID `json:"sessionId"`

	// TargetID (deprecated) (optional) Deprecated.
	TargetID TargetTargetID `json:"targetId,omitempty"`
}

// ProtoEvent name
func (evt TargetDetachedFromTarget) ProtoEvent() string {
	return "Target.detachedFromTarget"
}

// TargetReceivedMessageFromTarget Notifies about a new protocol message received from the session (as reported in
// `attachedToTarget` event).
type TargetReceivedMessageFromTarget struct {

	// SessionID Identifier of a session which sends a message.
	SessionID TargetSessionID `json:"sessionId"`

	// Message ...
	Message string `json:"message"`

	// TargetID (deprecated) (optional) Deprecated.
	TargetID TargetTargetID `json:"targetId,omitempty"`
}

// ProtoEvent name
func (evt TargetReceivedMessageFromTarget) ProtoEvent() string {
	return "Target.receivedMessageFromTarget"
}

// TargetTargetCreated Issued when a possible inspection target is created.
type TargetTargetCreated struct {

	// TargetInfo ...
	TargetInfo *TargetTargetInfo `json:"targetInfo"`
}

// ProtoEvent name
func (evt TargetTargetCreated) ProtoEvent() string {
	return "Target.targetCreated"
}

// TargetTargetDestroyed Issued when a target is destroyed.
type TargetTargetDestroyed struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`
}

// ProtoEvent name
func (evt TargetTargetDestroyed) ProtoEvent() string {
	return "Target.targetDestroyed"
}

// TargetTargetCrashed Issued when a target has crashed.
type TargetTargetCrashed struct {

	// TargetID ...
	TargetID TargetTargetID `json:"targetId"`

	// Status Termination status type.
	Status string `json:"status"`

	// ErrorCode Termination error code.
	ErrorCode int `json:"errorCode"`
}

// ProtoEvent name
func (evt TargetTargetCrashed) ProtoEvent() string {
	return "Target.targetCrashed"
}

// TargetTargetInfoChanged Issued when some information about a target has changed. This only happens between
// `targetCreated` and `targetDestroyed`.
type TargetTargetInfoChanged struct {

	// TargetInfo ...
	TargetInfo *TargetTargetInfo `json:"targetInfo"`
}

// ProtoEvent name
func (evt TargetTargetInfoChanged) ProtoEvent() string {
	return "Target.targetInfoChanged"
}
