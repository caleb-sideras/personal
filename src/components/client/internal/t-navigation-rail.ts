import { css, PropertyValues } from 'lit';
import { customElement, property } from 'lit/decorators.js';
import { MdList } from '@material/web/list/list';
import { ListItem } from './internal/t-list';

@customElement('t-navigation-rail')
export class TNavigationRail extends MdList {

  static styles = [
    css`
      :host {
        width: auto !important;
      }
  `,
    ...MdList.styles
  ]

  @property({ type: String }) url = '';

  protected override updated(_changedProperties: PropertyValues<TNavigationRail>) {
    if (_changedProperties.has('url')) this.activateItemFromId(this.url);
  }

  removeFirstLastSlash(text: string) {
    return text.replace(/^\/|\/$/g, "");
  }

  activateItemFromId(href: string) {
    for (const item of this.items as ListItem[]) {
      if (this.removeFirstLastSlash(item.href) === this.removeFirstLastSlash(href)) {
        const activationEvent = new Event('request-activation', { bubbles: true, composed: true });
        Object.defineProperty(activationEvent, 'target', { value: item });
        this.listController.onRequestActivation(activationEvent);
        return;
      }
    }
  }
}