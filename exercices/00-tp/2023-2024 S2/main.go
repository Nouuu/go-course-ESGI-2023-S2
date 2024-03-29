package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Player int

const (
	Empty = iota
	Player1
	Player2
)

// Player1Color : Special characters for console colors - Red
const Player1Color = "\033[1;31m"

// Player2Color : Special characters for console colors - Blue
const Player2Color = "\033[1;34m" // Blue

// ResetColor : Special characters for console colors - Reset
const ResetColor = "\033[0m"

// Player1Name : This constant is used to display the player 1's name with the associated color
const Player1Name = Player1Color + "●" + ResetColor

// Player2Name : This constant is used to display the player 2's name with the associated color
const Player2Name = Player2Color + "●" + ResetColor

// PlayerNameMap : This map is used to associate a player with a string
var PlayerNameMap = map[Player]string{
	Player1: Player1Name,
	Player2: Player2Name,
	Empty:   ".", // Default value for an empty cell
}

const gameFilename = "game.json"

// BoardHeight : This variable is used to store the height of the board
var BoardHeight = flag.Int("height", 6, "Height of the board")

// BoardWidth : This variable is used to store the width of the board
var BoardWidth = flag.Int("width", 7, "Width of the board")

// Board : This type is used to represent the board
type Board [][]Player

// Game : This type is used to represent the game state
type Game struct {
	Board         Board
	CurrentPlayer Player
	Score         map[Player]int
}

func main() {
	flag.Parse()

	// Initialize the game
	game := initGame()

	for {
		game.start()
		if !askRestartGame() {
			break
		}
		// Reset the game board
		game.Board = newBoard(*BoardHeight, *BoardWidth)
		game.CurrentPlayer = Player1
	}

	fmt.Println("Game over")
}

// initGame : This function is used to initialize the game
func initGame() *Game {
	if askLoadGame() {
		game, err := loadGame()
		if err == nil {
			fmt.Println("Game loaded")
			return game
		} else {
			fmt.Println("Error while loading the game. Starting a new game")
		}
	}
	return newGame(*BoardHeight, *BoardWidth)
}

// newBoard : This function is used to create a new board
func newBoard(height, width int) Board {
	board := make(Board, height)
	for i := range board {
		board[i] = make([]Player, width)
	}
	return board
}

// newGame : This function is used to create a new game
func newGame(height, width int) *Game {
	return &Game{
		Board:         newBoard(height, width),
		CurrentPlayer: Player1,
		Score: map[Player]int{
			Player1: 0,
			Player2: 0,
		},
	}
}

// clearScreen : This function is used to clear the console
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// start : This function is used to start the game. It is the main loop of the game
func (game *Game) start() {
	for {
		// Save the game state
		err := game.save()
		if err != nil {
			fmt.Println("Error while saving the game:", err)
		}
		// Clear the console before printing the board
		clearScreen()
		// Print the board
		game.printBoard()
		// Check if the current player has won
		if winner := game.checkWinner(); winner != Empty {
			game.Score[winner]++
			fmt.Printf("Player %s won !\n", PlayerNameMap[winner])
			break
		}
		// Else, check if the board is full (draw)
		if game.isBoardFull() {
			fmt.Println("Draw!")
			break
		}
		// Let the current player play
		game.playerTurn()
		// Switch to the next player
		game.switchPlayer()
	}

}

// printBoard : This function is used to print the board
func (game *Game) printBoard() {
	fmt.Printf("Power 4 - Player %s (%d) vs Player %s (%d)\n", PlayerNameMap[Player1], game.Score[Player1], PlayerNameMap[Player2], game.Score[Player2])
	fmt.Println("To play, enter the column number where you want to play")
	fmt.Println("   1   2   3   4   5   6   7")
	fmt.Println(" +---+---+---+---+---+---+---+")
	for i := range *BoardHeight {
		fmt.Print(" |")
		for j := range *BoardWidth {
			fmt.Printf(" %s |", PlayerNameMap[game.Board[i][j]])
		}
		fmt.Println()
	}
	fmt.Println(" +---+---+---+---+---+---+---+")
}

