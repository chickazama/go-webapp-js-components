import Example, { NewExample } from "./components/Example.js";

customElements.define("example-component", Example);
window.addEventListener("load", async () => {
    console.log("Page loaded.");
    const main = document.getElementById("main");
    const example = NewExample("Welcome!");
    main.appendChild(example);
    
})