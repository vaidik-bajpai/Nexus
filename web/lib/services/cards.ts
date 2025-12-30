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
        console.log(card)
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

export const toggleMemberToCard = async (member: { cardID: string, userID: string, listID: string, boardID: string }) => {
    try {
        const response = await apiClient.post(`/boards/${member.boardID}/lists/${member.listID}/cards/${member.cardID}/toggle-member`, { userID: member.userID });
        return response.data;
    } catch (error) {
        console.error("Error adding member to card:", error);
        throw error;
    }
}

export const addChecklistToCard = async (checklist: { cardID: string, name: string, listID: string, boardID: string }) => {
    try {
        const response = await apiClient.post(`/boards/${checklist.boardID}/lists/${checklist.listID}/cards/${checklist.cardID}/checklists/create`, { name: checklist.name });
        return response.data;
    } catch (error) {
        console.error("Error adding checklist to card:", error);
        throw error;
    }
}

export const getChecklist = async (checklist: { cardID: string, checklistID: string, listID: string, boardID: string }) => {
    try {
        const response = await apiClient.get(`/boards/${checklist.boardID}/lists/${checklist.listID}/cards/${checklist.cardID}/checklists/${checklist.checklistID}/detail`);
        return response.data;
    } catch (error) {
        console.error("Error getting checklist:", error);
        throw error;
    }
}

export const deleteChecklist = async (checklist: { cardID: string, checklistID: string, listID: string, boardID: string }) => {
    try {
        const response = await apiClient.delete(`/boards/${checklist.boardID}/lists/${checklist.listID}/cards/${checklist.cardID}/checklists/${checklist.checklistID}/delete`);
        return response.data;
    } catch (error) {
        console.error("Error deleting checklist:", error);
        throw error;
    }
}

export const addChecklistItem = async (item: { cardID: string, checklistID: string, listID: string, boardID: string, name: string }) => {
    try {
        const response = await apiClient.post(`/boards/${item.boardID}/lists/${item.listID}/cards/${item.cardID}/checklists/${item.checklistID}/items/create`, { name: item.name });
        return response.data;
    } catch (error) {
        console.error("Error adding checklist item:", error);
        throw error;
    }
}

export const deleteChecklistItem = async (item: { cardID: string, checklistID: string, listID: string, boardID: string, itemID: string }) => {
    try {
        const response = await apiClient.delete(`/boards/${item.boardID}/lists/${item.listID}/cards/${item.cardID}/checklists/${item.checklistID}/items/${item.itemID}/delete`);
        return response.data;
    } catch (error) {
        console.error("Error deleting checklist item:", error);
        throw error;
    }
}

export const updateChecklistItem = async (item: { cardID: string, checklistID: string, listID: string, boardID: string, itemID: string, name?: string, completed?: boolean }) => {
    try {
        const response = await apiClient.put(`/boards/${item.boardID}/lists/${item.listID}/cards/${item.cardID}/checklists/${item.checklistID}/items/${item.itemID}/update`, { name: item.name, completed: item.completed });
        return response.data;
    } catch (error) {
        console.error("Error updating checklist item:", error);
        throw error;
    }
}