---
title: "Steampipe Table: hibp_breach - Query HIBP Breaches using SQL"
description: "Allows users to query HIBP Breaches, specifically the data about breaches registered in the Have I Been Pwned database, providing insights into data leak incidents and potential vulnerabilities."
---

# Table: hibp_breach - Query HIBP Breaches using SQL

Have I Been Pwned (HIBP) is a service that allows users to check if their personal data has been compromised by data breaches. It collects and analyzes hundreds of database dumps and pastes containing information about billions of leaked accounts. The HIBP Breach resource provides information about the specific data breach incidents.

## Table Usage Guide

The `hibp_breach` table provides insights into data leak incidents registered in the Have I Been Pwned database. As a security analyst, explore breach-specific details through this table, including the breach name, domain, date, and associated data classes. Utilize it to uncover information about specific breaches, such as the affected accounts, the nature of the leaked data, and the actions taken to mitigate the breach.

**Important Notes**
- This table does not require an API key to be configured in the `hibp.spc` file.

## Examples

### List breaches from the last 3 months
Explore recent security breaches to understand potential vulnerabilities and patterns. This query is particularly useful for identifying recent threats and enhancing security measures accordingly.

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
Discover the segments that have experienced unverified security breaches. This can be useful in assessing potential vulnerabilities and prioritizing areas for security enhancement.

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
Discover the instances of security breaches involving either passwords or usernames. This can be helpful in understanding the magnitude and timing of such incidents, which can aid in improving data security measures.

```sql
select
  distinct(title),
  pwn_count as size,
  breach_date
from
  hibp_breach
where
  data_classes ?| array['Usernames','Passwords'];
```