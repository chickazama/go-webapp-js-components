import Example from "./components/Example.js";
import Navbar from "./components/Navbar.js";

customElements.define("navbar-component", Navbar);
customElements.define("example-component", Example);
window.addEventListener("load", () => {
    console.log("Page loaded.");
    const main = document.getElementById("main");
    const eg = new Example();
    const h1 = document.createElement("h1");
    h1.setAttribute("slot", "example");
    h1.innerText = "Whaddamegunnado";
    eg.appendChild(h1);
    main.appendChild(eg);
    // const testNav = new Navbar();
    // main.appendChild(testNav);
    // const testBtn = document.createElement("button");
    // testBtn.innerText = "Click Me!";
    // testBtn.type = "button";
    // testBtn.addEventListener("click", () => {
    //     console.log("Button clicked.");
    // })
    // main.appendChild(testBtn);
})