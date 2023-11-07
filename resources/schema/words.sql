CREATE TABLE words (
                       word   TEXT PRIMARY KEY NOT NULL,
                       length INTEGER NOT NULL
);


CREATE INDEX word_len_idx ON words (length);

