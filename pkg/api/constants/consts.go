package constants

import "time"

const (
	RoutePrefix                  = "/v2"
	Blobs                        = "blobs"
	Uploads                      = "uploads"
	DistAPIVersion               = "Docker-Distribution-API-Version"
	DistContentDigestKey         = "Docker-Content-Digest"
	SubjectDigestKey             = "OCI-Subject"
	BlobUploadUUID               = "Blob-Upload-UUID"
	DefaultMediaType             = "application/json"
	BinaryMediaType              = "application/octet-stream"
	DefaultMetricsExtensionRoute = "/metrics"
	AppNamespacePath             = "/zot"
	CallbackBasePath             = AppNamespacePath + "/auth/callback"
	LoginPath                    = AppNamespacePath + "/auth/login"
	LogoutPath                   = AppNamespacePath + "/auth/logout"
	APIKeyPath                   = AppNamespacePath + "/auth/apikey"
	SessionClientHeaderName      = "X-ZOT-API-CLIENT"
	SessionClientHeaderValue     = "zot-ui"
	APIKeysPrefix                = "zak_"
	CallbackUIQueryParam         = "callback_ui"
	APIKeyTimeFormat             = time.RFC3339
	// authz permissions.
	// method actions.
	// CreatePermission is an authz permission for method actions.
	CreatePermission = "create"
	ReadPermission   = "read"
	UpdatePermission = "update"
	DeletePermission = "delete"
	// behaviour actions.
	// DetectManifestCollisionPermission is an authz permission for behaviour actions.
	DetectManifestCollisionPermission = "detectManifestCollision"
	// zot scale-out hop count header.
	// ScaleOutHopCountHeader is the zot scale-out hop count header.
	ScaleOutHopCountHeader = "X-Zot-Cluster-Hop-Count"
	// log string keys.
	// these can be used together with the logger to add context to a log message.
	// RepositoryLogKey is a log string key that can be used together with the logger to add context to a log message.
	RepositoryLogKey = "repository"
)
