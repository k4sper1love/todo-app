// Service layer for user-related API interactions with Go backend

import {AddProfile, DeleteProfile, GetProfile, GetProfiles, UpdateProfile} from "../../wailsjs/go/main/App";
import {Profile} from "../types/profile";
import {mapProfile} from "../utils/profile";

export const fetchProfile = async (id: number): Promise<Profile | null> => {
    try {
        const data: any = await GetProfile(id);
        return mapProfile(data);
    } catch (error) {
        console.error("Failed to fetch profile:", error);
        return null;
    }
};

export const fetchProfiles = async (): Promise<Profile[]> => {
    try {
        const data: any = await GetProfiles();
        return data.map(mapProfile);
    } catch (error) {
        console.error("Failed to fetch profiles:", error);
        return [];
    }
}

export const createProfile = async(profile: Profile): Promise<number> => {
    try {
        return await AddProfile(profile.name);
    } catch (error){
        console.error("Failed to create profile:", error);
        return -1
    }
}

export const updateProfile = async(profile: Profile) => {
    try {
        await UpdateProfile(profile.id, profile.name);
    } catch(error) {
        console.error("Failed to update profile:", error);
    }
}

export const deleteProfile = async(id: number) => {
    try {
        await DeleteProfile(id);
    } catch (error) {
        console.error("Failed to delete profile:", error);
    }
}