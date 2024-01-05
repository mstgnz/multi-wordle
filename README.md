# Multi Wordle

It's the same Wordle game, the way of playing doesn't change.

### Standard way of playing
- Each guess must be a correct 5-letter word. Press enter to submit.
- After each guess the colors of the boxes will change according to the closeness of your guess.

There is only a two-player version with websocket and a scoring system. Each player takes turns guessing. The scoring system determines the winner.

### Features
- [ ] At the beginning of the game, the number of words to be played with the language option and word lengths can be set.  
- [ ] 5 points for finding the correct letter in the word.  
- [ ] 2 points for finding the letter in the word but misplacing it.  
- [ ] If a player finds the letter in the word and the location is wrong and the next player sees the letter and locates it correctly, then 3 points from 5-2.  
- [ ] if there is no letter in the word, no penalty for the first use but -1 point for the second use.  
- [ ] If the player whose turn it is does not answer within 10 seconds, it passes to the next player. -5 points penalty for repetition.  
- [ ] if a word that is not in the language used is used -2 points, if the same word is used again -4 points. (control with an api service) or there can be defined word sets. I haven't decided yet.

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

### Contributing
This project is open-source, and contributions are welcome. Feel free to contribute or provide feedback of any kind.

### License
This project is licensed under the Apache License. See the [LICENSE](https://github.com/mstgnz/multi-wordle/blob/main/LICENSE) file for more details.