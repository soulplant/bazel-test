```
# Generate BUILD files
bazel run //:gazelle

# Build binary
bazel build //foo/cmd/foo

bazel-bin/foo/cmd/foo/foo_/foo
```
