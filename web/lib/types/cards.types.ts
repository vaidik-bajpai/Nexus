export interface Card {
    id: string;
    title: string;
    cover: string;
    coverSize?: string;
    completed: boolean;
    position: number;
    description?: string;
}