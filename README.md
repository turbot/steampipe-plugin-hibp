# Have I Been Pwned? plugin for Steampipe

## Query HIBP with SQL

Use SQL to query HIBP. Example:

```sql
select title, breach_date
from hibp_breach
where breach_date > CURRENT_DATE - INTERVAL '3 months'
```

## Get Started

### Installation

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

Most of HIBP API's are public, but if you wish to search by account or passwords, you'll need an API key. It can be set in the following ways:

- HIBP_API_KEY Environment Variable

These can also be set in the configuration file:
`vi ~/.steampipe/config/hibp.spc`

## Credits

Heavily inspired by [theapsgroup/steampipe-plugin-gitlab](https://github.com/theapsgroup/steampipe-plugin-gitlab)
