package model



type PartyMember struct {
	UserId string `json:"user_id"`
	PartyId string `json:"party_id"`
	DateJoined int `json:"date_joined"`
	TotalSongsPlayed int `json:"total_songs_played"`
	TotalMusicListeningTime int `json:"total_music_listening_time"`
	Roles []string `json:"roles"`
}

