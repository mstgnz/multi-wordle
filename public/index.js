/*
## License
This project is licensed under the MIT Licence. Refer to https://github.com/mstgnz/multi-wordle/blob/main/LICENSE for more information.
*/
class MultiWordle {
    constructor() {
        this.room = null;
        this.objSize = 0;
        this.player = {}
        this.players = [];
        this.messages = [];
        this.isAnimate = false;
        this.game = document.getElementById("game")
        this.chat = document.getElementById("chat")
        this.form = document.getElementById("form")
        this.title = document.getElementById("title")
        this.input = document.getElementById("input")
        this.error = document.getElementById("error")
        this.wordle = document.getElementById("wordle")
        this.alphabet = document.getElementById("alphabet")
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

    socketProcess() {

        // on open
        this.socket.onopen = () => {
            this.send("login")
        }

        // on message
        this.socket.onmessage = (event) => {
            const response = JSON.parse(event.data)
            this.room = response.room
            this.players = response.players
            this.title.innerHTML = this.room && this.room.id ? this.room.id.toUpperCase(): ""
            console.log(response)
            switch (response.type) {
                case "login":
                    this.handleNewPlayer(response)
                    break
                case "animate":
                    this.handleAnimate(response)
                    break
                case "chat":
                    this.handleChat(response)
                    break
                case "name":
                    this.handleName(response)
                    break
                case "wordle":
                    this.handleWordle(response)
                    break
                case "disconnect":
                    this.handleDisconnect(response)
                    break
                case "error":
                    this.handleError(response)
                    break
                case "fatal":
                    this.handleFatal(response)
                    break
            }
        }

        // on close
        this.socket.onclose = () => {
            this.connected.style.display = "none"
            this.unconnected.style.display = "block"
        }

        // on error
        this.socket.onerror = (err) => {
            this.connected.style.display = "none"
            this.unconnected.style.display = "block"
        }
    }

    handleNewPlayer(response) {
        this.player = response.player
        this.initWordle()
        this.addPlayerToGameArea()
        this.handleChat(response)
    }

    handleAnimate(response) {
        const player = this.players.find((player) => player.name === response.player.name)
        player.position.x = response.player.position.x;
        player.position.y = response.player.position.y;
        this.animateElement(player)
    }

    handleChat(response) {
        this.messages = response.room.messages
        this.addMessageToChat(this.messages[this.messages.length -1])
        this.scrollTop()
    }

    handleError(response){
        this.error.innerHTML = response.message
        this.error.style.display = "block"
        setTimeout(function () {
            this.error.style.display = "none"
        },2000)
    }

    handleFatal(response){
        this.connected.style.display = "none"
        this.unconnected.style.display = "block"
        this.unconnected.innerHTML = response.message
    }

    handleName(response) {
        if (response.player.name !== this.player.name) {
            const player = this.players.find((player) => player.name === response.player.name)
            if(player){
                this.changeName(player, response.message)
                player.name = response.message
                this.handleChat(response)
            }
        }
    }

    handleWordle(response) {
        this.room = response.room
        this.handleChat(response)
        this.initWordle()
    }

    handleDisconnect(response) {
        this.players = this.players.filter((p) => p.name !== response.player.name);
        this.handleChat(response)
    }

    scrollTop() {
        setTimeout(() => {
            this.chat.scrollTop = this.chat.scrollHeight;
        })
    }

    animateElement(player) {
        const element = document.getElementById(player.name)
        if(element){
            const center = this.objSize / 2;
            element.style.left = (player.position.x - center) + "px";
            element.style.top = (player.position.y - center) + "px";
            this.isAnimate = false;
        }
    }

    send(type, message = "", position = {}) {
        this.socket.send(
            JSON.stringify({
                type: type,
                message: message,
                position: position,
                token: this.player && this.player.token ? this.player.token : ""
            })
        )
    }

    onMessage = (event) => {
        if (event.key === "Enter" && this.input.value.length) {
            this.send("chat", this.input.value)
            this.showBubble(this.player.name, this.input.value)
            this.checkCommand(this.input.value)
            this.input.value = "";
        }
    }

    onClickPlayer = (event) => {
        const element = document.getElementById(this.player.name)
        if(element){
            this.objSize = element.offsetWidth;
            const center = this.objSize / 2;
            this.player.position.x = event.offsetX - center;
            this.player.position.y = event.offsetY - center;
            if (!this.isAnimate) {
                this.isAnimate = true;
                this.send("animate","", {"x":event.offsetX - center, "y":event.offsetY - center})
                this.animateElement(this.player)
            }
        }
    }

    showBubble(name, message) {
        const element = document.getElementById(name)
        if(element){
            const messageElement = element.querySelector('.message')
            messageElement.style.display = 'block';
            messageElement.innerHTML = message;
            setTimeout(function () {
                messageElement.style.display = 'none';
            }, this.bubbleLifeTime(message))
        }
    }

    bubbleLifeTime(message) {
        const min = 500;
        const max = 3000;
        const msPerLetter = 40;
        let bubbleTime = min + message.length * msPerLetter;
        return bubbleTime > max ? max : bubbleTime;
    }

    checkCommand(command){
        command = command.split(" ")
        if(command[1]) {
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
                    if(this.room.len === command[1].length){
                        this.send("wordle", command[1])
                    }else{
                        alert("word count not matched")
                    }
                    break
            }
        }
        if(command[0] === ":change-bg"){
            fetch("https://source.unsplash.com/random/1920x1080").then((response) => {
                if (response.ok) {
                    document.body.style.backgroundImage = `url(${response.url})`
                }
            })
        }
    }

    changeName(player, name){
        const element = document.getElementById(player.name)
        if(element){
            const nameElement = element.querySelector('.name')
            nameElement.innerHTML = name
        }
    }

    addPlayerToGameArea() {
        this.players.forEach(player => {
            this.game.innerHTML += `<div class="circle" id="${player.name}" style="left:${player.position.x}px;top:${player.position.y}px; background-color: ${player.color}">
            <div class="relative">
                <span class="name">${player.name}</span>
                <div class="message"></div>
            </div>
        </div>`
        })
    }

    addMessageToChat(message) {
        this.chat.innerHTML += `<div class="item">
            <div class="content">
                <span>${message}</span>
            </div>
        </div>`;
    }

    // wordle layout initialization
    initWordle(){
        this.wordleBox.innerHTML = ""
        this.alphabet.innerHTML = ""
        // set wordle
        for (let i = 0; i < this.room.trial; i++) {
            const wordleRow = document.createElement("div");
            wordleRow.classList.add("wordle-row");
            for (let j = 0; j < this.room.len; j++) {
                const wordleItem = document.createElement("div");
                wordleItem.classList.add("wordle-item");
                const forecast = this.getLetter(i,j)
                if(forecast){
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

    chapter(array) {
        const pieceLength = Math.ceil(array.length / 3);
        const splitArray = [];
        for (let i = 0; i < array.length; i += pieceLength) {
            const part = array.slice(i, i + pieceLength);
            splitArray.push(part);
        }
        return splitArray;
    }

    getLetter(i,j){
        const forecasts = this.room.wordle.forecasts;
        if (i >= 0 && i < forecasts.length && j >= 0) {
            const forecastI = forecasts[i];
            if (j < forecastI.forecast.length) {
                return forecastI.forecast[j];
            }
        }
        return null
    }
}

new MultiWordle()