import {GetUsername, SetUsername} from "../../wailsjs/go/main/App";

export const fetchUsername = async (): Promise<string> => {
    try {
        return await GetUsername();
    } catch (error) {
        console.error("Failed to fetch username:", error);
        return "";
    }
};


export const setUsername = async (username: string) => {
    try {
        await SetUsername(username)
    } catch (error) {
        console.error("Failed to set username:", error)
    }
}