/**
 * @typedef {import('./domainsStore').DomainsStore} DomainsStore
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