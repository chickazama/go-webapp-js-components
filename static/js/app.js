import Example, { NewExample } from "./components/Example.js";

customElements.define("example-component", Example);
window.addEventListener("load", async () => {
    console.log("Page loaded.");
    const main = document.getElementById("main");
    const example = NewExample("Welcome!");
    main.appendChild(example);
    const socket = new WebSocket("ws://localhost:8192/ws");
    socket.onopen = () => {
        console.log("Socket connection opened.");
    }
    socket.onclose = () => {
        console.log("Socket connection closed.");
    }
    socket.onerror = (e) => {
        console.log(e);
    }
    socket.onmessage = (e) => {
        const p = document.createElement("p");
        p.innerText = e.data;
        main.prepend(p);
    }
})