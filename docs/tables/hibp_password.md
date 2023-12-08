---
title: "Steampipe Table: hibp_password - Query HIBP Passwords using SQL"
description: "Allows users to query HIBP Passwords, providing insights into leaked passwords and their exposure count."
---

# Table: hibp_password - Query HIBP Passwords using SQL

Have I Been Pwned (HIBP) is a service that collects and analyzes hundreds of database dumps and pastes containing information about billions of leaked accounts. It allows users to search for their own information by entering their username or email address. Users can also sign up to be notified if their email address appears in future dumps.

## Table Usage Guide

The `hibp_password` table provides insights into leaked passwords within HIBP. As a security analyst, use this table to explore details about leaked passwords, including their exposure count. Utilize it to uncover information about passwords, such as their frequency of occurrence in breaches, aiding in the development of more secure password policies.

**Important Notes**
- You can search by providing the `plaintext` password or the `hash` which is the `SHA-1` hash of the password that you are looking for. Alternatively, you can also search by the `hash_prefix` which is a prefix (at least 5 hex-digits) of the `SHA-1` password.
- This table does not require an API key to be configured in the `hibp.spc` file.

## Examples

### Get the number of times a password hash has been compromised (by hash)
Determine the frequency of a specific password hash's compromise. This query is useful for assessing the security of a particular password, helping to decide whether it needs to be changed to maintain data protection.

```sql+postgres
select
  hash,
  count
from
  hibp_password
where
  hash = '908f704ccaadfd86a74407d234c7bde30f2744fe';
```

```sql+sqlite
select
  hash,
  count
from
  hibp_password
where
  hash = '908f704ccaadfd86a74407d234c7bde30f2744fe';
```

### Get the number of times a password has been compromised (by plaintext)
Gain insights into the security of a specific password by determining how many times it has been compromised, helping to assess password strength and potential vulnerabilities.

```sql+postgres
select
  plaintext,
  hash,
  count
from
  hibp_password
where
  plaintext = '123457';
```

```sql+sqlite
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
Analyze the frequency of password compromises to understand potential vulnerabilities. This could be useful in strengthening security measures by identifying commonly compromised passwords.

```sql+postgres
select
  plaintext,
  hash,
  count
from
  hibp_password
where
  hash_prefix = '908f704cc';
```

```sql+sqlite
select
  plaintext,
  hash,
  count
from
  hibp_password
where
  hash_prefix = '908f704cc';
```