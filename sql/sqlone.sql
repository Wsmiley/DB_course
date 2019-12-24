USE DB_course
SELECT Name from sys.Databases
SELECT * FROM INFORMATION_SCHEMA.TABLES
-- drop table students
SELECT * FROM admins
SELECT * FROM students
SELECT * FROM scores
SELECT * FROM courses
SELECT * FROM classes
SELECT * FROM student_counts
-- INSERT INTO DB_course.dbo.admins(Username, Password,Name) VALUES ('20190101','123456',N'李老师')
-- INSERT INTO DB_course.dbo.student_counts(username,password,name) VALUES ('202170109','qqcc7711',N'陈威')
-- INSERT INTO DB_course.dbo.student_counts(username,password,name) VALUES ('202170108','12345',N'陈涛')