// SQL queries
package storage

const (
	GET_SONG_QUERY = `
		SELECT song_name,group_name,release_date,lyrics,link 
		FROM songs 
		WHERE group_name=$1 AND song_name=$2;
	`

	ADD_SONG_QUERY = `
		INSERT INTO songs 
		(group_name, song_name, release_date, lyrics, link) 
		VALUES ($1,$2,$3,$4,$5);
	`

	EDIT_SONG_QUERY = `UPDATE songs 
		SET release_date = $3, 
		lyrics = $4,
		link = $5      
		WHERE group_name = $1 AND song_name = $2;
	`

	DEL_SONG_QUERY = `
		DELETE FROM songs
		WHERE group_name=$1 AND song_name=$2;
	`
)
