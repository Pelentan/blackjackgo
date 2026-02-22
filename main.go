package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/yourusername/blackjackgo/game"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ask for player's name
	fmt.Print("Enter your name: ")
	playerName, _ := reader.ReadString('\n')
	playerName = strings.TrimSpace(playerName)

	// Create a new player
	player := game.NewPlayer(playerName)
	fmt.Printf("Welcome, %s! You start with $%d.\n", player.Name, player.Bank)

	// Create and shuffle the deck
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

		fmt.Printf("Your hand: %v\n", playerHand)
		fmt.Printf("Dealer's visible card: %v\n", dealerHand[0])

		// Player's turn
		for {
			fmt.Print("Do you want to (H)it, (S)tand, or (Q)uit? ")
			actionInput, _ := reader.ReadString('\n')
			action := strings.TrimSpace(strings.ToUpper(actionInput))

			switch action {
			case "H":
				playerHand = append(playerHand, deck.DealCard())
				fmt.Printf("Your hand: %v\n", playerHand)
				if game.CalculateScore(playerHand) > 21 {
					fmt.Println("Bust! You lose.")
					player.Bank -= bet
					player.Losses++
					break
				}
			case "S":
				dealerScore := game.CalculateScore(dealerHand)
				for dealerScore < 17 {
					dealerHand = append(dealerHand, deck.DealCard())
					dealerScore = game.CalculateScore(dealerHand)
				}

				playerScore := game.CalculateScore(playerHand)

				fmt.Printf("Your hand: %v (Total: %d)\n", playerHand, playerScore)
				fmt.Printf("Dealer's hand: %v (Total: %d)\n", dealerHand, dealerScore)

				if dealerScore > 21 || playerScore > dealerScore {
					fmt.Println("You win!")
					player.Bank += bet
					player.Wins++
				} else if playerScore < dealerScore {
					fmt.Println("You lose.")
					player.Bank -= bet
					player.Losses++
				} else {
					fmt.Println("It's a tie!")
				}
				break
			case "Q":
				fmt.Println("Thanks for playing!")
				return
			default:
				fmt.Println("Invalid action. Please enter H, S, or Q.")
			}
			// Check if the deck is running low and reshuffle if necessary
			if len(deck) < 10 {
				fmt.Println("Deck is running low. Reshuffling...")
				deck = game.NewDeck()
				deck.ShuffleDeck()
			}
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
	fmt.Printf("\nGame Over!\n")
	fmt.Printf("You won %d hands and lost %d hands.\n", player.Wins, player.Losses)
	fmt.Printf("Final bank balance: $%d\n", player.Bank)
}
