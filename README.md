![image](https://hub.steampipe.io/images/plugins/turbot/hibp-social-graphic.png)

# Have I Been Pwned Plugin for Steampipe

Use SQL to query breaches, account breaches, pastes and passwords from [Have I Been Pwned](https://haveibeenpwned.com/).

- **[Get started →](https://hub.steampipe.io/plugins/turbot/hibp)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/hibp/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
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
  breach_date > '2022-01-01'
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

## Credits

- This plugin was based on the [wedtm/steampipe-plugin-hibp](https://gitlab.com/wedtm/steampipe-plugin-hibp) project. The client and table designs in that project served as the base for this plugin, thanks to [@wedtm](https://github.com/wedtm) for his amazing work!

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-hibp/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-hibp/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [HIBP Plugin](https://github.com/turbot/steampipe-plugin-hibp/labels/help%20wanted)
