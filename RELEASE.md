# Development release steps

1. Update version number in `/cli/config.go`.
2. `git add .`, `git commit -m "<MESSAGE>"`, `git push origin main`.
3. `git tag -a v0.0.2 -m "new release: v0.0.2"` and `git push origin v0.0.2` (replacing with correct version; ensure it matches the version in `/cli/config.go`).
4. `export GITHUB_TOKEN=YOUR_GITHUB_TOKEN`.
5. Download the tagged and pushed version from `https://github.com/mattrltrent/quantum_crafter/archive/refs/tags/v0.0.2.tar.gz`.
6. Get the hash via `shasum -a 256 v0.0.2.tar.gz`.
7. Update `https://github.com/mattrltrent/homebrew-tap/blob/main/Formula/qc.rb` with new matching version number and `sha256`.