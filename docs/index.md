---
organization: turbot
category: ["security"]
icon_url: "/images/plugins/turbot/hibp.svg"
brand_color: "#3A9AC4"
display_name: Have I Been Pwned
name: hibp
description: Steampipe plugin to query breaches, account breaches, pastes and passwords from Have I Been Pwned.
og_description: Query HIBP data with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/hibp-social-graphic.png"
---

# Have I Been Pwned + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Have I Been Pwned](https://haveibeenpwned.com) (HIBP) is an online searchable index of data breaches where anyone can quickly assess if they may have been put at risk due to an online account of theirs having been compromised or "pwned" in a data breach.

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

- **[Table definitions & examples →](/plugins/turbot/hibp/tables)**

## Get started

### Install

Download and install the latest HIBP plugin:

```shell
steampipe plugin install hibp
```

### Configuration

Installing the latest hibp plugin will create a config file (`~/.steampipe/config/hibp.spc`) with a single connection named `hibp`:

```hcl
connection "hibp" {
  plugin  = "hibp"

  # Requests to HIBP API needs to carry an API KEY.
  # You can get one at https://haveibeenpwned.com/API/Key
  api_key = "03ef6bfxxxxxxxxxxxxxxx8ad568286b"
}
```

- `api_key` - (required) The API key to access the HIBP API. Can also be set with the `HIBP_API_KEY` environment variable.

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-hibp
- Community: [Slack Channel](https://steampipe.io/community/join)
