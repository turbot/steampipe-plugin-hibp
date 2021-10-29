# Table: hibp_password

Passwords indexed by HIBP.

You must include at _least_ 5 characters in the `prefix` field.

## Examples

### Show the number of times this password has been compromised

```sql
select
  *
from
  hibp_password
where
  hash = '5BAA61E4C9B93F3F0682250B6CF8331B7EE68FD8';
```

### Show all hashes for a 5 character prefix

```sql
select
  *
from
  hibp_password
where
  prefix = '5BAA6'
```
