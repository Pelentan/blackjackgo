package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/Pelentan/blackjackgo/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for player's name
	fmt.Print("Enter your name: ")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)

	// Create a new player with the supplied name
	player := game.NewPlayer(playerName)

	// Went with a template here demonstrate using named placeholders.
	// While printf is more accepted in Go, I prefer templates for more complext outputs.
	// Primarily because it makes it an easier read for someone who has to review the code.
	// Like me in a week trying to figure out what in the world was I _thiking_!
	tmpl := "Welcome, {{.Name}}! You start with ${{.Bank}}.\n"

	// Template for the Welcome message.
	t, err := template.New("welcome").Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Execute the template with the player data and print the result to stdout.
	err = t.Execute(os.Stdout, player)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}

	// Create and shuffle the deck.
	deck := game.NewDeck()
	deck.ShuffleDeck()

	// Main game loop
	for player.Bank > 0 {
		fmt.Printf("\nYou have $%d. Place your bet: ", player.Bank)
		betInput, _ := reader.ReadString('\n')
		betInput = strings.TrimSpace(betInput)

		// Validate the bet
		bet, err := strconv.Atoi(betInput)
		if err != nil || bet <= 0 || bet > player.Bank {
			fmt.Println("Invalid bet. Please enter a valid amount.")
			continue
		}

		// Deal cards to player and dealer
		playerHand := make(game.Hand, 0)
		dealerHand := make(game.Hand, 0)

		playerHand = append(playerHand, deck.DealCard(), deck.DealCard())
		dealerHand = append(dealerHand, deck.DealCard(), deck.DealCard())

		fmt.Printf("\nYour hand: %v\n", playerHand)
		fmt.Printf("Total: %v\n\n", playerHand.CountCards())
		fmt.Printf("Dealer's visible card: %v\n\n", dealerHand[0])

		// Label for the outer for loop to allow a simple "break" anywhere in the process.
	playerTurn:
		for {
			// Determine the player's action
			fmt.Print("Do you want to (H)it or (S)tand ")
			actionInput, _ := reader.ReadString('\n')

			// I want to take a minute to point out the ToUpper call in the below line.
			// I _still_ run into code that doesn't normalize the input before testing it
			// against fixed assumptions. But I'm sure no one reading this makes that error.
			// Right?
			action := strings.TrimSpace(strings.ToUpper(actionInput))

			// Process the player's action
			switch action {
			case "H":
				playerHand = append(playerHand, deck.DealCard())
				currentCount := playerHand.CountCards()
				fmt.Printf("Your hand: %v\n", playerHand)
				fmt.Printf("Total: %v\n", currentCount)
				if currentCount > 21 {
					fmt.Println("Bust! You lose.")
					player.Bank -= bet
					player.Losses++
					break playerTurn
				}
			case "S": // Time to see who wins
				dealerCount := dealerHand.CountCards()
				playerCount := playerHand.CountCards()

				// Dealer must keep taking cards until they hit 17
				for dealerCount < 17 {
					dealerHand = append(dealerHand, deck.DealCard())
					dealerCount = dealerHand.CountCards()
				}

				fmt.Printf("Your hand: %v \nTotal: %d\n\n", playerHand, playerCount)
				fmt.Printf("Dealer's hand: %v \nTotal: %d\n\n", dealerHand, dealerCount)
				fmt.Print("Calculating winner...\n\n")

				// Wait for it!!!
				time.Sleep(3 * time.Second)

				if dealerCount > 21 || playerCount > dealerCount {
					fmt.Println("You win!")
					player.Bank += bet
					player.Wins++
				} else if playerCount < dealerCount {
					fmt.Println("You lost.")
					player.Bank -= bet
					player.Losses++
				} else {
					fmt.Println("Push")
				}
				break playerTurn
			default:
				fmt.Println("Juuust a little bit not valid. Please enter H or S.")
			}
		}

		// Check if the deck is running low and reshuffle if necessary
		if len(deck) < 10 {
			fmt.Println("Deck is running low. <shuffle, shuffle, shuffle>")
			deck = game.NewDeck()
			deck.ShuffleDeck()
		}

		// Ask if the player wants to play another hand
		fmt.Print("Do you want to play another hand? (Y/N): ")
		playAgainInput, _ := reader.ReadString('\n')
		playAgain := strings.TrimSpace(strings.ToUpper(playAgainInput))

		if playAgain != "Y" {
			break
		}
	}

	// Game over summary
	fmt.Printf("\nThank you for playing!\n")
	fmt.Printf("You won %d hands and lost %d hands.\n", player.Wins, player.Losses)
	fmt.Printf("Final bank balance: $%d\n", player.Bank)
}
