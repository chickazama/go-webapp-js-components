import Comment, { NewComment } from "./Comment.js";
customElements.define("comment-component", Comment);
const template = document.createElement("template");
template.innerHTML = `
<style>
    #comments-section {
        width: 90%;
        margin: auto;
    }
</style>
<div class="story">
    <slot name="title" class="title">Story Title</slot>
    <slot name="score" class="score">Score</slot>
    <slot name="comments" class="comments" id="comments-dropdown">Comments</slot>
</div>
`;

export default class Story extends HTMLElement {
    static observedAttributes = ["show"];
    shadowRoot;
    comments;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"});
        const content = template.content.cloneNode(true);
        this.shadowRoot.appendChild(content);
    }

    async hideCommentsAsync() {
        const host = this.getRootNode().host;
        const dropdown = host.shadowRoot.getElementById("comments-dropdown");
        dropdown.innerText = "Show";
        const commentsBox = host.shadowRoot.getElementById("comments-section");
        commentsBox.setAttribute("style", "display: none;");
        host.setAttribute("show", "false");
    }

    async showCommentsAsync() {
        const host = this.getRootNode().host;
        const dropdown = host.shadowRoot.getElementById("comments-dropdown");
        dropdown.innerText = "Hide";
        console.log(host.id);
        if (!host.comments) {
            host.comments = new Map();
            const res = await fetch(`/api/comments?id=${host.id}`);
            const data = await res.json();
            // console.log(data);
            const div = document.createElement("div");
            div.id="comments-section";
            for (const item of data) {
                const comment = NewComment(item);
                div.appendChild(comment);
                host.comments.set(comment.id, comment);
            }
            host.shadowRoot.appendChild(div);
        }
        const commentsBox = host.shadowRoot.getElementById("comments-section");
        commentsBox.setAttribute("style", "display: block;");
        host.setAttribute("show", "true");
    }

    connectedCallback() {
        console.log("Story component connected to DOM.");
    }

    disconnectedCallback() {
        console.log("Story component disconnected from DOM.");
    }

    attributeChangedCallback(name, oldValue, newValue) {
        // console.log(`Attribute ${name} has changed: ${oldValue} -> ${newValue}.`);
        const dropdown = this.shadowRoot.getElementById("comments-dropdown");
        switch (newValue) {
            case "true":
                dropdown.removeEventListener("click", this.showCommentsAsync);
                dropdown.addEventListener("click", this.hideCommentsAsync);
                break;
            case "false":
                if (oldValue) {
                    dropdown.removeEventListener("click", this.hideCommentsAsync);
                }
                dropdown.addEventListener("click", this.showCommentsAsync);
                break;
        }
    }
}