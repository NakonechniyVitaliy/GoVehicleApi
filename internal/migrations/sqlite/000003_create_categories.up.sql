CREATE TABLE categories (
                                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                                    name TEXT NOT NULL,
                                    value INTEGER UNIQUE NOT NULL
);