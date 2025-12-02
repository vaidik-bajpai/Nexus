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
