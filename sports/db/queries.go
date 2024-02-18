package db

const (
	racesList   = "list"
	getRaceById = "getRaceById" // TODO unsure about the naming of this
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races
		`,
		getRaceById: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time 
			FROM races WHERE id = ?
		`,
	}
}
