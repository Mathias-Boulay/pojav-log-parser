package parser

/* Modify in place the errors by a translation */
func Localize(errors []string) {
	translations := LocalizeEnglish()
	for index := range errors {
		errors[index] = translations[errors[index]]
	}
}

/* The english translation map */
func LocalizeEnglish() map[string]string {
	return map[string]string{
		"mipmap.enabled":            "Mipmap is not disabled, this may destroy performance",
		"libgl.custom.value":        "Custom openGL declaration !",
		"java.args.unparsed":        "Some java arguments were not parsed",
		"java.args.unrecognised":    "Java arguments were not recognised",
		"build.not.mainline":        "The branch is not mainline",
		"build.unknown":             "Unknown build type, no support ?",
		"java.runtime.wrong.source": "Unexpected source for the java runtime !",
		"java.runtime.wrong.type":   "The runtime is not a JRE !",
		"java.runtime.oom":          "The game ran out of memory",
		"java.runtime.oom.start":    "The JVM couldn't allocated the desired memory amount at start",
		"renderer.unsupported.arch": "Unsupported renderer on 32 bits !",
		"fabric.incompatible.mods":  "Some fabric mods are incompatible",
		"load.shader.fail":          "Some shaders failed to be loaded",
		"runtime.early.fail":        "The program failed early",
		"optifine.render_region.on": "Optifine render region setting seems to be on",
		"account.offline":           "The account is offline or doesn't exist",
		"opengl.generic.error":      "An OpenGL error occured",
		"opengl.missing.functions":  "Some functions are missing !",
	}
}
