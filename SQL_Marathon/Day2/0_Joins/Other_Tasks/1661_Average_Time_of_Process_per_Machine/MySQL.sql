SELECT 
    tb1.machine_id, 
    ROUND(
        AVG(
            tb2.timestamp - tb1.timestamp
        )
    , 3) as processing_time
FROM Activity tb1
JOIN Activity tb2
ON tb1.machine_id = tb2.machine_id AND 
tb1.process_id = tb2.process_id AND 
tb1.activity_type = "start" AND
tb2.activity_type = "end"
GROUP BY tb1.machine_id