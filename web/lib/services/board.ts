import apiClient from "../apiClient";

export const createBoard = async (board: { name: string, background: string, visibility: string }) => {
    try {
        const response = await apiClient.post("/boards/create", board);
        return response.data;
    } catch (error) {
        console.error("Error creating board:", error);
        throw error;
    }
}