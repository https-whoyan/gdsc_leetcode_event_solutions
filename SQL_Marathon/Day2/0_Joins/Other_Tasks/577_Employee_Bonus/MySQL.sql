SELECT 
    tb1.name, 
    tb2.bonus
FROM employee tb1
LEFT JOIN bonus tb2
ON tb1.empId = tb2.empId
WHERE IFNULL(tb2.bonus, 0) < 1000