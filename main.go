package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"user-activity/events"
)

func main() {
	username := ""
	for {
		fmt.Print("github activity ")
		fmt.Scan(&username)
		if username == "exit" {
			break
		}
		events, err := getActivity(username)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		printActivity(events)
	}

}

func getActivity(username string) ([]events.GitHubEvent, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s/events", username), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "activity-go-app")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var events []events.GitHubEvent
	if err := json.Unmarshal(body, &events); err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return nil, errors.New("активности не найдено")
	}

	return events, nil
}

func printActivity(reqEevents []events.GitHubEvent) {
	var limitedEvents []events.GitHubEvent

	if len(reqEevents) > 10 {
		limitedEvents = reqEevents[:10]
	} else {
		limitedEvents = reqEevents
	}

	for _, event := range limitedEvents {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed to %s\n", event.Repo.Name)

		case "IssuesEvent":
			if event.Payload.Action == "opened" {
				fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
			} else {
				fmt.Printf("- %s an issue in %s\n", event.Payload.Action, event.Repo.Name)
			}

		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)

		default:
			fmt.Printf("- %s in %s\n", event.Type, event.Repo.Name)
		}

	}
	fmt.Println("--- last 10 actions activity")
}
