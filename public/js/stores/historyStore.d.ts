
export interface HistoryData {
}

export interface HistoryStore {
    data: HistoryData,
    setState: (data: AppData) => void
}