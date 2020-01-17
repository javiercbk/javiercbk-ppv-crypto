-- pass is 1234...yeah, see you in court.
INSERT INTO users (email, first_name, last_name, password, registration_date, created_at)
VALUES (
  'john@doe.com', 'John', 'Doe', '$2y$12$7TCLLtmM4WxHptX418/WSOnL6cYyn5w2xpcYN6gS4ZQbfZZWfClJa', NOW(), NOW()
);

INSERT INTO permissions_users ("user_id", "resource", "access") VALUES (1, 'event', 'write'), (1, 'user', 'write');

INSERT INTO users (email, first_name, last_name, password, registration_date, created_at)
VALUES (
  'jane@doe.com', 'Jane', 'Doe', '$2y$12$7TCLLtmM4WxHptX418/WSOnL6cYyn5w2xpcYN6gS4ZQbfZZWfClJa', NOW(), NOW()
);

INSERT INTO permissions_users ("user_id", "resource", "access") VALUES (2, 'event', 'read');