export interface Card {
    id: string;
    title: string;
    cover: string;
    coverSize?: string;
    completed: boolean;
    position: number;
    description?: string;
}

export interface CardDetail {
    id: string;
    title: string;
    description: string;
    position: number;
    cover: string;
    archived: boolean;
    completed: boolean;
    members: CardMember[];
    labels: CardLabel[];
    checklist: Checklists[];
}

export interface CardMember {
    userID: string;
    avatar: string;
    username: string;
    email: string;
    isCardMember: boolean;
}

export interface CardLabel {
    labelID: string;
    name: string;
    color: string;
}

export interface Checklists {
    title: string;
    items: ChecklistItem[];
}

export interface ChecklistItem {
    id: string;
    title: string;
    done: boolean;
    userID: string;
}