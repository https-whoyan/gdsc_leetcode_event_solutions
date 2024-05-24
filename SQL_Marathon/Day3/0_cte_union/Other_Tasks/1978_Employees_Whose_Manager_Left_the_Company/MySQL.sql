SELECT e1.employee_id
FROM Employees e1
LEFT JOIN Employees e2
ON e1.manager_id = e2.employee_id
WHERE e1.salary < 30000 
AND e2.employee_id IS NULL 
AND e1.manager_id IS NOT NULL
ORDER BY employee_id;