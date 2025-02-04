# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][],
and this project adheres to [Semantic Versioning][].

## [0.1.1][] — 2019-12-27

### Added

-   Debian PPA.
-   Brew tap.
-   Automated releases.
-   Repository example.
-   Private key handling.

### Changed

-   `copy` now uses `rsync` instead of `scp`.
-   Switched to `cobra` to get better UX and DX.

## [0.1.0][] — 2019-12-07

### Added

-   Configs recreation.

### Fixed

-   Empty lists handling.
-   Download link.

### Removed

-   Debian post-install logic.


## [0.1.0beta][] — 2019-12-06

### Added

-   Remote Hosts download.
-   Remote Hosts update.
-   SSH connect.
-   Local aliases.
-   Bash and Fish complitions.
-   `show` command.
-   `copy` command.

[keep a changelog]: https://keepachangelog.com/en/1.0.0/

[semantic versioning]: https://semver.org/spec/v2.0.0.html

[0.1.0beta]: https://github.com/mishamyrt/mysh/releases/tag/v0.1.0beta

[0.1.0]: https://github.com/mishamyrt/mysh/releases/tag/v0.1.0

[0.1.1]: https://github.com/mishamyrt/mysh/releases/tag/v0.1.1
