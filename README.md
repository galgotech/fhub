# Function Hub (FHub)

FHub is a function repository based on a **function specification**. This allows you to reuse code at the function level, build software that is scalable and flexible, provides interoperability between programming languages, and helps to developers code maintaince.

---

## Benefits
* **Increase productivity** helping developers save time by providing a central location to find functions.
* **Improve quality** of software by providing a way to test functions, track functions changes, keep version compatibility with old versions.
* **Reduced costs** of software development by providing a way to reuse functions avoiding duplicating work.


## Function specification

Simple schema using CUE. [Complete schema](./schema.cue)

```cue
name: "identity"
specVersion: "0.1"
version: "1.0"
env: [
  "MODE"
]
import: [
  "./fhub/fhub.file1.cue"
]
packages: {
  identity: {
    import: "fhub.dev/identity"
    launch: "Init"
    build: {
      container: {
        image: "golang:1.20"
      }
    }
    serving: {
      http: {
        url: "https://identity:3000/"
      }
    }
  }
}
functions: {
  login: {
    package: "identity"
    launch: "Login"
    input: {
      username: string
      password: string
    }
    output: {
      success: bool
    }
  }
}
```

### Scope

* **Standardized function definition** promoting code reusability, share, and use across different projects.
* **Serving** allows functions to be hosted and accessible as services with different protocols (http, grpc, ...).
* **Build** managing dependencies, compiling or packaging functions into executable units, and deployment to different environments (podman, docker, microvm, ...).


## Runtime
Responsible to interpret schema, builder, and servirng.
* [GO Lang](https://github.com/galgotech/fhub-runtime)
