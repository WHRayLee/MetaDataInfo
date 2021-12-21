select table_schema
     , table_name
     , constraint_name
     , constraint_schema
     , column_name
     , ordinal_position
     , position_in_unique_constraint
     , referenced_table_schema
     , referenced_table_name
     , referenced_column_name
from information_schema.key_column_usage
where table_schema not in ('mysql', 'information_schema', 'performance_schema', 'sys')