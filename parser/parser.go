package parser

import (
	"fmt"
	"strconv"
	"strings"

	regexp "github.com/wasilibs/go-re2"

	"github.com/mathias-boulay/parser/types"
)

/** Build a mc version from */
func ParseVersion(log string) types.PojavLauncherVersion {
	pattern := regexp.MustCompile(`Info: Launcher version: (.*?)-(.*?)-(.*?)-(\S*)`)
	results := pattern.FindStringSubmatch(log)
	return types.PojavLauncherVersion{MajorCode: results[1], CommitNumber: results[2], CommitSha: results[3], Branch: results[4]}
}

/** Build the minecraft parse version */
func ParseMinecraftVersion(log string) types.MinecraftVersion {
	versionNamePattern := regexp.MustCompile(`Selected Minecraft version: (.*)`)
	analyzedVersion := types.MinecraftVersion{Name: "UNKNOWN", Type: types.VANILLA}

	if regexVersionName := versionNamePattern.FindStringSubmatch(log); len(regexVersionName) > 0 {
		analyzedVersion.Name = strings.ToLower(regexVersionName[1])
	}

	types := map[string]types.McVersion{"forge": types.FORGE, "optifine": types.OPTIFINE, "fabric": types.FABRIC, "quilt": types.QUILT}
	for key, mcType := range types {
		if strings.Contains(log, key) {
			analyzedVersion.Type = mcType
			break
		}
	}

	return analyzedVersion
}

/** Parse the mc architecture */
func ParseArchitecture(log string) types.DeviceArchitecture {
	pattern := regexp.MustCompile(`(?m)Architecture: (.*)$`)
	result := pattern.FindStringSubmatch(log)[1]
	result = strings.TrimSpace(result)

	architectures := map[string]types.DeviceArchitecture{"arm64": types.ARM_64, "arm": types.ARM_32, "x86_64": types.X64, "x86": types.X86}
	for key, arch := range architectures {
		if strings.EqualFold(result, key) {
			return arch

		}
	}

	return types.ARCHITECTURE_UNKNOWN
}

/** Get all the env variables */
func ParseEnvVariables(log string) map[string]string {
	pattern := regexp.MustCompile(`(?m)Added custom env: (.*?)=(.*)$`)
	results := pattern.FindAllStringSubmatch(log, -1)
	envDictionary := make(map[string]string)

	for _, envTuple := range results {
		envDictionary[envTuple[1]] = envTuple[2]
	}

	return envDictionary
}

/** Get all the java args */
func ParseJavaArguments(log string) string {
	pattern := regexp.MustCompile(`(?m)Info: Custom Java arguments:.*?"(.*)"`)
	results := pattern.FindStringSubmatch(log)
	return results[1]
}

/** Parse the java runtime */
func ParseJavaRuntime(javaPath string) types.JavaRuntime {
	javaPath = strings.ToLower(javaPath)
	if strings.HasSuffix(javaPath, "/internal") {
		return types.JavaRuntime{Version: 8, Source: types.INTERNAL, Type: types.JRE}
	}
	if strings.HasSuffix(javaPath, "/internal17") || strings.HasSuffix(javaPath, "/internal-17") {
		return types.JavaRuntime{Version: 17, Source: types.INTERNAL, Type: types.JRE}
	}

	var runtime types.JavaRuntime
	runtime.Source = types.EXTERNAL
	if strings.Contains(javaPath, "jdk") {
		runtime.Type = types.JDK
	}
	if strings.Contains(javaPath, "jre") {
		runtime.Type = types.JRE
	}

	pattern := regexp.MustCompile(`(?:jdk|jre)(\d*)`)
	results := pattern.FindStringSubmatch(javaPath)
	if len(results) > 1 {
		version := results[1]
		if parsedVersion, err := strconv.ParseFloat(version, 32); err == nil {
			runtime.Version = float32(parsedVersion)
		} else {
			fmt.Println("Error during java runtime parsing")
		}
	}

	return runtime
}

/** Parse the renderer */
func ParseRenderer(log string) types.Renderer {
	pattern := regexp.MustCompile(`v1\.1\.(\d) built on `)
	result := pattern.FindStringSubmatch(log)
	if len(result) > 1 {
		gl4esVersion, _ := strconv.Atoi(result[1])
		isHoly := strings.Contains(log, "GLSL 320 es supported")
		if gl4esVersion == 5 {
			if isHoly {
				return types.HOLY_GL4ES_115
			} else {
				return types.GL4ES_115
			}
		}
		if gl4esVersion == 4 {
			if isHoly {
				return types.HOLY_GL4ES_114
			} else {
				return types.GL4ES_114
			}
		}
	}

	// TODO support zink, and perhaps vulkanmod as a special case ?
	m := map[string]types.Renderer{"VirGL: libvirgl_test_server": types.VIRGL, "POJAV_RENDERER=opengles3_desktopgl_angle_vulkan": types.ANGLE}
	for key, renderer := range m {
		if strings.Contains(log, key) {
			return renderer
		}
	}

	return types.RENDERER_UNKNOWN
}

/** Gather the build type from the package name */
func ParseBuildType(log string) types.PojavBuildType {
	m := map[string]types.PojavBuildType{"net.kdt.pojavlaunch.debug": types.DEBUG, "net.kdt.pojavlaunch.": types.BUILD_TYPE_UNKNOWN, "net.kdt.pojavlaunch": types.PRODUCTION}
	for key, buildType := range m {
		if strings.Contains(log, key) {
			return buildType
		}
	}

	return types.BUILD_TYPE_UNKNOWN
}

/** Get the mc username */
func ParseMinecraftUsername(log string) string {
	usernamePattern := regexp.MustCompile(`user: (.*)`)
	username := usernamePattern.FindStringSubmatch(log)
	if len(username) > 0 {
		return username[1]
	}
	return ""
}

/** Get all forge mods */
func ParseForgeMods(log string) [][]string {
	modPatterns := regexp.MustCompile(`^\t\| ([ULCHIJADE].*?) {1}\| (.*?) *\|`).FindAllStringSubmatch(log, -1)

	return modPatterns
}

func ParseLog(log string, shouldLocalize bool) types.ParsedLog {
	parsedLog := types.ParsedLog{
		Version:          ParseVersion(log),
		MinecraftVersion: ParseMinecraftVersion(log),
		Renderer:         ParseRenderer(log),
		EnvVariables:     ParseEnvVariables(log),
		Architecture:     ParseArchitecture(log),
		JavaArguments:    ParseJavaArguments(log),
		Username:         ParseMinecraftUsername(log),
	}
	parsedLog.BuildType = ParseBuildType(parsedLog.EnvVariables["JAVA_HOME"])
	parsedLog.JavaRuntime = ParseJavaRuntime(parsedLog.EnvVariables["JAVA_HOME"])

	parsedLog.Errors = CheckErrors(log, parsedLog)
	if shouldLocalize {
		Localize(parsedLog.Errors)
	}

	return parsedLog
}
