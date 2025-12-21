import { create } from "zustand";
import { BoardMetadata } from "../types/board.types";
import { Card } from "../types/cards.types";
import { List } from "../types/list.types";

interface BoardState {
    id: string;
    metadata: BoardMetadata;
    lists: List[];
    cards: Card[];

    setMetadata: (metadata: BoardMetadata) => void;
    setCardsAndLists: (id: string, cards: Card[], lists: List[]) => void;
}

export const useBoardStore = create<BoardState>()((set) => ({
    id: "",
    metadata: {
        id: "",
        name: "",
        background: "",
        labels: []
    },
    lists: [],
    cards: [],

    setMetadata: (metadata: BoardMetadata) => set({ metadata }),
    setCardsAndLists: (id: string, cards: Card[], lists: List[]) => set({ id, cards, lists })
}))