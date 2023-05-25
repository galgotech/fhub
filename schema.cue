name: string
specVersion: "0.1"
version: string
env: {
  [string]: string
}
packages: {
  [string]: {
    import: string
    launch?: string
    build: {
      container: {
        image?: string
        context?: string
        dockerfile?: string | *"Dockerfile"
        target: string
      }
    }
    serving: {
      http?: {
        url: string
      }
      grpc?: {}
    }
  }
}
functions: {
  [string]: {
    package: string
    launch: string
    inputs: {
      [string]: number | string | bool
    }
    outputs: {
      [string]: number | string | bool
    }
  }
}