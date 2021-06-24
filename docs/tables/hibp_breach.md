# Table: hibp_breach

Breaches indexed and catalogued by HIBP.

## Examples

### Breaches from the last 3 months

```sql
select title, breach_date
from hibp_breach
where breach_date > CURRENT_DATE - INTERVAL '3 months'
```

### Unverified breaches

```sql
select title, pwn_count as size, breach_date
from hibp_breach
where is_verified = false
```
