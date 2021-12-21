select table_schema,
       table_name,
       collation_name,
       ordinal_position,
       column_default,
       is_nullable,
       data_type,
       character_maximum_length,
       numeric_precision,
       numeric_scale,
       character_set_name,
       collation_name,
       column_type,
       column_key,
       column_comment
from information_schema.columns
where table_schema not in ('mysql', 'information_schema', 'performance_schema', 'sys')