---
organization: wedtm
category: ["internet"]
icon_url: "/images/plugins/wedtm/hibp.svg"
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
```

```
+------------------+-------------------+----------+---------------------+
| name             | compromised_count | verified | breach_date         |
+------------------+-------------------+----------+---------------------+
| DominosIndia     | 22527655          | true     | 2021-03-24 00:00:00 |
| Descomplica      | 4845378           | true     | 2021-03-14 00:00:00 |
| CityBee          | 110156            | true     | 2021-02-05 00:00:00 |
| DailyQuiz        | 8032404           | true     | 2021-01-13 00:00:00 |
| CardingMafia     | 297744            | true     | 2021-03-18 00:00:00 |
| Emotet           | 4324770           | true     | 2021-01-27 00:00:00 |
| Gab              | 66521             | true     | 2021-02-26 00:00:00 |
| PhoneHouse       | 5223350           | true     | 2021-04-08 00:00:00 |
| MangaDex         | 2987329           | true     | 2021-03-22 00:00:00 |
| NurseryCam       | 10585             | true     | 2021-02-12 00:00:00 |
| Oxfam            | 1834006           | true     | 2021-01-20 00:00:00 |
| ParkMobile       | 20949825          | true     | 2021-03-21 00:00:00 |
| Liker            | 465141            | true     | 2021-03-08 00:00:00 |
| WedMeGood        | 1306723           | true     | 2021-01-06 00:00:00 |
| SuperVPNGeckoVPN | 20339937          | true     | 2021-02-25 00:00:00 |
| Ticketcounter    | 1921722           | true     | 2021-02-22 00:00:00 |
| WeLeakInfo       | 11788             | true     | 2021-03-08 00:00:00 |
| Astoria          | 11498146          | false    | 2021-01-26 00:00:00 |
+------------------+-------------------+----------+---------------------+
```

## Documentation

- [hibp_account](/docs/tables/hibp_account.md)
- [hibp_breach](/docs/tables/hibp_breach.md)
- [hibp_password](/docs/tables/hibp_password.md)
- [hibp_paste](/docs/tables/hibp_paste.md)

## Get started

### Install

Download and install the latest HIBP plugin:

```shell
steampipe plugin install wedtm/hibp
```

Or if you prefer, you can clone this repository and build/install from source directly.

```shell
go build -o steampipe-plugin-hibp.plugin

mv steampipe-plugin-hibp.plugin ~/.steampipe/plugins/hub.steampipe.io/plugins/wedtm/hibp@latest/steampipe-plugin-hibp.plugin

cp config/hibp.spc ~/.steampipe/config/hibp.spc
```

### Configuration

Installing the latest HIBP plugin will create a config file (`~/.steampipe/config/hibp.spc`) with a single connection named `wedtm/hibp`:

```hcl
connection "hibp" {
  plugin     = "wedtm/hibp"
  api_key     = "use-it-if-you-got-it"
}
```

## Get involved

- Open source: https://gitlab.com/wedtm/steampipe-plugin-hibp
- Community: [Discussion forums](https://github.com/turbot/steampipe/discussions)
