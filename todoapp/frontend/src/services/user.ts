// Service layer for user-related API interactions with Go backend

import {GetUsername, SetUsername} from "../../wailsjs/go/main/App";

// Fetch username
export const fetchUsername = async (): Promise<string> => {
    try {
        return await GetUsername();
    } catch (error) {
        console.error("Failed to fetch username:", error);
        return "";
    }
};

// Set username
export const setUsername = async (username: string) => {
    try {
        await SetUsername(username);
    } catch (error) {
        console.error("Failed to set username:", error);
    }
};