INSERT INTO
    roles (name)
VALUES
    ('Admin'),
    ('User')
ON CONFLICT DO NOTHING;

INSERT INTO
    users (name, email, password_hash, role_id)
SELECT
    'Eleazar Fig',
    'eleazar.fig@example.com',
    '$2b$12$oSBDBlelP5C7nqokSfngTeSy9bCkfnLOfC3pRb7pxaRbSKSHk0q3S',
    -- '$2b$12$GPf.eB7OpIcB3hpCr/JhoOOVPHQ0YE9oLnDA0KyHq7oGBvAFospLK',
    role_id
FROM
    roles
WHERE
    name LIKE 'Admin';
