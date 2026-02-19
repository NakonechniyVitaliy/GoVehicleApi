CREATE TABLE brands (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        marka_id INTEGER UNIQUE NOT NULL,
                        category_id INTEGER NOT NULL,
                        cnt INTEGER NOT NULL,
                        country_id INTEGER NOT NULL,
                        eng TEXT NOT NULL,
                        name TEXT NOT NULL,
                        slang TEXT NOT NULL,
                        value INTEGER NOT NULL
);