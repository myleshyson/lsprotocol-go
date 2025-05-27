package protocol

// Compile-time interface compliance checks
var (
	_ Message         = (*InitializeRequest)(nil)
	_ Message         = (*InitializedNotification)(nil)
	_ Message         = (*InitializeResponse)(nil)
	_ Request         = (*InitializeRequest)(nil)
	_ Notification    = (*InitializedNotification)(nil)
	_ Response        = (*InitializeResponse)(nil)
	_ IncomingMessage = (*InitializeRequest)(nil)
	_ IncomingMessage = (*InitializedNotification)(nil)
)
