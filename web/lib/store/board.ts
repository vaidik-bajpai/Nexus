import { create } from "zustand";
import { BoardMetadata } from "../types/board.types";
import { Card, Checklist } from "../types/cards.types";
import { List } from "../types/list.types";

interface BoardState {
    id: string;
    metadata: BoardMetadata;
    lists: List[];
    cards: Card[];

    checklists: Record<string, Checklist>;
    setMetadata: (metadata: BoardMetadata) => void;
    setCardsAndLists: (id: string, cards: Card[], lists: List[]) => void;
    enrichCards: (cardId: string, fullCard: Partial<Card>) => void;
    addLabelToBoard: (label: any) => void;
    upsertChecklist: (checklist: Checklist) => void;
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
    checklists: {},

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
    })),
    upsertChecklist: (checklist: Checklist) => set((state) => ({
        checklists: {
            ...state.checklists,
            [checklist.id]: checklist
        }
    }))
}))