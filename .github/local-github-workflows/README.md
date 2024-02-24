
# pr open to main
 act pull_request-s GITHUB_TOKEN="$(gh auth token)" -e .github/local-github-workflows/pr-open.json -W .github/workflows/main-pr.yaml --container-architecture linux/arm64

# pr closed to main

act pull_request -e .github/local-github-workflows/pr-merge.json -s GITHUB_TOKEN="$(gh auth token)" -W .github/workflows/main-pr-merge-workflow.yaml --container-architecture linux/arm64 
