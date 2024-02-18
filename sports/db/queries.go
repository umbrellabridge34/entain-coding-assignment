package db

const (
	sportEventsList = "list"
)

func getSportQueries() map[string]string {
	return map[string]string{
		sportEventsList: `
			SELECT 
				id, 
				name, 
				advertised_start_time 
			FROM sportEvents
		`,
	}
}
