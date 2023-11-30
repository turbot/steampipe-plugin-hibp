---
title: "Steampipe Table: hibp_paste - Query HIBP Pastes using SQL"
description: "Allows users to query Pastes in HIBP, specifically the paste details including the source, id, title, author, date, email count, and the full URL of the paste."
---

# Table: hibp_paste - Query HIBP Pastes using SQL

HIBP (Have I Been Pwned) is a service that allows users to check whether their personal data has been compromised by data breaches. The service collects and analyzes hundreds of database dumps and pastes containing information about billions of leaked accounts. A Paste is information that has been "pasted" to a publicly facing website designed to share content, such as Pastebin.

## Table Usage Guide

The `hibp_paste` table provides insights into pastes within HIBP. As a security analyst, explore paste-specific details through this table, including the source, id, title, author, date, email count, and the full URL of the paste. Utilize it to uncover information about pastes, such as those containing compromised personal data, the sources of these pastes, and the extent of personal data leaks.

## Examples

### List pastes where `billy@example.com` was included in the paste
Explore pastes where a specific email address was included, helping to identify potential data breaches or unauthorized sharing of information.

```sql
select
  id,
  source
from
  hibp_paste
where
  account = 'billy@example.com';
```

### List pastes where `billy@example.com` was included in the last 10 years
Discover the instances where the email 'billy@example.com' has been involved in any data breaches within the past decade. This is useful for understanding the security history of this specific email address.

```sql
select
  id,
  source,
  date
from
  hibp_paste
where
  account = 'billy@example.com'
  and date > now() - interval '10 years';
```