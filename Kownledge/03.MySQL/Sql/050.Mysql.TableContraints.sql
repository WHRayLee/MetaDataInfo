select constraint_schema,
       constraint_name,
       table_name,
       constraint_type
from information_schema.table_constraints
where constraint_schema not in ('mysql', 'information_schema', 'performance_schema', 'sys')