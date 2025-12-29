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
    checklist?: Checklists[];
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