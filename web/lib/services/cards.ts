import apiClient from "../apiClient";

export const createCard = async (card: { title: string, listID: string, boardID: string }) => {
    try {
        const response = await apiClient.post(`/boards/${card.boardID}/lists/${card.listID}/cards/create`, { title: card.title });
        return response.data;
    } catch (error) {
        console.error("Error creating card:", error);
        throw error;
    }
}

export const updateCard = async (card: { cardID: string, listID: string, boardID: string, completed?: boolean, title?: string }) => {
    try {
        const response = await apiClient.put(`/boards/${card.boardID}/lists/${card.listID}/cards/${card.cardID}/update`, {
            completed: card.completed,
            title: card.title
        });
        return response.data;
    } catch (error) {
        console.error("Error updating card:", error);
        throw error;
    }
}
