select table_schema,
       table_name,
       partition_name,
       subpartition_name,
       partition_ordinal_position,
       subpartition_ordinal_position,
       partition_method,
       subpartition_method,
       partition_expression,
       subpartition_expression,
       partition_description,
       create_time,
       update_time,
       partition_comment,
       tablespace_name
from information_schema.partitions
where table_schema not in ('mysql', 'information_schema', 'performance_schema', 'sys')