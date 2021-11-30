select schema_name,
       default_character_set_name,
       default_collation_name
from information_schema.schemata
where schema_name not in ('mysql', 'information_schema', 'performance_schema', 'sys')