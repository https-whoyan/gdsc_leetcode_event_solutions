SELECT 
    tb2.name as Department, 
    tb1.name as Employee, 
    tb1.salary as Salary
FROM (
    SELECT 
        tmp.salary, 
        tmp.name, 
        tmp.departmentId, 
        DENSE_RANK() OVER (
            PARTITION BY departmentId
            ORDER BY salary DESC
        ) as RNum
    FROM Employee tmp
) tb1
JOIN Department tb2 
ON tb1.departmentId = tb2.id
WHERE tb1.RNum <= 3