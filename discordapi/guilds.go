// Package discordapi interacts with the Discord API to fetch server information and provide data types
// to represent the data.
package discordapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/faulteh/nap-and-go/config"
	"github.com/faulteh/nap-and-go/db"
	"github.com/faulteh/nap-and-go/discordtypes"
)

// UserAdminGuildList returns a list of servers the user has admin permissions in queried from the Discord API
func UserAdminGuildList(token *oauth2.Token) ([]discordtypes.Guild, error) {
	// Retrieve list of servers for user from discord
	oauth2Config := config.LoadDiscordConfig().OAuth2Config()
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://discord.com/api/users/@me/guilds")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var guilds []discordtypes.Guild
	if err := json.NewDecoder(resp.Body).Decode(&guilds); err != nil {
		return nil, err
	}

	// Filter guilds where the user has admin permissions
	var adminGuilds []discordtypes.Guild
	const ADMINISTRATOR = 0x00000008
	for _, guild := range guilds {
		if guild.Permissions&ADMINISTRATOR != 0 {
			adminGuilds = append(adminGuilds, guild)
		}
	}

	return adminGuilds, nil
}

// BotGuildList returns a list of servers the bot is in queried from the Discord API
// and syncs the database with the list of servers
func BotGuildList() ([]discordtypes.Guild, error) {
	// Bot uses a simple Authorization: Bot <token> header
	token := config.LoadDiscordConfig().BotToken
	if token == "" {
		return nil, fmt.Errorf("missing bot token")
	}
	// Use the bot authorization token to request the guild list
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://discord.com/api/users/@me/guilds", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bot "+token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var guilds []discordtypes.Guild
	if err := json.NewDecoder(resp.Body).Decode(&guilds); err != nil {
		return nil, err
	}

	// Sync the guilds with the database
	err = db.SyncGuilds(guilds)
	if err != nil {
		return guilds, err
	}

	return guilds, nil
}
