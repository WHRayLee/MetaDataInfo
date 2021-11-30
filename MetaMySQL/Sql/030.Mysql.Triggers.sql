select trigger_schema,
       trigger_name,
       event_manipulation,
       event_object_schema,
       event_object_table,
       action_order,
       action_condition,
       action_statement,
       action_orientation,
       action_timing,
       created,
       character_set_client,
       database_collation
from information_schema.triggers
where trigger_schema not in ('sys')