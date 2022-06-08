# Table: hibp_password

Passwords indexed by HIBP.

You must include at _least_ 5 characters in the `prefix` field.

## Examples

### List the number of times a password hash has been compromised

```sql
select
  *
from
  hibp_password
where
  hash = '908f704ccaadfd86a74407d234c7bde30f2744fe'
```

### List the number of times a password has been compromised (by plain text)

```sql
select
  *
from
  hibp_password
where
  plain = '123457'
```
