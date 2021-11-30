select table_schema,
       table_name,
       view_definition
from information_schema.views
where table_schema not in ('sys')