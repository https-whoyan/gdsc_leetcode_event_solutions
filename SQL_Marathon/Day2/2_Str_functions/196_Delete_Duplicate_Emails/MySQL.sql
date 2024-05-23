DELETE R1 
FROM Person R1, Person R2
WHERE R1.email = R2.email AND
R1.id > R2.id