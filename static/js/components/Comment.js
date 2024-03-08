const template = document.createElement("template");
template.innerHTML = `
<style>
    #comments-section {
        width: 90%;
        margin: auto;
    }
</style>
<div class="data">
    <slot name="author" class="author">Author</slot>
    <slot name="text" class="text">Text</slot>
    <div class="replies">
        <slot name="comments" class="comments" id="comments-dropdown">No replies</slot>
    </div>
</div>
`;
export default class Comment extends HTMLElement {
    static observedAttributes = ["show"];
    shadowRoot;
    comments;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"});
        const content = template.content.cloneNode(true);
        this.shadowRoot.appendChild(content);
    }
    connectedCallback() {
        console.log("Comment Added to DOM.");

    }
    disconnectedCallback() {
        console.log("Comment Removed from DOM.");
    }
    async hideCommentsAsync() {
        const host = this.getRootNode().host;
        const commentsBox = host.shadowRoot.getElementById("comments-section");
        commentsBox.setAttribute("style", "display: none;");
        host.setAttribute("show", "false");
    }

    async showCommentsAsync() {
        const host = this.getRootNode().host;
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

    attributeChangedCallback(name, oldValue, newValue) {
        console.log(`Attribute ${name} has changed: ${oldValue} -> ${newValue}.`);
        const dropdown = this.shadowRoot.getElementById("comments-dropdown");
        switch (newValue) {
            case "true":
                console.log("Showing");
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

export function NewComment(data) {
    console.log(data);
    const comment = new Comment();
    comment.id = data.id;
    comment.setAttribute("show", "false");
    const author = document.createElement("h5");
    author.setAttribute("slot", "author")
    author.innerText = data.by;
    const text = document.createElement("p");
    text.setAttribute("slot", "text");
    text.innerText = data.text;
    comment.appendChild(author);
    comment.appendChild(text);
    if (data.kids) {
        const replies = document.createElement("h5");
        replies.setAttribute("slot", "comments");
        replies.innerText = `Replies: ${data.kids.length}`;
        comment.appendChild(replies);
    }
    return comment;
}