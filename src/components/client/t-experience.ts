import { LitElement, html, css } from "lit";
import { property, customElement } from "lit/decorators.js"

@customElement('t-experience')
export class TExperience extends LitElement {
  static styles = css`
  .experience_container {
    border: solid;
    border-width: 2px;
    padding-left: 16px;
    padding-right: 16px;
    padding-bottom: 16px;
    margin-top: 32px;
  }
  h2{  
    color: var(--md-sys-color-on-primary);
  }
  ::slotted([slot="description"]) {  
    color: var(--md-sys-color-on-primary-80);
  }
  span {
    color: var(--md-sys-color-on-primary-80);
    font-weight: normal;
    font-style: italic;
  }
  `
  @property({ type: String }) company = ""
  @property({ type: String }) position = ""
  @property({ type: String }) date = ""

  render() {
    return html`
      <div class="experience_container">
        <h2>${this.company} <span>${this.position}</span></h2>
        <p>${this.date}</p>
        <slot name="list"></slot>      
        <slot name="description"></slot>
      </div>
    `
  }
}
