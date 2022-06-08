# Table: hibp_paste

A "paste" is information that has been "pasted" to a publicly facing website designed to share content such as Pastebin. These services are favoured by hackers due to the ease of anonymously sharing information and they're frequently the first place a breach appears.

HIBP searches through pastes that are broadcast by the accounts in the Paste Sources Twitter list and reported as having emails that are a potential indicator of a breach. Finding an email address in a paste does not immediately mean it has been disclosed as the result of a breach. Review the paste and determine if your account has been compromised then take appropriate action such as changing passwords.

## Examples

### Pastes where `billy@example.com` was included in the paste

```sql
select
  id,
  source
from
  hibp_paste
where
  account = 'billy@example.com'
```
