CREATE TABLE game_log (
                          game_id       INTEGER NOT NULL,
                          attempted_word TEXT NOT NULL,
                          attempt_num    INTEGER NOT NULL,
                          FOREIGN KEY (game_id) REFERENCES game (id)
);
