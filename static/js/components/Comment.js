const template = document.createElement("template");
template.innerHTML = `
<style>
    .data {
        background-color: cyan;
    }
</style>
<div class="data">
    <slot name="author" class="author">Author</slot>
    <slot name="text" class="text">Text</slot>
    <slot name="replies" class="replies">No replies</slot>
</div>
`;
export default class Comment extends HTMLElement {
    shadowRoot;
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
}

export function NewComment(data) {
    console.log(data);
    const comment = new Comment();
    comment.id = data.id;
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
        replies.setAttribute("slot", "replies");
        replies.innerText = `Replies: ${data.kids.length}`;
        comment.appendChild(replies);
    }
    return comment;
}