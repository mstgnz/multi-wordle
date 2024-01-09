/*
## License
This project is licensed under the MIT Licence. Refer to https://github.com/mstgnz/multi-wordle/blob/main/LICENSE for more information.
*/
class MultiWordle {
    constructor() {
        this.init = {}
        this.room = null;
        this.objSize = 0;
        this.player = {}
        this.players = [];
        this.response = {}
        this.messages = [];
        this.isAnimate = false;
        this.game = document.getElementById("game")
        this.chat = document.getElementById("chat")
        this.form = document.getElementById("form")
        this.title = document.getElementById("title")
        this.total = document.getElementById("total")
        this.input = document.getElementById("input")
        this.error = document.getElementById("error")
        this.wordle = document.getElementById("wordle")
        this.wordleToken = localStorage.getItem("wordle-token")
        this.alphabet = document.getElementById("alphabet")
        this.initForm = document.getElementById("init-form")
        this.connected = document.getElementById("connected")
        this.wordleBox = document.getElementById("wordle-box")
        this.unconnected = document.getElementById("unconnected")
        this.game.addEventListener("click", this.onClickPlayer)
        this.input.addEventListener("keypress", this.onMessage)
        this.form.addEventListener("submit", (event) => {
            event.preventDefault()
        })
        this.socket = new WebSocket("ws://localhost:3000/ws")
        this.socketProcess()
    }

    socketProcess = () => {

        // on open
        this.socket.onopen = () => {
            if(this.wordleToken){
                this.send("login")
            }else{
                this.initForm.addEventListener("submit", (event) => {
                    event.preventDefault()
                    this.init.lang = this.initForm.querySelector('#room-lang').value.toString()
                    this.init.limit = parseInt(this.initForm.querySelector('#set-count').value)
                    this.init.length = parseInt(this.initForm.querySelector('#wordle-length').value)
                    this.init.trial = parseInt(this.initForm.querySelector('#wordle-trial').value)
                    this.send("login")
                })
            }
        }

        // on message
        this.socket.onmessage = (event) => {
            this.response = JSON.parse(event.data)
            if(this.response.room){
                this.room = this.response.room
                this.title.innerHTML = this.room && this.room.name ? this.room.name.toUpperCase() : ""
            }
            if(this.response.players){
                this.players = this.response.players
            }
            console.log(this.response)
            switch (this.response.type) {
                case "login":
                    this.handleLogin()
                    break
                case "animate":
                    this.handleAnimate()
                    break
                case "chat":
                    this.handleChat()
                    break
                case "name":
                    this.handleName()
                    break
                case "wordle":
                    this.handleWordle()
                    break
                case "total":
                    this.handleTotal()
                case "disconnect":
                    this.handleDisconnect()
                    break
                case "error":
                    this.handleError()
                    break
                case "fatal":
                    this.handleFatal()
                    break
            }
            this.input.focus()
        }

        // on close
        this.socket.onclose = () => {
            this.close()
        }

        // on error
        this.socket.onerror = (err) => {
            this.close()
        }
    }

    handleLogin = () => {
        this.player = this.response.player
        localStorage.setItem("wordle-token", this.player.token)
        this.initWordle()
        this.addPlayerToGameArea()
        this.handleChat()
        this.initForm.style.display = "none"
        this.connected.style.display = "block"
    }

    handleAnimate = () => {
        const player = this.players.find((player) => player.name === this.response.player.name)
        player.position.x = this.response.player.position.x;
        player.position.y = this.response.player.position.y;
        this.animateElement(player)
    }


    handleChat = () => {
        this.messages = this.response.room.messages
        this.addMessageToChat(this.messages[this.messages.length - 1])
        this.scrollTop()
    }

    handleError = (error) => {
        this.error.innerHTML = error ? error : this.response.message
        this.error.style.display = "block"
        setTimeout(function () {
            this.error.style.display = "none"
        }.bind(this), 5000)
    }

    handleFatal = () => {
        this.close()
        this.unconnected.innerHTML = this.response.message
    }

    handleName = () => {
        if (this.response.player.name !== this.player.name) {
            const player = this.players.find((player) => player.name === this.response.player.name)
            if (player) {
                this.changeName(player, this.response.message)
                player.name = this.response.message
                this.handleChat(this.response)
            }
        }
    }

    handleWordle = () => {
        this.handleChat()
        this.changeScore()
        this.initWordle()
    }

    handleTotal = () => {
        this.total.innerHTML = this.response.message
    }

    handleDisconnect = () => {
        if(this.response.player){
            this.players = this.players.filter((p) => p.name !== this.response.player.name);
            this.handleChat()
        }
    }

    scrollTop = () => {
        setTimeout(() => {
            this.chat.scrollTop = this.chat.scrollHeight;
        })
    }

    animateElement = (player) => {
        const element = document.getElementById(player.name)
        if (element) {
            const center = this.objSize / 2;
            element.style.left = (player.position.x - center) + "px";
            element.style.top = (player.position.y - center) + "px";
            this.isAnimate = false;
        }
    }

    send = (type, message = "", position = {}) => {
        this.socket.send(
            JSON.stringify({
                type: type,
                init: this.init,
                message: message,
                position: position,
                token: this.wordleToken ?? ""
            })
        )
    }

