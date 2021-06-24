# Table: hibp_account

Breaches indexed and catalogued by HIBP filtered by account compromised.
Account can either be an email address or a phone number.

## Examples

### Breaches from the last 3 months

```sql
select title, breach_date
from hibp_breach
where breach_date > CURRENT_DATE - INTERVAL '3 months'
and account = 'account-exists@hibp-integration-tests.com'
```

### Unverified breaches

```sql
select title, pwn_count as size, breach_date
from hibp_breach
where is_verified = false
and account = 'account-exists@hibp-integration-tests.com'
```
