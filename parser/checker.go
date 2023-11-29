package parser

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	regexp "github.com/wasilibs/go-re2"

	"github.com/mathias-boulay/parser/types"
)

/* Generic version of a check function */
func checkGeneric(log string) []string {
	errors := make([]string, 0)
	for _, arrayElements := range types.GenericChecks() {
		for elementIndex, elementToCheck := range arrayElements {
			if elementIndex == 0 {
				// TODO skip the first round, as it is the error code
				continue
			}

			if strings.Contains(log, elementToCheck) {
				errors = append(errors, arrayElements[0])
				break
			}
		}
	}
	return errors
}

/*
Reimplementation of the java argument parser
Warns if something was left unparsed or something
*/
func checkJavaArguments(javaArgs string) bool {
	separators := []string{"-XX:-", "-XX:+", "-XX:", "--", "-D", "-X", "-javaagent:", "-verbose"}

	for _, separator := range separators {
		for strings.Contains(javaArgs, separator) {
			startIndex := strings.Index(javaArgs, separator)
			endIndex := len(javaArgs)

			for _, endSeparator := range separators {
				if tempEnd := strings.Index(javaArgs[startIndex+len(separator):], endSeparator); tempEnd != -1 {
					if endSeparatorIndex := tempEnd + startIndex + len(separator); endSeparatorIndex < endIndex {
						endIndex = endSeparatorIndex
					}
				}
			}

			extractedArgument := javaArgs[startIndex:endIndex]
			javaArgs = strings.ReplaceAll(javaArgs, extractedArgument, "")

			if strings.Count(extractedArgument, "=") > 1 {
				return false
			}
		}
	}

	return len(javaArgs) == 0
}

/* Check if the branch is the main one */
func checkBranch(version types.PojavLauncherVersion) bool {
	return version.Branch == "v3_openjdk"
}

/* Check if the release name is the latest, to be updated manually */
func checkRelease(version types.PojavLauncherVersion) bool {
	return version.MajorCode == "edelweiss"
}

/* Check the env variables */
func checkEnvVariables(env map[string]string) []string {
	var errors []string
	if mipmap, _ := strconv.Atoi(env["LIBGL_MIPMAP"]); mipmap != 3 {
		errors = append(errors, "mipmap.enabled")
	}

	if gl, _ := strconv.Atoi(env["LIBGL_GL"]); gl != 0 {
		errors = append(errors, "libgl.custom.value")
	}
	return errors
}

/* Check if we have the default java runtime */
func checkRuntime(runtime types.JavaRuntime) []string {
	var errors []string
	if runtime.Source != types.INTERNAL {
		errors = append(errors, "java.runtime.wrong.source")
	}
	if runtime.Type != types.JRE {
		errors = append(errors, "java.runtime.wrong.type")
	}
	return errors
}

/* Check common errors */
func checkForgeModErrors(log string, version types.MinecraftVersion) []string {
	errors := make([]string, 0)
	if version.Type != types.FORGE {
		return errors
	}
	for _, mod := range ParseForgeMods(log) {
		if strings.Contains(mod[1], "E") { // Meaning the mod has crashed
			errors = append(errors, "Error with mod: "+mod[2])
		}
	}
	return errors
}

/* Check whether the account is offline, we never know */
func checkOfflineAccount(log, minecraftUsername string) []string {
	errors := make([]string, 0)
	if minecraftUsername == "" {
		return errors
	}
	usernamePattern := regexp.MustCompile(`(?i)[a-zA-Z0-9_]{3,}`)
	if !usernamePattern.MatchString(strings.TrimSpace(minecraftUsername)) {
		errors = append(errors, "account.offline")
		return errors
	}
	resp, err := http.Get("https://api.ashcon.app/mojang/v2/user/" + minecraftUsername)
	if err != nil {
		fmt.Printf("Unable to request minetools with the username: %s\n", minecraftUsername)
		fmt.Println(err)
		return errors
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		errors = append(errors, "account.offline")
	}
	return errors
}

/* Aggregate all the issues */
func CheckErrors(logs string, parsedLogs types.ParsedLog) []string {
	errors := make([]string, 0)

	errors = append(errors, checkEnvVariables(parsedLogs.EnvVariables)...)
	if !checkJavaArguments(parsedLogs.JavaArguments) {
		errors = append(errors, "java.args.unparsed")
	}
	if !checkBranch(parsedLogs.Version) {
		errors = append(errors, "build.not.mainline")
	}
	if !checkRelease(parsedLogs.Version) {
		errors = append(errors, "build.not.release")
	}

	// Fork of pojav tend to have custom errors, warn about those
	if parsedLogs.BuildType == types.BUILD_TYPE_UNKNOWN {
		errors = append(errors, "build.unknown")
	}

	// Virgl + 32 bits combo check
	errors = append(errors, checkRuntime(parsedLogs.JavaRuntime)...)
	if parsedLogs.Renderer == types.VIRGL && (parsedLogs.Architecture == types.ARM_32 || parsedLogs.Architecture == types.X86) {
		errors = append(errors, "renderer.unsupported.arch")
	}

	errors = append(errors, checkForgeModErrors(logs, parsedLogs.MinecraftVersion)...)
	errors = append(errors, checkOfflineAccount(logs, parsedLogs.Username)...)
	// Generic errors to look forward
	errors = append(errors, checkGeneric(logs)...)
	// TODO more verification

	return errors
}
