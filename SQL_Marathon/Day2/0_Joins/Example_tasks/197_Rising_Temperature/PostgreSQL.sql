SELECT data1.id
FROM weather data1, 
weather data2
WHERE data1.temperature > data2.temperature 
AND data1.recordDate - data2.recordDate = 1