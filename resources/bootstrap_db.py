import getpass
import json
import os
import sqlite3

config = {}
config_path = "config/config.json"
tables = ['words', 'attempts', 'game', 'game_log']


def include_word(word):
    return word.islower() and word.isalpha() and len(word) < 26


def persist(word):
    cursor.execute("insert into words (word, length) VALUES (?, ?) ON CONFLICT  do nothing", (word, str(len(word))))
    db_handle.commit()


def process():
    with open("resources/raw_dict.txt") as f:
        while True:
            word = f.readline()
            if word is not None and len(word) > 0:
                word = word.strip()
                if include_word(word):
                    persist(word)
            else:
                # print("breaking out")
                break
    for i in range(5, 25):
        cursor.execute("insert into attempts (word_len, max_attempts) VALUES (?, ?) ON CONFLICT  do nothing",
                       (i, i + 1))
        db_handle.commit()


def load_config():
    with open(config_path, 'r') as f:
        data = f.read()
        global config
        config = json.loads(data)


def create_tables():
    for table in tables:
        res = cursor.execute("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='{}';".format(table))
        result = res.fetchone()
        r = int(result[0])
        if r == 0:
            with open("resources/schema/{}.sql".format(table), 'r') as f:
                data = f.read()
                try:
                    cursor.executescript(data)
                    db_handle.commit()
                except sqlite3.OperationalError:
                    pass


def init_db():
    current_user = getpass.getuser()
    dbPathParameterized = config.get("SQLITE_DB_PATH")
    if dbPathParameterized == "" or dbPathParameterized is None:
        print("database path not provided")
        exit(1)

    db_path = dbPathParameterized.replace("${current_user}", current_user)

    if not os.path.exists(db_path):
        directory = os.path.dirname(db_path)
        if not os.path.exists(directory):
            os.makedirs(directory, exist_ok=True)
    connection = sqlite3.connect(db_path)
    return connection


if __name__ == "__main__":
    load_config()
    db_handle = init_db()
    cursor = db_handle.cursor()
    create_tables()
    process()
    cursor.close()
    db_handle.close()
