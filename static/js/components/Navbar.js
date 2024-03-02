export default class Navbar extends HTMLElement {
    constructor() {
        super();
    }
    
    connectedCallback() {
        console.log("Navbar component added to DOM.");
    }

    disconnectedCallback() {
        console.log("Navbar component removed from DOM.");
    }
}