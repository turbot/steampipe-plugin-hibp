# Table: hibp_password

Pwned Passwords are more than half a billion passwords which have previously been exposed in data breaches. Each password is stored as a `SHA-1` hash of a UTF-8 encoded password.

You can also search by the `plain text` password which is converted to a `SHA-1` hash under the hood before calling the HIBP API.

## Examples

### List the number of times a password hash has been compromised

```sql
select
  password_hash,
  count
from
  hibp_password
where
  password_hash = '908f704ccaadfd86a74407d234c7bde30f2744fe'
```

### List the number of times a password has been compromised (by plain text)

```sql
select
  password,
  password_hash,
  count
from
  hibp_password
where
  password = '123457'
```
