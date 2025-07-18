import { LitElement } from 'lit';
import { TemplateResult } from 'lit';

export declare class SortableExample extends LitElement {
    sortContainer: HTMLElement;
    title: string;
    value: string;
    items: SortableItem[];
    firstUpdated(): void;
    protected createRenderRoot(): this;
    render(): TemplateResult<1>;
}

declare interface SortableItem {
    name: string;
}

export { }
