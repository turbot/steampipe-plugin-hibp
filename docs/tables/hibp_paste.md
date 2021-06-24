# Table: hibp_paste

Pastes indexed by HIBP.

## Examples

### Pastes where test@test.org was included in the paste

```sql
select * from hibp_paste
 where account = 'test@test.org'
```
