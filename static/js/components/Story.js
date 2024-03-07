const template = document.createElement("template");
template.innerHTML = `
<style>
    .story {
        background-color: lightgreen;
    }
    .title {
        font-family: sans-serif;
    }
    .comments {
        color: orange;
    }
</style>
<div class="story">
    <slot name="title" class="title">Story Title</slot>
    <slot name="comments" class="comments" id="comments-dropdown">Comments</slot>
</div>
`;

export default class Story extends HTMLElement {
    shadowRoot;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"});
        const content = template.content.cloneNode(true);
        this.shadowRoot.appendChild(content);
    }

    connectedCallback() {
        console.log("Story component connected to DOM.");
        const dropdown = this.shadowRoot.getElementById("comments-dropdown");
        dropdown.addEventListener("click", () => {
            console.log(`${this.id} clicked.`);
        })
    }

    disconnectedCallback() {
        console.log("Story component disconnected from DOM.");
    }
}