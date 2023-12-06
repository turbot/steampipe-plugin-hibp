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
