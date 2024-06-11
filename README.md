# EdgeX Secrets Module

[![Build Status](https://jenkins.agile-edgex.org/view/EdgeX%20Foundry%20Project/job/agile-edgex/job/go-mod-secrets/job/main/badge/icon)](https://jenkins.agile-edgex.org/view/EdgeX%20Foundry%20Project/job/agile-edgex/job/go-mod-secrets/job/main/) [![Code Coverage](https://codecov.io/gh/agile-edgex/go-mod-secrets/branch/main/graph/badge.svg?token=KrqJoby1fK)](https://codecov.io/gh/agile-edgex/go-mod-secrets) [![Go Report Card](https://goreportcard.com/badge/github.com/agile-edgex/go-mod-secrets)](https://goreportcard.com/report/github.com/agile-edgex/go-mod-secrets) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/agile-edgex/go-mod-secrets?include_prereleases&sort=semver&label=latest-dev)](https://github.com/agile-edgex/go-mod-secrets/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/agile-edgex/go-mod-secrets?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/agile-edgex/go-mod-secrets)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/agile-edgex/go-mod-secrets) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/agile-edgex/go-mod-secrets)](https://github.com/agile-edgex/go-mod-secrets/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/agile-edgex/go-mod-secrets)](https://github.com/agile-edgex/go-mod-secrets/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/agile-edgex/teams/go-mod-secrets-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/agile-edgex/go-mod-secrets)](https://github.com/agile-edgex/go-mod-secrets/commits)

## Delayed Start Go Build Tags

The delayed start feature is by default built-in and included in this module.
If other go services would like to **exclude** the delayed start feature in the builds,
please do the go builds with tags `non_delayedstart` explicitly like the following:

```console
user$ go build <exclude_delayed_start_service.go> -tags non_delayedstart
```

## Community
- Wiki: https://wiki.agile-edgex.org/
- Website: https://www.agile-edgex.org/
- Mailing lists: https://lists.agile-edgex.org/mailman/listinfo

## License
[Apache-2.0](LICENSE)
