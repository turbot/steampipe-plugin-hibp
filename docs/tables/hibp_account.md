# Table: hibp_account

This table returns data similar to the `hibp_breach` table, with the requirement of an `account` field.

While the `hibp_breaches` table will return all of the known breaches, there are millions of accounts associated with those
breaches. To find the breaches associated with a certain account, you must use this table.

## Examples

### Breaches from the last 3 months

```sql
select title, breach_date
from hibp_account
where breach_date > CURRENT_DATE - INTERVAL '3 months'
and account = 'account-exists@hibp-integration-tests.com'
```

### Unverified breaches

```sql
select title, pwn_count as size, breach_date
from hibp_account
where is_verified = false
and account = 'account-exists@hibp-integration-tests.com'
```
