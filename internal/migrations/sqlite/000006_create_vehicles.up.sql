CREATE TABLE vehicles (
                          id INTEGER PRIMARY KEY AUTOINCREMENT,
                          brand INTEGER NOT NULL,
                          driver_type INTEGER NOT NULL,
                          gearbox INTEGER NOT NULL,
                          body_style INTEGER NOT NULL,
                          category INTEGER NOT NULL,
                          mileage INTEGER,
                          model TEXT,
                          price INTEGER NOT NULL,

                          FOREIGN KEY (brand)
                              REFERENCES brands(id)
                              ON DELETE RESTRICT,

                          FOREIGN KEY (driver_type)
                              REFERENCES driver_types(id)
                              ON DELETE RESTRICT,

                          FOREIGN KEY (gearbox)
                              REFERENCES gearboxes(id)
                              ON DELETE RESTRICT,

                          FOREIGN KEY (body_style)
                              REFERENCES body_styles(id)
                              ON DELETE RESTRICT,

                          FOREIGN KEY (category)
                              REFERENCES categories(id)
                              ON DELETE RESTRICT
);