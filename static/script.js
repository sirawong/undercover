const ws = new WebSocket("ws://" + location.host + "/ws");

ws.onopen = function() {
    console.log("WebSocket connection established");
};

ws.onmessage = function(event) {
    const message = JSON.parse(event.data);
    const messages = document.getElementById('messages');
    const roleDiv = document.getElementById('role');
    const playersDiv = document.getElementById('players');
    
    if (message.type === "info") {
        const messageElement = document.createElement('div');
        messageElement.textContent = message.body;
        messages.appendChild(messageElement);
    } else if (message.type === "role") {
        roleDiv.textContent = "Your role is: " + message.body;
    } else if (message.type === "player_list") {
        playersDiv.innerHTML = "<h3>Current Players:</h3>";
        const players = JSON.parse(message.body);
        players.forEach(player => {
            const playerElement = document.createElement('div');
            playerElement.textContent = player;
            playersDiv.appendChild(playerElement);
        });
    }
};

ws.onclose = function() {
    console.log("WebSocket connection closed");
};

ws.onerror = function(error) {
    console.log("WebSocket error:", error);
};

function joinGame() {
    const nameInput = document.getElementById('nameInput').value;
    const message = {
        type: "join",
        body: nameInput
    };
    ws.send(JSON.stringify(message));
}

function startGame() {
    const message = {
        type: "start",
        body: "Start the game"
    };
    ws.send(JSON.stringify
