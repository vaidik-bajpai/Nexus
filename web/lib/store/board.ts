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
    enrichCards: (cardId: string, fullCard: Partial<Card>) => void;
    addLabelToBoard: (label: any) => void;
}

export const useBoardStore = create<BoardState>()((set) => ({
    id: "",
    metadata: {
        id: "",
        name: "",
        background: "",
        labels: [],
        members: []
    },
    lists: [],
    cards: [],

    setMetadata: (metadata: BoardMetadata) => set({ metadata }),
    setCardsAndLists: (id: string, cards: Card[], lists: List[]) => set({ id, cards, lists }),
    enrichCards: (cardId: string, fullCard: Partial<Card>) => set((state) => {
        const cardIndex = state.cards.findIndex((card: Card) => card.id === cardId);
        if (cardIndex !== -1) {
            state.cards[cardIndex] = { ...state.cards[cardIndex], ...fullCard };
        }
        return { cards: [...state.cards] };
    }),
    addLabelToBoard: (label: any) => set((state) => ({
        metadata: {
            ...state.metadata,
            labels: [...(state.metadata.labels || []), label]
        }
    }))
}))