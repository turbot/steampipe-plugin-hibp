---
title: "Steampipe Table: hibp_breached_account - Query HIBP Breached Accounts using SQL"
description: "Allows users to query breached accounts in Have I Been Pwned (HIBP), specifically the account details associated with each breach, providing insights into potential security risks and compromised data."
---

# Table: hibp_breached_account - Query HIBP Breached Accounts using SQL

Have I Been Pwned (HIBP) is a free resource that allows anyone to quickly assess if they may have been put at risk due to an online account of theirs having been compromised in a data breach. HIBP provides a comprehensive list of accounts that have been compromised in a data breach. The service is widely appreciated for its usefulness in tracking down potentially compromised accounts.

## Table Usage Guide

The `hibp_breached_account` table provides insights into breached accounts within Have I Been Pwned (HIBP). As a security analyst, explore account-specific details through this table, including breach names, breach dates, and compromised data types. Utilize it to uncover information about breaches, such as the specific accounts affected, the extent of the breach, and the types of data compromised.

**Important Notes**
- This table returns data similar to the `hibp_breach` table, with the requirement and addition of an `account` field.. While the `hibp_breaches` table will return all of the known breaches, this table can be used to find breaches for a particular account.
- This table requires an API key to be configured in the `hibp.spc` file.

## Examples

### List breaches from the last 3 months for an account
Discover the segments that have experienced security breaches within the last three months for a specific user account. This query can be used to monitor recent security incidents and take necessary actions to mitigate risks.

```sql+postgres
select
  title,
  breach_date
from
  hibp_breached_account
where
  breach_date > current_date - interval '3 months'
  and account = 'billy@example.com';
```

```sql+sqlite
select
  title,
  breach_date
from
  hibp_breached_account
where
  breach_date > date('now','-3 month')
  and account = 'billy@example.com';
```

### List unverified breaches for an account
Uncover the details of unverified security breaches associated with a specific account to understand the potential risk and take necessary actions. This information is useful for improving security measures and mitigating potential threats.

```sql+postgres
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  is_verified = false
  and account = 'billy@example.com';
```

```sql+sqlite
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  is_verified = 0
  and account = 'billy@example.com';
```

### List breaches for an account for the `"Passwords"` or `"Usernames"` data classes
Discover the instances of security breaches for a specific account, focusing on cases where either the usernames or passwords were compromised. This can be useful to understand the extent of data exposure and take necessary protective measures.

```sql+postgres
select
  distinct(title),
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  account = 'billy@example.com'
  and data_classes ?| array['Usernames','Passwords'];
```

```sql+sqlite
Error: SQLite does not support array operations and "?|" operator.
```

### List breaches for active Okta users (requires [Okta plugin](https://hub.steampipe.io/plugins/turbot/okta))
Determine the areas in which active Okta users may be at risk by identifying any breaches associated with their accounts. This helps in enhancing user security by proactively identifying potential vulnerabilities.

```sql+postgres
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  account in
  (
    select
      email
    from
      okta_user
    where
      filter = 'status eq "ACTIVE"'
  );
```

```sql+sqlite
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  account in
  (
    select
      email
    from
      okta_user
    where
      filter = 'status eq "ACTIVE"'
  );
```

### List breaches for LDAP users (requires [LDAP plugin](https://hub.steampipe.io/plugins/turbot/ldap))
This query is used to identify potential security breaches associated with LDAP users, particularly those in the 'Devs' group. It's a useful tool for maintaining the security of your system by pinpointing any instances where user details may have been compromised.

```sql+postgres
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  account in
  (
    select
      mail
    from
      ldap_user
    where
      filter = '(memberof=CN=Devs,OU=Steampipe,OU=SP,DC=sp,DC=turbot,DC=com)'
  );
```

```sql+sqlite
select
  title,
  pwn_count,
  breach_date
from
  hibp_breached_account
where
  account in
  (
    select
      mail
    from
      ldap_user
    where
      filter = '(memberof=CN=Devs,OU=Steampipe,OU=SP,DC=sp,DC=turbot,DC=com)'
  );
```