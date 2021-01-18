/**
 * @typedef {import('./historyStore').DomainsStore} DomainsStore
 */

/**
 * @type {DomainsStore}
 */
const domainsStore = {
    data: {},
    setState(data) {
        Object.assign(this.data, data);
    }
}