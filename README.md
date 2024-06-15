# Install

Ensure wails3 is installed per docs
Install templ cli `go install github.com/a-h/templ/cmd/templ@latest`

## Taskfile changes

```yaml
  build:frontend:
    summary: Build the frontend project
    dir: frontend
    sources:
      - "**/*"
    deps:
      - task: generate:bindings
        vars:
          BUILD_FLAGS: '{{.BUILD_FLAGS}}'
    cmds:
+      - templ generate
+      - npx tailwind build -i frontend/static/styles.css -o frontend/static/tailwind.css
```
