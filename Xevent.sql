
/*
http://dba.stackexchange.com/questions/47025/identifying-unused-stored-procedures

If you are on SQL Server 2008+ you can also use extended events with a histogram target. 
Possibly this would be more light weight than a trace.

AFAIK you would need to create a different session for each database of interest though as 
I couldn't see any indication that bucketizing on multiple columns was possible. 

*/
DECLARE @db_id int;  
SET @db_id = DB_ID(N'DBnew');   

CREATE EVENT SESSION [count_module_start_dbnew]
ON SERVER
ADD EVENT sqlserver.module_start
(  
        WHERE (source_database_id = @db_id) 
)
ADD TARGET package0.asynchronous_bucketizer(
     SET  filtering_event_name='sqlserver.module_start', 
            source_type=0, 
            source='object_id',
            slots = 10000
)
WITH (MAX_DISPATCH_LATENCY = 5 SECONDS)
GO




--Start Xevent
ALTER EVENT SESSION [count_module_start_dbnew]
ON SERVER
STATE=START

/*
And then after running some stored procedures in that DB a few times and retrieving the data with
*/

SELECT CAST(target_data as XML) target_data
FROM sys.dm_xe_sessions AS s 
JOIN sys.dm_xe_session_targets t
    ON s.address = t.event_session_address
WHERE s.name = 'count_module_start_dbnew'
/*
The output is

<HistogramTarget truncated="0" buckets="16384">
  <Slot count="36">
    <value>1287675635</value>
  </Slot>
  <Slot count="3">
    <value>1271675578</value>
  </Slot>
  <Slot count="2">
    <value>1255675521</value>
  </Slot>
</HistogramTarget>

Showing that the procedure with object_id of 1287675635 was executed 36 times for example. 
The asynchronous_bucketizer is memory only so it would be best to set up something that polls this every so often and saves to persistent storage.
*/
