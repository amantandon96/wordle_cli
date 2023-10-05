CREATE TABLE game (
                      id         INTEGER PRIMARY KEY AUTOINCREMENT,
                      word       TEXT NOT NULL,
                      mode       TEXT,
                      state      TEXT,
                      created_at DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
