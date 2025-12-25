import apiClient from "../apiClient";

export const createLabel = async (label: { boardID: string, name: string, color: string }) => {
    try {
        const response = await apiClient.post(`/boards/${label.boardID}/labels/create`, { name: label.name, color: label.color });
        return response.data;
    } catch (error) {
        console.error("Error creating label:", error);
        throw error;
    }
}

export const toggleLabel = async (label: { boardID: string, listID: string, cardID: string, labelID: string, type: "add" | "remove" }) => {
    try {
        const response = await apiClient.post(`/boards/${label.boardID}/lists/${label.listID}/cards/${label.cardID}/labels/toggle`, { type: label.type, label_id: label.labelID });
        return response.data;
    } catch (error) {
        console.error("Error toggling label:", error);
        throw error;
    }
}
