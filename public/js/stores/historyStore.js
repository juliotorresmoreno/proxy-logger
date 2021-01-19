/**
 * @typedef {import('./historyStore').HistoryStore} HistoryStore
 */

/**
 * @type {HistoryStore}
 */
const historyStore = {
    data: {},
    setState(data) {
        Object.assign(this.data, data);
    }
}