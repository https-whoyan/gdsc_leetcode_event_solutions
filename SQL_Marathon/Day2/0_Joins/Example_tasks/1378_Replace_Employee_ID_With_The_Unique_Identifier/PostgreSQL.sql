SELECT
    employeeUNI.unique_id AS unique_id,
    employees.name
FROM employees
LEFT JOIN employeeUNI ON employees.id=employeeUNI.id