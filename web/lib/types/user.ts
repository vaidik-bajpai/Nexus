interface storeUser {
    username: string;
    email: string;
    id: string;
}

interface UserState {
    user: storeUser | null
    setUser: (user: storeUser) => void
    clearUser: () => void
}