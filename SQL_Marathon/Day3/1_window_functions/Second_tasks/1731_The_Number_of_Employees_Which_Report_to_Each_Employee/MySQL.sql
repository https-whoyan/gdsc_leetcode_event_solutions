SELECT
    tb1.employee_id,
    tb1.name, 
    COUNT(tb2.reports_to) as reports_count, 
    ROUND(
        AVG(tb2.age), 0
    ) as average_age 
FROM Employees tb1
JOIN (
    SELECT age, reports_to
    FROM Employees
) tb2 
ON tb1.employee_id = tb2.reports_to 
GROUP BY tb1.employee_id
ORDER BY tb1.employee_id