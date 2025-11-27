import { create } from "zustand";
import { persist } from "zustand/middleware";

export const useUserStore = create<UserState>()(
    persist(
        (set) => ({
            user: null,
            setUser: (user) => set(() => ({ user: user })),
            clearUser: () => set(() => ({ user: null })),
        }),
        { name: 'user-storage' },
    ),
)