SELECT tb1.name
FROM Employee tb1
LEFT JOIN Employee tb2
ON tb2.managerId = tb1.id 
GROUP BY tb1.id, tb1.name
HAVING COUNT(DISTINCT tb2.id) >= 5