SELECT *
FROM Users
WHERE regexp_like(mail, '^[a-zA-Z][a-zA-Z0-9_.-]*@leetcode[.]com$')