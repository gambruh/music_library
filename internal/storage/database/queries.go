// SQL queries
package storage

const (
	GET_SONG_QUERY = `
		SELECT 
			s.song_name, g.group_name, s.release_date, s.lyrics, s.link
		FROM 
			songs s
		JOIN 
			groups g ON s.group_id = g.id
		WHERE 
			g.group_name = $1 AND s.song_name = $2;
	`

	ADD_GROUP_QUERY = `
		INSERT INTO groups (group_name)
		SELECT $1::varchar
		WHERE NOT EXISTS (
			SELECT 1 FROM groups WHERE group_name = $1::varchar
		);
	`

	GET_GROUP_ID_QUERY = `
		SELECT id FROM groups WHERE group_name = $1;
	`

	ADD_SONG_QUERY = `
		INSERT INTO songs 
		(group_id, song_name, release_date, lyrics, link) 
		VALUES (
			(SELECT id FROM groups WHERE group_name = $1),
			$2, $3, $4, $5
		);
	`

	EDIT_SONG_QUERY = `
		UPDATE songs s
		SET 
			release_date = $3, 
			lyrics = $4,
			link = $5
		FROM 
			groups g
		WHERE 
			s.group_id = g.id AND g.group_name = $1 AND s.song_name = $2;
	`

	DEL_SONG_QUERY = `
	DELETE FROM songs
	USING groups
	WHERE 
		songs.group_id = groups.id AND
		groups.group_name = $1 AND
		songs.song_name = $2;
`
)
