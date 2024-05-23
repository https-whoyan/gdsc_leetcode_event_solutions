SELECT 
    employeeUNI.unique_id as unique_id, 
    employees.name
FROM employees
LEFT JOIN employeeUNI ON employees.id=employeeUNI.id