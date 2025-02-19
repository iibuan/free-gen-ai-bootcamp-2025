CREATE TABLE words (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bahasa_indonesia TEXT,
    english TEXT
);

CREATE TABLE groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT
);

CREATE TABLE words_groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id INTEGER,
    group_id INTEGER,
    FOREIGN KEY(word_id) REFERENCES words(id),
    FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE study_sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER,
    created_at DATETIME,
    study_activity_id INTEGER,
    FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE study_activities (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    study_session_id INTEGER,
    group_id INTEGER,
    created_at DATETIME,
    FOREIGN KEY(study_session_id) REFERENCES study_sessions(id),
    FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE word_review_items (
    word_id INTEGER,
    study_session_id INTEGER,
    correct BOOLEAN,
    created_at DATETIME,
    FOREIGN KEY(word_id) REFERENCES words(id),
    FOREIGN KEY(study_session_id) REFERENCES study_sessions(id)
);