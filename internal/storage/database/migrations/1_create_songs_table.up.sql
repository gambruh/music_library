CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL,
    song_name VARCHAR(255) NOT NULL,
    release_date DATE,
    lyrics TEXT,
    link TEXT,
    FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
    CONSTRAINT unique_song_group UNIQUE (song_name, group_id)
);