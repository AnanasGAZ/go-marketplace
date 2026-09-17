package main

func GetUserHistory(
	history map[int][]string, userID int,
) []string {
	userHistory, exists := history[userID]
	if !exists {
		return nil
	}
	historyCopy := make([]string, len(userHistory))
	copy(historyCopy, userHistory) // копируем строки
	return historyCopy
}
