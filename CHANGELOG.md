## v1.1.1 [2025-04-18]

_Bug fixes_

- Fixed Linux AMD64 plugin build failures for `Postgres 14 FDW`, `Postgres 15 FDW`, and `SQLite Extension` by upgrading GitHub Actions runners from `ubuntu-20.04` to `ubuntu-22.04`.

## v1.1.0 [2025-04-17]

_Dependencies_

- Recompiled plugin with Go version `1.23.1`. ([#35](https://github.com/turbot/steampipe-plugin-hibp/pull/35))
- Recompiled plugin with [steampipe-plugin-sdk v5.11.5](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.11.5/CHANGELOG.md#v5115-2025-03-31) that addresses critical and high vulnerabilities in dependent packages. ([#35](https://github.com/turbot/steampipe-plugin-hibp/pull/35))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#33](https://github.com/turbot/steampipe-plugin-hibp/pull/33))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#33](https://github.com/turbot/steampipe-plugin-hibp/pull/33))

## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#29](https://github.com/turbot/steampipe-plugin-hibp/pull/29))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#29](https://github.com/turbot/steampipe-plugin-hibp/pull/29))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hibp/blob/main/docs/LICENSE). ([#29](https://github.com/turbot/steampipe-plugin-hibp/pull/29))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#28](https://github.com/turbot/steampipe-plugin-hibp/pull/28))

## v0.4.2 [2023-12-06]

_Bug fixes_

- Fixed the invalid Go module path of the plugin. ([#26](https://github.com/turbot/steampipe-plugin-hibp/pull/26))

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#19](https://github.com/turbot/steampipe-plugin-hibp/pull/19))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#17](https://github.com/turbot/steampipe-plugin-hibp/pull/17))
- Recompiled plugin with Go version `1.21`. ([#17](https://github.com/turbot/steampipe-plugin-hibp/pull/17))

## v0.3.0 [2023-04-10]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#12](https://github.com/turbot/steampipe-plugin-hibp/pull/12))

## v0.2.1 [2023-01-05]

_Bug fixes_

- Fixed `hibp_breach` and `hibp_password` tables to work without an API key. ([#10](https://github.com/turbot/steampipe-plugin-hibp/pull/10))

## v0.2.0 [2022-11-11]

_Dependencies_

- Recompiled plugin with [go-hibp v1.0.4](https://github.com/wneessen/go-hibp/releases/tag/v1.0.4). ([#6](https://github.com/turbot/steampipe-plugin-hibp/pull/6)) (Thanks [@wneessen](https://github.com/wneessen) for the contribution!)
- Recompiled plugin with [steampipe-plugin-sdk v4.1.8](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v418-2022-09-08) which increases the default open file limit. ([#7](https://github.com/turbot/steampipe-plugin-hibp/pull/7))

## v0.1.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#4](https://github.com/turbot/steampipe-plugin-hibp/pull/4))
- Recompiled plugin with Go version `1.19`. ([#4](https://github.com/turbot/steampipe-plugin-hibp/pull/4))

## v0.0.2 [2022-06-17]

_Enhancements_

- Added column `hash_prefix` to `hibp_password` table. ([#3](https://github.com/turbot/steampipe-plugin-hibp/pull/3))
- Recompiled plugin with [go-hibp v1.0.3](https://github.com/wneessen/go-hibp/releases/tag/v1.0.3). ([#3](https://github.com/turbot/steampipe-plugin-hibp/pull/3))

## v0.0.1 [2022-06-09]

_What's new?_

- New tables added
  - [hibp_breach](https://hub.steampipe.io/plugins/turbot/hibp/tables/hibp_breach)
  - [hibp_breached_account](https://hub.steampipe.io/plugins/turbot/hibp/tables/hibp_breached_account)
  - [hibp_password](https://hub.steampipe.io/plugins/turbot/hibp/tables/hibp_password)
  - [hibp_paste](https://hub.steampipe.io/plugins/turbot/hibp/tables/hibp_paste)

Thanks to [@wedtm](https://github.com/wedtm) for his ideas and original work on this plugin!
