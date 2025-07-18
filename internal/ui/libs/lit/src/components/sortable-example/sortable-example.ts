import { LitElement, html } from 'lit'
import { customElement, property, query } from 'lit/decorators.js'
import Sortable from 'sortablejs'

interface SortableItem {
  name: string
}

@customElement('sortable-example')
export class SortableExample extends LitElement {
  @query('#sortable-container')
  sortContainer!: HTMLElement

  @property({ type: String }) title: string = ''
  @property({ type: String }) value: string = ''
  @property({ type: Array }) items: SortableItem[] = []

  firstUpdated() {
    new Sortable(this.sortContainer, {
      animation: 150,
      ghostClass: 'opacity-25',
      onEnd: (evt) => {
        this.value = `Moved from ${evt.oldIndex} to ${evt.newIndex}`
        this.dispatchEvent(
          new CustomEvent('change', {
            detail: `Moved from ${evt.oldIndex} to ${evt.newIndex}`,
          }),
        )
      },
    })
  }

  protected createRenderRoot() {
    return this
  }

  render() {
    console.log(this)
    return html`
      <div class="flex flex-col gap-2 p-4 border border-primary rounded">
          <div class="text-lg">${this.title}: <span class="font-bold">${this.value}</span></div>
          <div>Open your console to see event results</div>
          <div id="sortable-container" class="flex flex-col gap-4">
            ${this.items?.length > 0 && this.items.map(
              (item) => html` <div class="bg-primary text-primary-content p-4 rounded-box">${item.name}</div> `,
            )}
        </div>
      </div>
    `
  }
}
