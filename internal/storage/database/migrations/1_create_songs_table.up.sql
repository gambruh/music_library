CREATE TABLE artists (
    id SERIAL PRIMARY KEY,
    artist_name VARCHAR(255) NOT NULL,
);

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_name VARCHAR(255) NOT NULL,
    release_date DATE,
    lyrics TEXT,
    link TEXT,
    FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE
);