CREATE TABLE `profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` text,
  `gender` varchar(10) NOT NULL,
  `martial_status` varchar(50) NOT NULL,
  `phone` varchar(15) NOT NULL, 
  PRIMARY KEY (`id`)
);


INSERT INTO `profile` (`name`, `email`, `password`, `gender`, `martial_status`, `phone`) VALUES
('Test1', 'test1@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'male', 'single', '9876543210'),
('Test2', 'test2@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'female', 'married', '9876533210'),
('Test3', 'test3@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'female', 'married', '9876563210'),
('Test4', 'test4@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'female', 'married', '9876583210'),
('Test5', 'test5@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'female', 'married', '9876523210'),
('Test6', 'test6@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'male', 'single', '9876543910'),
('Test7', 'test7@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'male', 'single', '9876543310'),
('Test8', 'test8@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'male', 'single', '9826543210'),
('Test9', 'test9@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'male', 'single', '9826543250'),
('Test10', 'test10@gmail.com', '$2a$10$Sz0GkK/GlUcmX9EZLJVXNOEXqyaCDICfpjF2NX6Pn7YR3wrLky6X2', 'male', 'single', '9821543250');




CREATE TABLE `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `profile_id` int(11) NOT NULL,
  `auth_provider` varchar(100) NOT NULL,
  `auth_id` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `profile_id` (`profile_id`)
);


INSERT INTO `auth` (`profile_id`, `auth_provider`, `auth_id`, `name`) VALUES
((Select id from `profile` WHERE name = 'Test1'), 'linkedin', '1234567890', 'Test1'),
((Select id from `profile` WHERE name = 'Test1'), 'facebook', '987654321124345', 'Testfacebook'),
((Select id from `profile` WHERE name = 'Test2'), 'linkedin', '12345678903434', 'Test2'),
((Select id from `profile` WHERE name = 'Test2'), 'github', '987654321124345', 'Test2github'),
((Select id from `profile` WHERE name = 'Test3'), 'linkedin', '1234569747890', 'Test3'),
((Select id from `profile` WHERE name = 'Test3'), 'facebook', '987654321124345', 'Test3facebook'),
((Select id from `profile` WHERE name = 'Test4'), 'linkedin', '123456546537890', 'Test4'),
((Select id from `profile` WHERE name = 'Test4'), 'facebook', '987654325451124345', 'Test4acebook'),
((Select id from `profile` WHERE name = 'Test5'), 'linkedin', '12378684567890', 'Test5'),
((Select id from `profile` WHERE name = 'Test5'), 'facebook', '98765435621124345', 'Test5facebook'),
((Select id from `profile` WHERE name = 'Test6'), 'linkedin', '123456789039798434', 'Test6'),
((Select id from `profile` WHERE name = 'Test6'), 'github', '9876543211243455677', 'Test6github'),
((Select id from `profile` WHERE name = 'Test7'), 'linkedin', '123456778998903434', 'Test7'),
((Select id from `profile` WHERE name = 'Test7'), 'github', '987654321124345', 'Test7github'),
((Select id from `profile` WHERE name = 'Test8'), 'linkedin', '12345678903476834', 'Test8'),
((Select id from `profile` WHERE name = 'Test8'), 'github', '987654321124379745', 'Test8github'),
((Select id from `profile` WHERE name = 'Test9'), 'linkedin', '1234567890343564', 'Test9'),
((Select id from `profile` WHERE name = 'Test9'), 'github', '98765432112434985', 'Test9github'),
((Select id from `profile` WHERE name = 'Test10'), 'linkedin', '1234567890343234', 'Test10'),
((Select id from `profile` WHERE name = 'Test10'), 'github', '98765432112434455', 'Test10github');

ALTER TABLE `auth`
  ADD CONSTRAINT `auth_ibfk_1` FOREIGN KEY (`profile_id`) REFERENCES `profile` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;




CREATE TABLE `access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ;


INSERT INTO `access` (`name`) VALUES
('Module1'),
('Module2'),
('Module3'),
('Module4'),
('Module5');



CREATE TABLE `access_profile_mapping` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `profile_id` int(11) NOT NULL,
  `access_id` int(11) NOT NULL,
  `view` smallint(6) NOT NULL,
  `edit` smallint(6) NOT NULL,
  `remove` smallint(6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `profile_id` (`profile_id`),
  KEY `access_id` (`access_id`)
);


ALTER TABLE `access_profile_mapping`
  ADD CONSTRAINT `access_profile_mapping_ibfk_1` FOREIGN KEY (`profile_id`) REFERENCES `profile` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `access_profile_mapping_ibfk_2` FOREIGN KEY (`access_id`) REFERENCES `access` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;


INSERT INTO `access_profile_mapping` (`profile_id`, `access_id`, `view`, `edit`, `remove`) VALUES
((Select id from `profile` WHERE name = 'Test1'), (Select id from `access` WHERE name = 'Module1'), '1', '0', '0'), 
((Select id from `profile` WHERE name = 'Test2'), (Select id from `access` WHERE name = 'Module1'), '1', '1', '1'),
((Select id from `profile` WHERE name = 'Test3'), (Select id from `access` WHERE name = 'Module2'), '1', '1', '1'),
((Select id from `profile` WHERE name = 'Test4'), (Select id from `access` WHERE name = 'Module2'), '1', '0', '0'),
((Select id from `profile` WHERE name = 'Test5'), (Select id from `access` WHERE name = 'Module3'), '1', '1', '0'),
((Select id from `profile` WHERE name = 'Test6'), (Select id from `access` WHERE name = 'Module3'), '1', '1', '1'),
((Select id from `profile` WHERE name = 'Test7'), (Select id from `access` WHERE name = 'Module4'), '1', '1', '0'),
((Select id from `profile` WHERE name = 'Test8'), (Select id from `access` WHERE name = 'Module4'), '1', '1', '1'),
((Select id from `profile` WHERE name = 'Test9'), (Select id from `access` WHERE name = 'Module5'), '1', '1', '1');