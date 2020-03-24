use diaverse;
​
​
INSERT INTO users (username, password, email, first_name, last_name, user_bio) values ('harrison', 'password', 'affelh@wit.edu', 'harrison', 'affel', 'hardwarestuff');
INSERT INTO users (username, password, email, first_name, last_name, user_bio) values ('william', 'password', 'medeirosp@wit.edu', 'paul', 'medeiros', 'uistuff');
INSERT INTO users (username, password, email, first_name, last_name, user_bio) values ('william', 'diaverse!123', 'shuew@wit.edu', 'william', 'shue', 'serverstuff');
​
​
INSERT INTO hardwareToken (username, hardwareToken) values ('harrison', 'token1');
INSERT INTO hardwareToken (username, hardwareToken) values ('william', 'token2');
INSERT INTO hardwareToken (username, hardwareToken) values ('paul', 'token3');
​
INSERT INTO scripts(script_id, hardwareToken, script) values (1, 'token1', 'json1');
INSERT INTO scripts(script_id, hardwareToken, script) values (2, 'token3', 'json3');
INSERT INTO scripts(script_id, hardwareToken, script) values (3, 'token2', 'json2');
