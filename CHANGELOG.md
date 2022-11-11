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
