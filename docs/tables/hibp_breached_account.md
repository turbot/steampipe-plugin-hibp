# Table: hibp_account

The most common use of the API is to return a list of all breaches a particular account has been involved in. This table returns data similar to the `hibp_breach` table, with the requirement and addition of an `account` field.

While the `hibp_breaches` table will return all of the known breaches, this table can be used to find breaches for a particular account.

## Examples

### Breaches from the last 3 months for an account

```sql
select
  title,
  breach_date
from
  hibp_breached_account
where
  breach_date > current_date - interval '3 months'
  and account = 'billy@example.com'
```

### Unverified breaches for an account

```sql
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  is_verified = false
  and account = 'billy@example.com'
```

### List all breaches for an account for the `"Passwords"` or `"Usernames"` data classes

```sql
select
  distinct(title),
  pwn_count,
  breach_date
from
  hibp_breached_account,
  jsonb_array_elements(data_classes) as dc
where
  account = 'billy@example.com'
  and dc::text in
  (
    '"Passwords"',
    '"Usernames"'
  )
```

### List all breaches for accounts for all `Active` users in the organization (uses the [Okta Plugin](https://hub.steampipe.io/plugins/turbot/okta))

```sql
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account,
where
  account in
  (
    select
      email
    from
      okta_user
    where
      filter = 'status eq "ACTIVE"'
  )
```
