import apiClient from "../apiClient";

export const createList = async (list: { name: string, boardID: string }) => {
    try {
        const response = await apiClient.post(`boards/${list.boardID}/lists/create`, { name: list.name, position: 1 });
        return response.data;
    } catch (error) {
        console.error("Error creating list:", error);
        throw error;
    }
}
