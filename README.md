![image](https://hub.steampipe.io/images/plugins/turbot/hibp-social-graphic.png)

# HIBP Plugin for Steampipe

Use SQL to query infrastructure including breaches, pastes, and more from Have I Been Pwned.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/hibp)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/hibp/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-hibp/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install hibp
```

Run a query:

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
  breach_date limit 5;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-hibp.git
cd steampipe-plugin-hibp
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/hibp.spc
```

Try it!

```
steampipe query
> .inspect hibp
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-aws/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [HIBP Plugin](https://github.com/turbot/steampipe-plugin-hibp/labels/help%20wanted)
