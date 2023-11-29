# pojav-log-parser

A simple log parser to point out classic PojavLauncher problems.

Ideally, this wouldn't be a thing that exists. However, many pitfalls are present in certain versions of PojavLauncher, making level 1 support with customers a boring and repetitive task.

# API Documentation.

The API exposes only one method to the root of the exposed URL:

```
POST https://us-west1-pojavlauncher.cloudfunctions.net/pojav-parser-function/?should_localize=<boolean>
```

The body is simply the log that you want to analyze.

**Returns:** an `application/json` object

```typescript
interface response {
  version: {
    major_code: string;
    commit_number: number;
    commit_short: string;
    branch: string;
  };

  minecraft_version: {
    name: string;
    type: "UNKNOWN" | "VANILLA" | "OPTIFINE" | "FORGE" | "FABRIC" | "QUILT";
  };

  minecraft_username: string;
  renderer: "GL4ES_114" | "GL4ES_115" | "HOLY_GL4ES_114" | "HOLY_GL4ES_115" | "ANGLE" | "VIRGL" | "ZINK" | "UNKNOWN";
  env_variables: {
    [key: string]: string;
  };
  build_type: "DEBUG" | "PRODUCTION" | "UNKNOWN";
  architecture: "ARM_32" | "ARM_64" | "X86_32" | "X86_64" | "UNKNOWN";

  java_runtime: {
    version: number;
    source: "INTERNAL" | "EXTERNAL" | "UNKNOWN";
    type: "JDK" | "JRE" | "UNKNOWN";
  };
  java_arguments: {
    [key: string]: string;
  };

  errors: string[];
}
```
