CREATE TABLE driver_types (
                              id INTEGER PRIMARY KEY AUTOINCREMENT,
                              name TEXT NOT NULL,
                              value INTEGER UNIQUE NOT NULL
);