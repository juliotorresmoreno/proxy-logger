
import { Session } from '../models/session'


export interface AppData {
    route: string;
    session: Session
}

export interface AppStore {
    data: AppData,
    setState: (data: AppData) => void
}