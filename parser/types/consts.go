package types

/* Gets a list of things to look for */
func GenericChecks() [][]string {
	return [][]string{
		// Early fails
		{"runtime.early.fail", "Could not create the Java Virtual Machine", "Error occurred during initialization of VM"},
		{"java.args.unrecognised", "Unrecognized VM option"},
		// OOMs
		{"java.runtime.oom", "java.lang.OutOfMemoryError: Java heap space"},
		{"java.runtime.oom.start", "Could not reserve enough space for"},
		// OGL errors
		{"opengl.generic.error", "opengl error"},
		{"opengl.missing.functions", "Are you running essential and/or <1.13?", "stub called:"},
		// fabric basic stuff
		{"fabric.incompatible.mods", "Incompatible mod set"},
		// Optifine stuff
		{"optifine.render_region.on", "[OptiFine] OpenGL error: 1280 (Invalid enum), at: Copy VBO"},
		// Shader related issues
		{"load.shader.fail", "could not reload shaders", "shader version mismatch", "Failed to compile shader"},
	}
}
