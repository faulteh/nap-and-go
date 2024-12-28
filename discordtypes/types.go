// Package discordtypes provides the types used in the Discord API.
// Used by db and discordapi packages but kept separate to avoid circular dependencies.
package discordtypes

// Guild represents a Discord guild/server
type Guild struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Icon           string `json:"icon"`
	Owner          bool   `json:"owner"`
	Permissions    int64  `json:"permissions"`
	PermissionsNew string `json:"permissions_new"`
	HasBot         bool
}
