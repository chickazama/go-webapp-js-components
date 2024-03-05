const template = document.createElement("template");
template.innerHTML = `
<style>
    .eg {
        background-color: red;
    }
</style>
<div class="eg">
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

    async connectedCallback() {
        console.log("Example component connected to DOM. Fetching data...");
        const res = await fetch("/test");
        const body = await res.json();
        console.log(body);
    }

    disconnectedCallback() {
        console.log("Example component disconnected from DOM.");
    }
}

export function NewExample(str) {
    const example = new Example();
    const h1 = document.createElement("h1");
    h1.setAttribute("slot", "example");
    h1.innerText = str;
    example.appendChild(h1);
    return example;
}