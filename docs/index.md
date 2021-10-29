---
organization: turbot
category: ["internet"]
icon_url: "/images/plugins/turbot/hibp.svg"
brand_color: "#3A9AC4"
display_name: Security
name: hibp
description: Steampipe plugin to query Have I Been Pwned breaches, pastes, and passwords
og_description: Query compromised data with SQL! Open source CLI. No DB required.
og_image: "TBD"
---

# HaveIBeenPwned? + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

Compromised accounts and passwords available via [Have I Been Pwned](https://haveibeenpwned.com) (HIBP) along with the breaches that compromised them are searchable.

For example:

```sql
select
  name,
  pwn_count as compromised_count,
  is_verified as verified,
  breach_date
from
  hibp_breach
where
  breach_date > '2020-01-01'
order by
  breach_date limit 5
```

```
+-------------+-------------------+----------+----------------------+
| name        | compromised_count | verified | breach_date          |
+-------------+-------------------+----------+----------------------+
| HTCMania    | 1488089           | true     | 2020-01-04T00:00:00Z |
| MobiFriends | 3512952           | true     | 2020-01-06T00:00:00Z |
| Zoosk2020   | 23927853          | true     | 2020-01-12T00:00:00Z |
| Mathway     | 25692862          | true     | 2020-01-13T00:00:00Z |
| MeetMindful | 1422717           | true     | 2020-01-26T00:00:00Z |
+-------------+-------------------+----------+----------------------+
```

## Get started

### Install

Download and install the latest HIBP plugin:

```shell
steampipe plugin install hibp
```

### Configuration

Installing the latest HIBP plugin will create a config file (`~/.steampipe/config/hibp.spc`) with a single connection named `hibp`:

```hcl
connection "hibp" {
  plugin     = "hibp"
  api_key     = "use-it-if-you-got-it"
}
```

- `api_key` (required) - Making calls to the HIBP API requires a key. [Get API key](https://haveibeenpwned.com/API/Key)

## Get Involved

- Open source: https://github.com/turbot/steampipe-plugin-aws
- Community: [Slack Channel](https://steampipe.io/community/join)
