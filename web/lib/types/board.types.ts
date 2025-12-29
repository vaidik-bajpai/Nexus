import { List } from "./list.types";

export interface Board {
    id: string;
    name: string;
    background: string;
}

export interface BoardDetail extends Board {
    lists: List[];
}

export interface BoardMetadata {
    id: string;
    name: string;
    background: string;
    labels: BoardLabel[];
    members: BoardMember[];
}

export interface BoardMember {
    id: string;
    fullName: string;
    username: string;
    role: string;
    email: string;
}

