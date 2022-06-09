# Table: hibp_breached_account

A "breach" is an instance of a system having been compromised by an attacker and the data disclosed. For example, Adobe was a breach, Gawker was a breach etc. This returns the details of each of breach in the system which currently stands at 606 breaches.

## Examples

### List breaches from the last 3 months

```sql
select
  title,
  breach_date
from
  hibp_breach
where
  breach_date > current_date - interval '3 months';
```

### List unverified breaches

```sql
select
  title,
  pwn_count,
  breach_date
from
  hibp_breach
where
  is_verified = false;
```

### List breaches for the `"Passwords"` or `"Usernames"` data classes

```sql
select
  distinct(title),
  pwn_count as size,
  breach_date
from
  hibp_breach,
  jsonb_array_elements(data_classes) as dc
where
  dc::text in
  (
    '"Passwords"',
    '"Usernames"'
  );
```
