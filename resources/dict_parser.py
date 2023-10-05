# readline
# parse string
# ignore proper nouns: ignore if first character is capital
# ignore strings with numbers or special characters
# so to invert only words with lowercase characters allowed, and length > 3
# generate a sql statement and insert into a db
#
import os

import psycopg2


def include_word(word):
    return word.islower() and word.isalpha() and len(word) < 26


def persist(word):
    cursor.execute("insert into words (word, length) VALUES (%s, %s) ON CONFLICT  do nothing", (word, str(len(word))))
    db_handle.commit()


def process():
    with open("raw_dict.txt") as f:
        while True:
            word = f.readline()
            if word is not None:
                word = word.strip()
                if include_word(word):
                    persist(word)
            else:
                break


def init_db():
    dbname = os.environ.get("wordle_pg_db")
    user = os.environ.get("wordle_pg_user")
    password = os.environ.get("wordle_pg_password")
    connection = psycopg2.connect("dbname={} user={} password={}".format(dbname, user, password))
    return connection


# wordle_pg_db=wordle wordle_pg_user=postgres word_pg_password=password
if __name__ == "__main__":
    db_handle = init_db()
    cursor = db_handle.cursor()
    process()
    cursor.close()
    db_handle.close()
