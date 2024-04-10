BEGIN;

-- seed profiles
INSERT INTO profiles (name, type_user)
SELECT 'ADMIN', 'ADMIN'
WHERE NOT EXISTS (SELECT 1 FROM profiles WHERE name = 'ADMIN');

INSERT INTO profiles (name, type_user)
SELECT 'ADMIN_SCHOOL', 'ADMIN_SCHOOL'
WHERE NOT EXISTS (SELECT 1 FROM profiles WHERE name = 'ADMIN_SCHOOL');

INSERT INTO profiles (name, type_user)
SELECT 'TEACHER', 'TEACHER'
WHERE NOT EXISTS (SELECT 1 FROM profiles WHERE name = 'TEACHER');

INSERT INTO profiles (name, type_user)
SELECT 'STUDENT', 'STUDENT'
WHERE NOT EXISTS (SELECT 1 FROM profiles WHERE name = 'STUDENT');

-- seed profile_permissions
INSERT INTO profile_permissions (profile_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM profiles WHERE name = 'ADMIN'), 15, 15, 15, 15
WHERE NOT EXISTS (SELECT 1 FROM profile_permissions WHERE profile_id = (SELECT id FROM profiles WHERE name = 'ADMIN'));

INSERT INTO profile_permissions (profile_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM profiles WHERE name = 'ADMIN_SCHOOL'), 15, 15, 0 , 15
WHERE NOT EXISTS (SELECT 1 FROM profile_permissions WHERE profile_id = (SELECT id FROM profiles WHERE name = 'ADMIN_SCHOOL'));

INSERT INTO profile_permissions (profile_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM profiles WHERE name = 'TEACHER'),0, 15, 0 , 15
WHERE NOT EXISTS (SELECT 1 FROM profile_permissions WHERE profile_id = (SELECT id FROM profiles WHERE name = 'TEACHER'));

INSERT INTO profile_permissions (profile_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM profiles WHERE name = 'STUDENT'),0, 15,0 , 15
WHERE NOT EXISTS (SELECT 1 FROM profile_permissions WHERE profile_id = (SELECT id FROM profiles WHERE name = 'STUDENT'));

-- seed user
INSERT INTO users (name, type_user, email, password, profile_id)
SELECT 'ADMIN', 'ADMIN', 'admin@admin.com','$2a$12$B7EW8DXW3mFGJi10Rcceo.YfXW2Rm0eEqCc9ShXWq4kONNp8vhZqy', (SELECT id FROM profiles WHERE name = 'ADMIN')
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@admin.com');

INSERT INTO users (name, type_user, email, password, profile_id)
SELECT 'ADMIN_SCHOOL', 'ADMIN_SCHOOL', 'admin@admin_school.com','$2a$12$B7EW8DXW3mFGJi10Rcceo.YfXW2Rm0eEqCc9ShXWq4kONNp8vhZqy', (SELECT id FROM profiles WHERE name = 'ADMIN_SCHOOL')
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@admin_school.com');

INSERT INTO users (name, type_user, email, password, profile_id)
SELECT 'TEACHER', 'TEACHER', 'teacher@dev.com','$2a$12$B7EW8DXW3mFGJi10Rcceo.YfXW2Rm0eEqCc9ShXWq4kONNp8vhZqy', (SELECT id FROM profiles WHERE name = 'TEACHER')
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'teacher@dev.com');

-- seed user_permissions
INSERT INTO user_permissions (user_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM users WHERE email = 'admin@admin.com'), 15, 15,15,15
WHERE NOT EXISTS (SELECT 1 FROM user_permissions WHERE user_id = (SELECT id FROM users WHERE email = 'admin@admin.com'));

INSERT INTO user_permissions (user_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM users WHERE email = 'admin@admin_school.com'), 15, 15, 0 , 15
WHERE NOT EXISTS (SELECT 1 FROM user_permissions WHERE user_id = (SELECT id FROM users WHERE email = 'admin@admin_school.com'));

INSERT INTO user_permissions (user_id, users, classes, profiles, lessons)
SELECT (SELECT id FROM users WHERE email = 'teacher@dev.com'), 0, 15, 0 , 15
WHERE NOT EXISTS (SELECT 1 FROM user_permissions WHERE user_id = (SELECT id FROM users WHERE email = 'teacher@dev.com'));

COMMIT;