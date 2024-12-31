# Multi Wordle

It's the same Wordle game, the way of playing doesn't change.

### Standard way of playing
- Each guess must be a correct 5-letter word. Press enter to submit.
- After each guess the colors of the boxes will change according to the closeness of your guess.

There is a two-player version with websocket and a scoring system. Each player takes turns guessing. The scoring system determines the winner.

### Features
- [x] There will be rooms for two people.
- [x] Although the rooms are for two players, a single player game can be played with the “:start” command. No second person is assigned to the room as soon as it is started.
- [x] At the beginning of the game, the number of words to be played with the language option and word lengths can be set.
- [x] 10 points for knowing the word.
- [x] 5 points for finding the correct letter in the word in the correct position.
- [x] 3 points for finding the letter in the word but misplacing it.
- [x] If a player finds the letter in the word and the location is wrong and the next player sees the letter and locates it correctly, then 2 points from 5-3.
- [x] if there is no letter in the word, no penalty for the first use but -1 point for the second use.
- [x] If a correctly placed letter is not used in the next guess, a penalty of -2 points is awarded.
- [x] If a word is used that is not in the game language, -2 points penalty. The word list is embedded in the project.
- [x] If the player whose turn it doesn't answer within the specified time, he/she is penalized -5 points and the turn passes to the next player.

## Getting Started

Follow the steps below to start the project locally or on a web server.

### Prerequisites

- Install the Go programming language locally.

### Installation

1. Clone this repository to a local directory:
    ```bash
    git clone https://github.com/mstgnz/multi-wordle
    cd multi-wordle
    ```

2. Run project
   ```bash
   make run
   ```
   OR
   ```bash
    docker compose up -d
    ```
   OR
    ```bash
    docker build -t multi-wordle:latest . && docker run -d --restart=always -p 3000:3000 --name=multi-wordle multi-wordle
    ```


### Usage
Launch the application in your browser and join with a random username.
Click anywhere on the game area to move your character in that direction.
Type a message in the text box at the bottom and press "Enter" to send a message.

#### commands

| command              | description                 |
|----------------------|-----------------------------|
| :start               | Start the Wordle Game       |
| :reset               | Reset the Wordle Game       |
| :wordle word         | Guess the Wordle            |
| :change-name newName | Change Nick Name            |
| :change-bg           | Change Random Background    |
| :change-bg newUrl    | Change  Background with Url |

### Contributing
This project is open-source, and contributions are welcome. Feel free to contribute or provide feedback of any kind.

### License
This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for more details.