    onMessage = (event) => {
        if (event.key === "Enter" && this.input.value.length) {
            const value = this.input.value.trim()
            if(value.length !== 0){
                this.send("chat", value)
                this.showBubble(this.player.name, value)
                this.checkCommand(value)
            }
            this.input.value = "";
        }
    }

    onClickPlayer = (event) => {
        const element = document.getElementById(this.player.name)
        if (element) {
            this.objSize = element.offsetWidth;
            const center = this.objSize / 2;
            this.player.position.x = event.offsetX - center;
            this.player.position.y = event.offsetY - center;
            if (!this.isAnimate) {
                this.isAnimate = true;
                this.send("animate", "", {"x": event.offsetX - center, "y": event.offsetY - center})
                this.animateElement(this.player)
            }
        }
    }

    showBubble = (name, message) => {
        const element = document.getElementById(name)
        if (element) {
            const messageElement = element.querySelector('.message')
            messageElement.style.display = 'block';
            messageElement.innerHTML = message;
            setTimeout(function () {
                messageElement.style.display = 'none';
            }, this.bubbleLifeTime(message))
        }
    }

    bubbleLifeTime = (message) => {
        const min = 500;
        const max = 3000;
        const msPerLetter = 40;
        let bubbleTime = min + message.length * msPerLetter;
        return bubbleTime > max ? max : bubbleTime;
    }

    checkCommand = (command) => {
        command = command.trim().split(" ")
        if (command[1]) {
            switch (command[0]) {
                case ":change-name":
                    const newName = command[1].substring(0, 5)
                    this.send("name", newName)
                    this.changeName(this.player, newName)
                    this.player.name = command[1];
                    break
                case ":change-bg":
                    const newBg = command[1];
                    document.body.style.backgroundImage = `url(${newBg})`
                    break
                case ":wordle":
                    if (this.room.len === command[1].length) {
                        this.send("wordle", command[1])
                    } else {
                        this.handleError(`The number of letters in this word "${command[1]}" does not match.`)
                    }
                    break
            }
        }
        if (command[0] === ":change-bg") {
            fetch("https://source.unsplash.com/random/1920x1080").then((response) => {
                if (response.ok) {
                    document.body.style.backgroundImage = `url(${response.url})`
                }
            })
        }
    }

    changeName = (player, name) => {
        const element = document.getElementById(player.name)
        if (element) {
            const nameElement = element.querySelector('.name')
            nameElement.innerHTML = name
        }
    }

    changeScore = () => {
        const element = document.getElementById(this.response.player.name)
        if (element) {
            const nameElement = element.querySelector('.score')
            nameElement.innerHTML = "S: " + this.response.player.score
        }
    }

    addPlayerToGameArea = () => {
        this.game.innerHTML = ""
        this.players.forEach(player => {
            this.game.innerHTML += `<div class="circle" id="${player.name}" style="left:${player.position.x}px;top:${player.position.y}px; background-color: ${player.color}">
            <div class="relative">
                <span class="name">${player.name}</span>
                <span class="score">S: ${player.score}</span>
                <div class="message"></div>
            </div>
        </div>`
        })
    }

    addMessageToChat = (message) => {
        this.chat.innerHTML += `<div class="item">
            <div class="content">
                <span>${message}</span>
            </div>
        </div>`;
    }

    // wordle layout initialization
    initWordle = () => {
        this.wordleBox.innerHTML = ""
        this.alphabet.innerHTML = ""
        // set wordle
        for (let i = 0; i < this.room.trial; i++) {
            const wordleRow = document.createElement("div");
            wordleRow.classList.add("wordle-row");
            for (let j = 0; j < this.room.len; j++) {
                const wordleItem = document.createElement("div");
                wordleItem.classList.add("wordle-item");
                const forecast = this.getLetter(i, j)
                if (forecast) {
                    wordleItem.textContent = String.fromCharCode(forecast.letter)
                    wordleItem.style.backgroundColor = forecast.color
                }
                wordleRow.appendChild(wordleItem);
            }
            this.wordleBox.appendChild(wordleRow);
        }
        // set alphabet
        const letters = this.chapter(this.room.wordle.alphabet)
        for (let i = 0; i < letters.length; i++) {
            const alphabetRow = document.createElement("div")
            alphabetRow.classList.add("alphabet-row")
            for (let j = 0; j < letters[i].length; j++) {
                const alphabetItem = document.createElement("div")
                alphabetItem.classList.add("alphabet-item")
                alphabetItem.textContent = String.fromCharCode(letters[i][j].letter)
                alphabetItem.style.backgroundColor = letters[i][j].color
                alphabetRow.appendChild(alphabetItem)
            }
            this.alphabet.appendChild(alphabetRow);
        }
    }

    chapter = (array) => {
        const pieceLength = Math.ceil(array.length / 3);
        const splitArray = [];
        for (let i = 0; i < array.length; i += pieceLength) {
            const part = array.slice(i, i + pieceLength);
            splitArray.push(part);
        }
        return splitArray;
    }

    getLetter = (i, j) => {
        const forecasts = this.room.wordle.forecasts;
        if (i >= 0 && i < forecasts.length && j >= 0) {
            const forecastI = forecasts[i];
            if (j < forecastI.forecast.length) {
                return forecastI.forecast[j];
            }
        }
        return null
    }

    close = () => {
        this.connected.style.display = "none"
        this.unconnected.style.display = "block"
    }
}

new MultiWordle()