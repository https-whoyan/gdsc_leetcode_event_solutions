SELECT tb1.name
FROM employee tb1
LEFT JOIN employee tb2 
ON tb2.managerId = tb1.id
GROUP BY tb1.id
HAVING COUNT(tb2.id) >= 5;