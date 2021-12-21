select routine_schema,
       routine_name,
       routine_type,
       data_type,
       dtd_identifier,
       collation_name,
       routine_body,
       routine_definition,
       created,
       last_altered,
       character_set_client,
       collation_connection,
       database_collation
from information_schema.routines
where routine_schema not in ('mysql', 'information_schema', 'performance_schema', 'sys')