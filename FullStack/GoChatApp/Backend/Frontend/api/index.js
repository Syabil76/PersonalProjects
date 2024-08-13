// api/index.js


// LOGIC -> the function 'connect' essentialy listens to the endpoint of
// the server and listens for events like open, close, messages or errors
// the send msg function allows messages from front end to be sent to backend

var socket = new WebSocket("ws://localhost:8000/ws")

let connect = cb => {
    console.log("Attempting connection....");

    socket.onopen = () => {
        console.log("successfuly connected");
    };

    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    socket.onclose = event => {
        console.log("console connection closed", event);
    };

    socket.onerror = error => {
        console.log("error encountered", error);
    };
};

let sendMsg = msg => {
    console.log("sending msg", msg);
    socket.send(msg);
};

export { connect ,sendMsg };