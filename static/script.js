const ws = new WebSocket("ws://" + location.host + "/ws");

ws.onopen = function() {
    console.log("WebSocket connection established");
};

ws.onmessage = function(event) {
    const message = JSON.parse(event.data);
    const messages = document.getElementById('messages');
    const messageElement = document.createElement('div');
    messageElement.textContent = message.body;
    messages.appendChild(messageElement);
};

ws.onclose = function() {
    console.log("WebSocket connection closed");
};

ws.onerror = function(error) {
    console.log("WebSocket error:", error);
};

function joinGame() {
    console.log("Join Game button clicked");
    const nameInput = document.getElementById('nameInput').value;
    const message = {
        type: "join",
        body: {
            name: nameInput
        }
    };
    ws.send(JSON.stringify(message));
}