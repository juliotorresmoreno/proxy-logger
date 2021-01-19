
interface Session {
    token: string;
    profile: {
        username: string;
    }
}

export interface AppData {
    route: string;
    session: Session
}

export interface AppStore {
    data: AppData,
    setState: (data: AppData) => void
}