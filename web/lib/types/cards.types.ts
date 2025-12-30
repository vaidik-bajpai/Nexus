export interface Card {
    id: string;
    board_id: string;
    list_id: string;
    title: string;
    cover: string;
    coverSize: string;
    due: string;
    completed: boolean;
    position: number;
    labels: BoardLabel[];
    description?: string;
    member_ids?: string[];
    checklist_ids?: string[];
    archived?: boolean;
    start?: string;
}

export interface CardDetail {
    id: string;
    title: string;
    description: string;
    position: number;
    cover: string;
    coverSize: string;
    archived: boolean;
    completed: boolean;
    member_ids: string[];
    labels: CardLabel[];
    checklist_ids: string[];
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

export interface Checklist {
    id: string;
    name: string;
    position: number;
    checkItems: ChecklistItem[];
}

export interface ChecklistItem {
    id: string;
    name: string;
    completed: boolean;
    position: number;
}