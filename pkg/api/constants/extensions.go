package constants

// https://github.com/opencontainers/distribution-spec/tree/main/extensions#extensions-api-for-distribution
const (
	ExtCatalogPrefix     = "/_catalog"
	ExtOciDiscoverPrefix = "/_oci/ext/discover"

	BaseExtension = "_zot"

	// zot specific extensions.
	// BasePrefix is the base prefix for zot specific extensions.
	BasePrefix = "/_zot"
	ExtPrefix  = BasePrefix + "/ext"

	// search extension.
	// ExtSearch is the search extension path.
	ExtSearch        = "/search"
	ExtSearchPrefix  = ExtPrefix + ExtSearch
	FullSearchPrefix = RoutePrefix + ExtSearchPrefix

	// mgmt extension.
	// Mgmt is the mgmt extension path.
	Mgmt     = "/mgmt"
	ExtMgmt  = ExtPrefix + Mgmt
	FullMgmt = RoutePrefix + ExtMgmt

	// signatures extension.
	// Notation is the notation signatures extension path.
	Notation     = "/notation"
	ExtNotation  = ExtPrefix + Notation
	FullNotation = RoutePrefix + ExtNotation
	Cosign       = "/cosign"
	ExtCosign    = ExtPrefix + Cosign
	FullCosign   = RoutePrefix + ExtCosign

	// user preferences extension.
	// UserPrefs is the user preferences extension path.
	UserPrefs     = "/userprefs"
	ExtUserPrefs  = ExtPrefix + UserPrefs
	FullUserPrefs = RoutePrefix + ExtUserPrefs
)
