# Mono Repo For Unoplat Utilities

## Image Scan Utility

- Ability to parse Helm values yaml to extract images/containers present and store their respective paths.

Example:
Input
```
app:
  image:
    registry: docker.io
    repository: myapp
    tag: "1.0"
ap2:
  image:
    registry: docker.io
    repository: app2
    tag: "1.1"
```
Output
```
{
    "docker.io/myapp:1.0": [
        "app.image.tag"
    ],
    "docker.io/app2:1.1": [
        "ap2.image.tag"
    ]
}
```

This enables us to scan for vulnerabilities inside container images and scan them using trivy and then patch them using copacectic and modify the values for patched images automatically using their known locations.
