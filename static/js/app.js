import Example, { NewExample } from "./components/Example.js";

customElements.define("example-component", Example);
window.addEventListener("load", () => {
    console.log("Page loaded.");
})