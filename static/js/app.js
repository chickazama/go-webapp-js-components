import Navbar from "./components/Navbar.js";

customElements.define("navbar-component", Navbar);

window.addEventListener("load", () => {
    console.log("Page loaded.");
    const main = document.getElementById("main");
    const testNav = new Navbar();
    main.appendChild(testNav);
    const testBtn = document.createElement("button");
    testBtn.innerText = "Click Me!";
    testBtn.type = "button";
    testBtn.addEventListener("click", () => {
        console.log("Button clicked.");
    })
    main.appendChild(testBtn);
})