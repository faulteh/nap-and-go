// Initializes and runs the Discord bot and connects to the database.
package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"

    "github.com/bwmarrin/discordgo"

    "github.com/faulteh/nap-and-go/db"
    "github.com/faulteh/nap-and-go/models"
)

func main() {
    // Initialize the database connection
    db.Connect()
    // Get the database connection
    conn := db.GetDB()

    // AutoMigrate the models
    log.Println("Auto-migrating models")
    err := conn.AutoMigrate(&models.Server{}, &models.User{})
    if err != nil {
        log.Fatalf("failed to auto-migrate models: %v", err)
        return
    }

	// Load the bot token from an environment variable
	token := os.Getenv("DISCORD_BOT_TOKEN")
	if token == "" {
		log.Println("Please set the DISCORD_BOT_TOKEN environment variable")
		return
	}

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Error creating Discord session:", err)
		return
	}

	// Register the message handler
	dg.AddHandler(messageCreate)

	// Open a connection to Discord
	err = dg.Open()
	if err != nil {
		log.Println("Error opening connection:", err)
		return
	}
	log.Println("Bot is now running. Press CTRL+C to exit.")

	// Wait for a signal to exit
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Cleanly close the Discord session
	err = dg.Close()
    if err != nil {
        log.Println("Error closing Discord session:", err)
    }
    log.Println("Bot has been stopped.")
}

// messageCreate is called whenever a new message is created in a channel the bot has access to
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Respond to "!ping"
	if m.Content == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
        if err != nil {
            log.Println("Error sending message:", err)
        }
	}
}
