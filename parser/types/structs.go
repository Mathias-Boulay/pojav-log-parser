package types

/** The response struct */
type ParsedLog struct {
	Version          PojavLauncherVersion `json:"version"`
	MinecraftVersion MinecraftVersion     `json:"minecraft_version"`
	Renderer         Renderer             `json:"renderer"`
	EnvVariables     map[string]string    `json:"env_variables"`
	BuildType        PojavBuildType       `json:"build_type"`
	Architecture     DeviceArchitecture   `json:"architecture"`
	JavaArguments    string               `json:"java_arguments"`
	Username         string               `json:"username"`
	// Optional properties
	JavaRuntime JavaRuntime `json:"java_runtime"`
	Errors      []string    `json:"errors"`
}

/** The pojav version details */
type PojavLauncherVersion struct {
	MajorCode    string `json:"major_code"`
	CommitNumber string `json:"commit_number"`
	CommitSha    string `json:"commit_sha"`
	Branch       string `json:"branch"`
}

type MinecraftVersion struct {
	Name string
	Type McVersion
}

type JavaRuntime struct {
	Version float32
	Source  RuntimeSource
	Type    RuntimeType
}