// switchPlayer : This function is used to switch the current player
func (game *Game) switchPlayer() {
	if game.CurrentPlayer == Player1 {
		game.CurrentPlayer = Player2
	} else {
		game.CurrentPlayer = Player1
	}
}

// playerTurn: This function is used to handle a player's turn
func (game *Game) playerTurn() {
	validMove := false
	for !validMove {
		fmt.Printf("Player %s, enter the column number where you want to play: ", PlayerNameMap[game.CurrentPlayer])
		// Ask the player to play
		var column int
		// Check for invalid inputs
		inputs, err := fmt.Scanf("%d", &column)
		if err != nil || inputs != 1 || column < 1 || column > *BoardWidth {
			fmt.Println("Invalid input. Please enter a valid column number")
			continue
		}
		// Check if the column is full
		columnIndex := column - 1
		if game.isColumnFull(columnIndex) {
			fmt.Println("This column is full. Please enter another column number")
			continue
		}
		// Play the move
		validMove = game.playMove(columnIndex)
		if !validMove {
			fmt.Println("Invalid move. Please enter another column number")
		}
	}
}

func (game *Game) isColumnFull(column int) bool {
	return game.Board[0][column] != Empty
}

func (game *Game) playMove(column int) bool {
	for row := *BoardHeight - 1; row >= 0; row-- {
		if game.Board[row][column] == Empty {
			game.Board[row][column] = game.CurrentPlayer
			return true
		}
	}
	return false
}

// checkWinner : This function is used to check if a player has won the game
func (game *Game) checkWinner() Player {
	for i := range *BoardHeight {
		for j := range *BoardWidth {
			// Check row
			if j <= *BoardWidth-4 && isEqual(game.Board[i][j], game.Board[i][j+1], game.Board[i][j+2], game.Board[i][j+3]) {
				return game.Board[i][j]
			}
			// Check column
			if i <= *BoardHeight-4 && isEqual(game.Board[i][j], game.Board[i+1][j], game.Board[i+2][j], game.Board[i+3][j]) {
				return game.Board[i][j]
			}
			// Check diagonal (top-left to bottom-right)
			if i <= *BoardHeight-4 && j <= int(*BoardWidth)-4 && isEqual(game.Board[i][j], game.Board[i+1][j+1], game.Board[i+2][j+2], game.Board[i+3][j+3]) {
				return game.Board[i][j]
			}
			// Check diagonal (top-right to bottom-left)
			if i >= 3 && j <= *BoardWidth-4 && isEqual(game.Board[i][j], game.Board[i-1][j+1], game.Board[i-2][j+2], game.Board[i-3][j+3]) {
				return game.Board[i][j]
			}
		}
	}
	return Empty
}

func isEqual(a, b, c, d Player) bool {
	return a != Empty && a == b && a == c && a == d
}

// isBoardFull : This function is used to check if the board is full
func (game *Game) isBoardFull() bool {
	for _, cell := range game.Board[0] {
		if cell == Empty {
			return false
		}
	}
	return true
}

// save : This function is used to save the game state to a file
func (game *Game) save() error {
	data, err := json.MarshalIndent(game, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(gameFilename, data, 0666)
}

// loadGame : This function is used to load the game state from a file
func loadGame() (*Game, error) {
	data, err := os.ReadFile(gameFilename)
	if err != nil {
		return nil, err
	}
	var game Game
	err = json.Unmarshal(data, &game)
	if err != nil {
		return nil, err
	}
	*BoardHeight = len(game.Board)
	*BoardWidth = len(game.Board[0])
	return &game, nil
}

// askLoadGame : This function is used to ask the user if he wants to load a saved game
func askLoadGame() bool {
	fmt.Print("Do you want to load a saved game? (y/n): ")
	var answer string
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false
	}
	return strings.ToLower(strings.TrimSpace(answer)) == "y"
}

// askRestartGame : This function is used to ask the user if he wants to restart the game
func askRestartGame() bool {
	fmt.Print("Do you want to restart the game? (y/n): ")
	var answer string
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false
	}
	return strings.ToLower(strings.TrimSpace(answer)) == "y"
}
