# Table: hibp_password

Pwned Passwords are more than half a billion passwords which have previously been exposed in data breaches. Each password is stored as a `SHA-1` hash of a UTF-8 encoded password.

You can search by providing the `plaintext` password or the `hash` - which is the `SHA-1` hash of the password that you are looking for. Alternatively, you can also search by the `hash_prefix` which is a prefix (at least 5 hex-digits) of the `SHA-1` password.

## Examples

### Get the number of times a password hash has been compromised (by hash)

```sql
select
  hash,
  count
from
  hibp_password
where
  hash = '908f704ccaadfd86a74407d234c7bde30f2744fe';
```

### Get the number of times a password has been compromised (by plaintext)

```sql
select
  plaintext,
  hash,
  count
from
  hibp_password
where
  plaintext = '123457';
```

### Get the number of times a collection of passwords has been compromised (by prefix)

```sql
select
  plaintext,
  hash,
  count
from
  hibp_password
where
  hash_prefix = '908f704cc';
```
