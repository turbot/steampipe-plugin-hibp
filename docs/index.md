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
  breach_date > '2022-01-01'
```

```
+----------------+-------------------+----------+---------------------------+
| name           | compromised_count | verified | breach_date               |
+----------------+-------------------+----------+---------------------------+
| AmartFurniture | 108940            | true     | 2022-05-16T05:30:00+05:30 |
| BlackBerryFans | 174168            | true     | 2022-05-06T05:30:00+05:30 |
| Fanpass        | 112251            | true     | 2022-04-30T05:30:00+05:30 |
| GiveSendGo     | 89966             | true     | 2022-02-07T05:30:00+05:30 |
| CDEK           | 19218203          | false    | 2022-03-09T05:30:00+05:30 |
| Doxbin         | 370794            | true     | 2022-01-05T05:30:00+05:30 |
| NVIDIA         | 71335             | true     | 2022-02-23T05:30:00+05:30 |
| MacGeneration  | 101004            | true     | 2022-01-29T05:30:00+05:30 |
| PayHere        | 1580249           | true     | 2022-03-27T05:30:00+05:30 |
+----------------+-------------------+----------+---------------------------+
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
