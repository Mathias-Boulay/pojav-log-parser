package main

// Used to develop locally, not used by the function outside

import (
	"encoding/json"
	"log"

	// Blank-import the function package so the init() runs

	"github.com/mathias-boulay/parser"
)

func main() {
	file := `--------- beginning with launcher debug
	Info: Launcher version: dahlia-20230727-aeb2cda-v3_openjdk
	Info: Architecture: arm64
	Info: Device model: Xiaomi 22071212AG
	Info: API version: 33
	Info: Selected Minecraft version: 1.20.1-OptiFine_HD_U_I6_pre6
	Info: Custom Java arguments: ""
	Added custom env: TMPDIR=/data/user/0/net.kdt.pojavlaunch.beta/cache
	Added custom env: AWTSTUB_WIDTH=1898
	Added custom env: FORCE_VSYNC=false
	Added custom env: POJAV_NATIVEDIR=/data/app/~~ATYdNnlnBDrIW8ne30fHbg==/net.kdt.pojavlaunch.beta-uQ-cWziAsx0lJd2NgymzCA==/lib/arm64
	Added custom env: REGAL_GL_VERSION=4.5
	Added custom env: REGAL_GL_VENDOR=Android
	Added custom env: LIBGL_MIPMAP=3
	Added custom env: allow_higher_compat_version=true
	Added custom env: MESA_GLSL_CACHE_DIR=/data/user/0/net.kdt.pojavlaunch.beta/cache
	Added custom env: HOME=/storage/emulated/0/Android/data/net.kdt.pojavlaunch.beta/files
	Added custom env: PATH=/data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/bin:/product/bin:/apex/com.android.runtime/bin:/apex/com.android.art/bin:/system_ext/bin:/system/bin:/system/xbin:/odm/bin:/vendor/bin:/vendor/xbin
	Added custom env: LIBGL_NOINTOVLHACK=1
	Added custom env: force_glsl_extensions_warn=true
	Added custom env: LIBGL_NORMALIZE=1
	Added custom env: LD_LIBRARY_PATH=/data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/jli:/data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib:/system/lib64:/vendor/lib64:/vendor/lib64/hw:/data/app/~~ATYdNnlnBDrIW8ne30fHbg==/net.kdt.pojavlaunch.beta-uQ-cWziAsx0lJd2NgymzCA==/lib/arm64
	Added custom env: POJAV_RENDERER=opengles3_virgl
	Added custom env: LIBGL_ES=2
	Added custom env: VTEST_SOCKET_NAME=/data/user/0/net.kdt.pojavlaunch.beta/cache/.virgl_test
	Added custom env: MESA_LOADER_DRIVER_OVERRIDE=zink
	Added custom env: MESA_GLSL_VERSION_OVERRIDE=430
	Added custom env: JAVA_HOME=/data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17
	Added custom env: MESA_GL_VERSION_OVERRIDE=4.3
	Added custom env: allow_glsl_extension_directive_midshader=true
	Added custom env: REGAL_GL_RENDERER=Regal
	Added custom env: AWTSTUB_HEIGHT=854
	--------- beginning of main
	I/jrelog  (22703): dlopen libOSMesa_81.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjli.so success
	I/jrelog  (22703): dlopen libjvm.so failed: dlopen failed: library "libjvm.so" not found
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/server/libjvm.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libverify.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjava.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libnet.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libnio.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libawt.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libawt_h
	eadless.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libfreetype.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libfontmanager.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libawt_headless.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjawt.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libdt_socket.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjavajpeg.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libfontmanager.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/server/libjvm.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/server/libj
	sig.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libmanagement_agent.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libj2gss.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libverify.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/librmi.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libmanagement_ext.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libmlib_image.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libprefs.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libmanagement.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libnio.so success
	I
	/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjaas.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libj2pkcs11.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjava.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjli.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjdwp.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libextnet.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjimage.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libzip.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libjsig.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.
	pojavlaunch.beta/runtimes/Internal-17/lib/libfreetype.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libinstrument.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libsctp.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libawt.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/liblcms.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libnet.so success
	I/jrelog  (22703): dlopen /data/user/0/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/libawt_xawt.so success
	I/jrelog  (22703): dlopen /data/app/~~ATYdNnlnBDrIW8ne30fHbg==/net.kdt.pojavlaunch.beta-uQ-cWziAsx0lJd2NgymzCA==/lib/arm64/libopenal.so success
	I/jrelog  (22703): Done processing args
	I/jrelog  (22703): Found JLI lib
	I/jrelog  (22703): Calling JLI_Launch
	
	2023-11-10 17:44:51,299 main ERROR appender Console has no parameter that matches element Policies
	[17:44:51] [main/INFO]: Loading tweak class name optifine.OptiFineTweaker
	[17:44:51] [main/INFO]: Using primary tweak class name optifine.OptiFineTweaker
	[17:44:51] [main/INFO]: Calling tweak class optifine.OptiFineTweaker
	[17:44:51] [main/INFO]: [OptiFine] OptiFineTweaker: acceptOptions
	[17:44:51] [main/INFO]: [OptiFine] OptiFineTweaker: injectIntoClassLoader
	[17:44:51] [main/INFO]: [OptiFine] OptiFine ClassTransformer
	[17:44:51] [main/INFO]: [OptiFine] OptiFine ZIP file: /storage/emulated/0/Android/data/net.kdt.pojavlaunch.beta/files/.minecraft/libraries/optifine/OptiFine/1.20.1_HD_U_I6_pre6/OptiFine-1.20.1_HD_U_I6_pre6.jar
	[17:44:51] [main/INFO]: [OptiFine] OptiFineTweaker: getLaunchArguments
	[17:44:51] [main/INFO]: [OptiFine] OptiFineTweaker: getLaunchTarget
	[17:44:51] [main/INFO]: Launching wrapped minecraft {net.minecraft.client.main.Main}
	2023-11-10 17:44:51,856 main ERROR appender Console has no parameter that matches element Policies
	[17:44:52] [main/WARN]: Failed retrieving info for group processor
	java.lang.NoClassDefFoundError: Could not initialize class com.sun.jna.Native
		at com.sun.jna.NativeLong.<clinit>(NativeLong.java:35) ~[jna-5.12.1.jar:5.12.1 (b0)]
		at oshi.util.FileUtil.readNativeLongFromBuffer(FileUtil.java:240) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.driver.linux.proc.Auxv.queryAuxv(Auxv.java:59) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.software.os.linux.LinuxOperatingSystem.<clinit>(LinuxOperatingSystem.java:106) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxCentralProcessor.initProcessorCounts(LinuxCentralProcessor.java:171) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.common.AbstractCentralProcessor.<init>(AbstractCentralProcessor.java:83) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxCentralProcessor.<init>(LinuxCentralProcessor.java:71) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxHardwareAbstractionLayer.createProcessor(LinuxHardwareAbstractionLayer.java:62) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.util.Memoizer$1.get(Memoizer.java:87) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.common.AbstractHardwareAbstractionLayer.getProcessor(AbstractHardwareAbstractionLayer.java:68) ~[oshi-core-6.2.2.jar:6.2.2]
		at ab.c(SourceFile:75) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:82) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:75) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.c(SourceFile:52) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:82) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.<init>(SourceFile:52) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at o.<init>(CrashReport.java:34) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at o.h(CrashReport.java:345) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at net.minecraft.client.main.Main.main(SourceFile:164) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method) ~[?:?]
		at jdk.internal.reflect.NativeMethodAccessorImpl.invoke(Unknown Source) ~[?:?]
		at jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(Unknown Source) ~[?:?]
		at java.lang.reflect.Method.invoke(Unknown Source) ~[?:?]
		at net.minecraft.launchwrapper.Launch.launch(Launch.java:159) ~[launchwrapper-of-2.3.jar:2.3]
		at net.minecraft.launchwrapper.Launch.main(Launch.java:30) ~[launchwrapper-of-2.3.jar:2.3]
	Caused by: java.lang.ExceptionInInitializerError: Exception java.lang.UnsatisfiedLinkError: /storage/emulated/0/Android/data/net.kdt.pojavlaunch.beta/files/.cache/JNA/temp/jna13793163275006896723.tmp: dlopen failed: library "/storage/emulated/0/Android/data/net.kdt.pojavlaunch.beta/files/.cache/JNA/temp/jna13793163275006896723.tmp" needed or dlopened by "/data/data/net.kdt.pojavlaunch.beta/runtimes/Internal-17/lib/server/libjvm.so" is not accessible for the namespace "clns-4" [in thread "main"]
		at jdk.internal.loader.NativeLibraries.load(Native Method) ~[?:?]
		at jdk.internal.loader.NativeLibraries$NativeLibraryImpl.open(Unknown Source) ~[?:?]
		at jdk.internal.loader.NativeLibraries.loadLibrary(Unknown Source) ~[?:?]
		at jdk.internal.loader.NativeLibraries.loadLibrary(Unknown Source) ~[?:?]
		at java.lang.ClassLoader.loadLibrary(Unknown Source) ~[?:?]
		at java.lang.Runtime.load0(Unknown Source) ~[?:?]
		at java.lang.System.load(Unknown Source) ~[?:?]
		at com.sun.jna.Native.loadNativeDispatchLibraryFromClasspath(Native.java:1045) ~[jna-5.12.1.jar:5.12.1 (b0)]
		at com.sun.jna.Native.loadNativeDispatchLibrary(Native.java:1015) ~[jna-5.12.1.jar:5.12.1 (b0)]
		at com.sun.jna.Native.<clinit>(Native.java:221) ~[jna-5.12.1.jar:5.12.1 (b0)]
		at com.sun.jna.platform.linux.Udev.<clinit>(Udev.java:37) ~[jna-platform-5.12.1.jar:5.12.1 (b0)]
		at oshi.software.os.linux.LinuxOperatingSystem.<clinit>(LinuxOperatingSystem.java:93) ~[oshi-core-6.2.2.jar:6.2.2]
		... 21 more
	[17:44:52] [main/WARN]: Failed retrieving info for group memory
	java.lang.NoClassDefFoundError: Could not initialize class oshi.software.os.linux.LinuxOperatingSystem
		at oshi.hardware.platform.linux.LinuxGlobalMemory.<clinit>(LinuxGlobalMemory.java:47) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxHardwareAbstractionLayer.createMemory(LinuxHardwareAbstractionLayer.java:57) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.util.Memoizer$1.get(Memoizer.java:87) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.common.AbstractHardwareAbstractionLayer.getMemory(AbstractHardwareAbstractionLayer.java:80) ~[oshi-core-6.2.2.jar:6.2.2]
		at ab.a(SourceFile:77) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:82) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:77) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.c(SourceFile:52) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:82) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.<init>(SourceFile:52) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at o.<init>(CrashReport.java:34) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at o.h(CrashReport.java:345) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at net.minecraft.client.main.Main.main(SourceFile:164) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method) ~[?:?]
		at jdk.internal.reflect.NativeMethodAccessorImpl.invoke(Unknown Source) ~[?:?]
		at jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(Unknown Source) ~[?:?]
		at java.lang.reflect.Method.invoke(Unknown Source) ~[?:?]
		at net.minecraft.launchwrapper.Launch.launch(Launch.java:159) ~[launchwrapper-of-2.3.jar:2.3]
		at net.minecraft.launchwrapper.Launch.main(Launch.java:30) ~[launchwrapper-of-2.3.jar:2.3]
	Caused by: java.lang.ExceptionInInitializerError: Exception java.lang.NoClassDefFoundError: Could not initialize class com.sun.jna.Native [in thread "main"]
		at com.sun.jna.NativeLong.<clinit>(NativeLong.java:35) ~[jna-5.12.1.jar:5.12.1 (b0)]
		at oshi.util.FileUtil.readNativeLongFromBuffer(FileUtil.java:240) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.driver.linux.proc.Auxv.queryAuxv(Auxv.java:59) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.software.os.linux.LinuxOperatingSystem.<clinit>(LinuxOperatingSystem.java:106) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxCentralProcessor.initProcessorCounts(LinuxCentralProcessor.java:171) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.common.AbstractCentralProcessor.<init>(AbstractCentralProcessor.java:83) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxCentralProcessor.<init>(LinuxCentralProcessor.java:71) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.platform.linux.LinuxHardwareAbstractionLayer.createProcessor(LinuxHardwareAbstractionLayer.java:62) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.util.Memoizer$1.get(Memoizer.java:87) ~[oshi-core-6.2.2.jar:6.2.2]
		at oshi.hardware.common.AbstractHardwareAbstractionLayer.getProcessor(AbstractHardwareAbstractionLayer.java:68) ~[oshi-core-6.2.2.jar:6.2.2]
		at ab.c(SourceFile:75) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:82) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		at ab.a(SourceFile:75) ~[1.20.1-OptiFine_HD_U_I6_pre6.jar:?]
		... 12 more
	[17:44:52] [main/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.eventbus.api.Event$Result
	[17:44:52] [main/INFO]: [OptiFine] (Reflector) Method not present: net.minecraftforge.common.extensions.IForgeEntity.canUpdate
	[17:44:52] [main/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.logging.CrashReportExtender
	[17:44:53] [main/INFO]: [OptiFine] (Reflector) Method not present: ane.create
	[17:44:56] [Datafixer Bootstrap/INFO]: 188 Datafixer optimizations took 438 milliseconds
	[17:44:58] [Render thread/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.client.ForgeHooksClient
	[17:44:58] [Render thread/INFO]: [STDERR]: [LWJGL] Failed to load a library. Possible solutions:
		a) Add the directory that contains the shared library to -Djava.library.path or -Dorg.lwjgl.librarypath.
		b) Add the JAR that contains the shared library to the classpath.
	[17:44:58] [Render thread/INFO]: [STDERR]: [LWJGL] Enable debug mode with -Dorg.lwjgl.util.Debug=true for better diagnostics.
	[17:44:58] [Render thread/INFO]: [STDERR]: [LWJGL] Enable the SharedLibraryLoader debug mode with -Dorg.lwjgl.util.DebugLoader=true for better diagnostics.
	[17:44:58] [Render thread/INFO]: Environment: authHost='https://authserver.mojang.com', accountsHost='https://api.mojang.com', sessionHost='https://sessionserver.mojang.com', servicesHost='https://api.minecraftservices.com', name='PROD'
	[17:44:59] [Yggdrasil Key Fetcher/ERROR]: Failed to request yggdrasil public key
	com.mojang.authlib.exceptions.AuthenticationUnavailableException: Cannot contact authentication server
		at com.mojang.authlib.yggdrasil.YggdrasilAuthenticationService.makeRequest(YggdrasilAuthenticationService.java:119) ~[YggdrasilAuthenticationService.class:?]
		at com.mojang.authlib.yggdrasil.YggdrasilAuthenticationService.makeRequest(YggdrasilAuthenticationService.java:91) ~[YggdrasilAuthenticationService.class:?]
		at com.mojang.authlib.yggdrasil.YggdrasilServicesKeyInfo.fetch(YggdrasilServicesKeyInfo.java:94) ~[YggdrasilServicesKeyInfo.class:?]
		at com.mojang.authlib.yggdrasil.YggdrasilServicesKeyInfo.lambda$get$1(YggdrasilServicesKeyInfo.java:81) ~[YggdrasilServicesKeyInfo.class:?]
		at java.util.concurrent.Executors$RunnableAdapter.call(Unknown Source) ~[?:?]
		at java.util.concurrent.FutureTask.runAndReset(Unknown Source) ~[?:?]
		at java.util.concurrent.ScheduledThreadPoolExecutor$ScheduledFutureTask.run(Unknown Source) ~[?:?]
		at java.util.concurrent.ThreadPoolExecutor.runWorker(Unknown Source) ~[?:?]
		at java.util.concurrent.ThreadPoolExecutor$Worker.run(Unknown Source) ~[?:?]
		at java.lang.Thread.run(Unknown Source) ~[?:?]
	Caused by: java.net.SocketException: Connection reset by peer
		at sun.nio.ch.NioSocketImpl.implWrite(Unknown Source) ~[?:?]
		at sun.nio.ch.NioSocketImpl.write(Unknown Source) ~[?:?]
		at sun.nio.ch.NioSocketImpl$2.write(Unknown Source) ~[?:?]
		at java.net.Socket$SocketOutputStream.write(Unknown Source) ~[?:?]
		at sun.security.ssl.SSLSocketOutputRecord.flush(Unknown Source) ~[?:?]
		at sun.security.ssl.HandshakeOutStream.flush(Unknown Source) ~[?:?]
		at sun.security.ssl.ClientHello$ClientHelloKickstartProducer.produce(Unknown Source) ~[?:?]
		at sun.security.ssl.SSLHandshake.kickstart(Unknown Source) ~[?:?]
		at sun.security.ssl.ClientHandshakeContext.kickstart(Unknown Source) ~[?:?]
		at sun.security.ssl.TransportContext.kickstart(Unknown Source) ~[?:?]
		at sun.security.ssl.SSLSocketImpl.startHandshake(Unknown Source) ~[?:?]
		at sun.security.ssl.SSLSocketImpl.startHandshake(Unknown Source) ~[?:?]
		at sun.net.www.protocol.https.HttpsClient.afterConnect(Unknown Source) ~[?:?]
		at sun.net.www.protocol.https.AbstractDelegateHttpsURLConnection.connect(Unknown Source) ~[?:?]
		at sun.net.www.protocol.http.HttpURLConnection.getInputStream0(Unknown Source) ~[?:?]
		at sun.net.www.protocol.http.HttpURLConnection.getInputStream(Unknown Source) ~[?:?]
		at sun.net.www.protocol.https.HttpsURLConnectionImpl.getInputStream(Unknown Source) ~[?:?]
		at com.mojang.authlib.HttpAuthenticationService.performGetRequest(HttpAuthenticationService.java:140) ~[HttpAuthenticationService.class:?]
		at com.mojang.authlib.yggdrasil.YggdrasilAuthenticationService.makeRequest(YggdrasilAuthenticationService.java:96) ~[YggdrasilAuthenticationService.class:?]
		... 9 more
	[17:44:59] [Render thread/INFO]: Setting user: Ghoib
	Registered forkAndExec
	[17:44:59] [Render thread/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.client.settings.KeyConflictContext
	[17:44:59] [Render thread/INFO]: [OptiFine] (Reflector) Method not present: enl.getKeyModifier
	[17:44:59] [Render thread/ERROR]: Error parsing option value off for option Fullscreen: Not a boolean: "off"
	[17:44:59] [Render thread/ERROR]: Error parsing option value  for option Text Background Opacity: Not a boolean: ""
	[17:44:59] [Render thread/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.forge.snapshots.ForgeSnapshotsMod
	[17:44:59] [Render thread/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.versions.forge.ForgeVersion
	[17:44:59] [Render thread/INFO]: Backend library: LWJGL version 3.2.3 SNAPSHOT
	VirGL: OSMesa buffer flush is DISABLED!
	VirGL: libvirgl_test_server = 0x76fbaeeff0596c53
	EGLBridge: Initializing
	EGLBridge: Initialized!
	EGLBridge: ThreadID=22774
	EGLBridge: EGLDisplay=0xb4000072bf336500, EGLSurface=0x71f796c500
	VirGL: created EGL context 0x71f796c680
	EGLBridge: eglMakeCurrent() succeed!
	VirGL: vtest_main = 0x71f8ab3278
	VirGL: Calling VTest server's main function
	OSMDroid: width=1898;height=854, reserving 6483568 bytes for frame buffer
	OSMDroid: created frame buffer
	[17:45:00] [Render thread/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.internal.BrandingControl
	[17:45:00] [Render thread/INFO]: [OptiFine] (Reflector) Class not present: net.minecraftforge.fml.loading.ImmediateWindowHandler
	OSMDroid: generating context`

	json, _ := json.Marshal(parser.ParseLog(file))
	log.Printf("%s", json)
}
