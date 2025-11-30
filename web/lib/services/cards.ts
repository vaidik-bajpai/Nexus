import apiClient from "../apiClient";

export const createCard = async (card: { title: string, listID: string, boardID: string, position: number }) => {
    try {
        const response = await apiClient.post(`/boards/${card.boardID}/lists/${card.listID}/cards/create`, { title: card.title, position: card.position });
        return response.data;
    } catch (error) {
        console.error("Error creating card:", error);
        throw error;
    }
}

export const updateCard = async (card: { cardID: string, listID: string, boardID: string, completed?: boolean, title?: string, cover?: string, coverSize?: string, position?: number, description?: string }) => {
    try {
        const response = await apiClient.put(`/boards/${card.boardID}/lists/${card.listID}/cards/${card.cardID}/update`, {
            completed: card.completed,
            title: card.title,
            cover: card.cover,
            coverSize: card.coverSize,
            position: card.position,
            description: card.description
        });
        return response.data;
    } catch (error) {
        console.error("Error updating card:", error);
        throw error;
    }
}

export const getCardDetail = async (card: { cardId: string, boardId: string, listId: string }) => {
    try {
        const response = await apiClient.get(`/boards/${card.boardId}/lists/${card.listId}/cards/${card.cardId}/detail`);
        return response.data;
    } catch (error) {
        console.error("Error getting card detail:", error);
        throw error;
    }
}