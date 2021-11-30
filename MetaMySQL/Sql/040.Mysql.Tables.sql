select table_schema,
       table_name,
       table_type,
       engine,
       table_collation
from information_schema.tables
where table_schema not in ('mysql', 'information_schema', 'performance_schema', 'sys')