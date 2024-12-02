# EdgeX Secrets Module

[![Build Status](https://jenkins.agile-edge.org/view/EdgeX%20Foundry%20Project/job/agile-edge/job/go-mod-secrets/job/main/badge/icon)](https://jenkins.agile-edge.org/view/EdgeX%20Foundry%20Project/job/agile-edge/job/go-mod-secrets/job/main/) [![Code Coverage](https://codecov.io/gh/agile-edge/go-mod-secrets/branch/main/graph/badge.svg?token=KrqJoby1fK)](https://codecov.io/gh/agile-edge/go-mod-secrets) [![Go Report Card](https://goreportcard.com/badge/github.com/agile-edge/go-mod-secrets)](https://goreportcard.com/report/github.com/agile-edge/go-mod-secrets) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/agile-edge/go-mod-secrets?include_prereleases&sort=semver&label=latest-dev)](https://github.com/agile-edge/go-mod-secrets/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/agile-edge/go-mod-secrets?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/agile-edge/go-mod-secrets)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/agile-edge/go-mod-secrets) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/agile-edge/go-mod-secrets)](https://github.com/agile-edge/go-mod-secrets/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/agile-edge/go-mod-secrets)](https://github.com/agile-edge/go-mod-secrets/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/agile-edge/teams/go-mod-secrets-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/agile-edge/go-mod-secrets)](https://github.com/agile-edge/go-mod-secrets/commits)

## Delayed Start Go Build Tags

The delayed start feature is by default built-in and included in this module.
If other go services would like to **exclude** the delayed start feature in the builds,
please do the go builds with tags `non_delayedstart` explicitly like the following:

```console
user$ go build <exclude_delayed_start_service.go> -tags non_delayedstart
```

## Community
- Wiki: https://wiki.agile-edge.org/
- Website: https://www.agile-edge.org/
- Mailing lists: https://lists.agile-edge.org/mailman/listinfo

## License
[Apache-2.0](LICENSE)
