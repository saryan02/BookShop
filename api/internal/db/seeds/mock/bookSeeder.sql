INSERT INTO books (id, title, author, price) VALUES
                                             (1,'The Go Programming Language', 'Alan Donovan', 29.99),
                                             (2,'Clean Code', 'Robert Martin', 24.99),
                                             (3, 'Domain-Driven Design', 'Eric Evans', 34.99)
    ON CONFLICT (id) DO NOTHING;