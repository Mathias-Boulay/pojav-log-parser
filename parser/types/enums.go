package types

const _UNKNOWN = "UNKNOWN"

// Mc version type
type McVersion string

const (
	MC_VERSION_UNKNOWN McVersion = _UNKNOWN
	VANILLA            McVersion = "VANILLA"
	OPTIFINE           McVersion = "OPTIFINE"
	FORGE              McVersion = "FORGE"
	FABRIC             McVersion = "FABRIC"
	QUILT              McVersion = "QUILT"
)

// Renderer enums
type Renderer string

const (
	RENDERER_UNKNOWN Renderer = _UNKNOWN
	GL4ES_114        Renderer = "GL4ES_114"
	GL4ES_115        Renderer = "GL4ES_115"
	HOLY_GL4ES_114   Renderer = "HOLY_GL4ES_114"
	HOLY_GL4ES_115   Renderer = "HOLY_GL4ES_115"
	ANGLE            Renderer = "ANGLE"
	VIRGL            Renderer = "VIRGL"
	ZINK             Renderer = "ZINK"
)

// Build type
type PojavBuildType string

const (
	BUILD_TYPE_UNKNOWN PojavBuildType = _UNKNOWN
	DEBUG              PojavBuildType = "DEBUG"
	PRODUCTION         PojavBuildType = "PRODUCTION"
)

// Architecture types
type DeviceArchitecture string

const (
	ARCHITECTURE_UNKNOWN DeviceArchitecture = _UNKNOWN
	ARM_32               DeviceArchitecture = "ARM_32"
	ARM_64               DeviceArchitecture = "ARM_64"
	X86                  DeviceArchitecture = "X86"
	X64                  DeviceArchitecture = "X64"
)

// Java runtime sources
type RuntimeSource string

const (
	RUNTIME_SOURCE_UNKNOWN RuntimeSource = _UNKNOWN
	INTERNAL               RuntimeSource = "INTERNAL"
	EXTERNAL               RuntimeSource = "EXTERNAL"
)

// Java runtime types
type RuntimeType string

const (
	RUNTIME_TYPE_UNKNOWN RuntimeType = _UNKNOWN
	JDK                  RuntimeType = "JDK"
	JRE                  RuntimeType = "JRE"
)
