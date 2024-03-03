const template = document.createElement("template");
template.innerHTML = `
    <div>
        <slot name="example">Example Text</slot>
    </div>
`;

export default class Example extends HTMLElement {
    shadowRoot;
    constructor() {
        super();
        this.shadowRoot = this.attachShadow({mode: "open"});
        const tmpl = template.content.cloneNode(true);
        this.shadowRoot.appendChild(tmpl);
    }

    connectedCallback() {
        console.log("Example component connected to DOM.");
    }

    disconnectedCallback() {
        console.log("Example component disconnected from DOM.");
    }
}