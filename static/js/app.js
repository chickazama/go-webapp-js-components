import Example, { NewExample } from "./components/Example.js";
import Navbar from "./components/Navbar.js";

customElements.define("navbar-component", Navbar);
customElements.define("example-component", Example);
window.addEventListener("load", () => {
    console.log("Page loaded.");
    const main = document.getElementById("main");
    const example = NewExample("Hi :)");
    main.appendChild(example);
})