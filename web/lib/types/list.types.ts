import { Card } from "./cards.types";

export interface List {
    id: string;
    name: string;
    cards: Card[];
}