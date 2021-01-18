/**
 * @typedef {import('./homeStore').HomeStore} HomeStore
 */

/**
 * @type {HomeStore}
 */
const homeStore = {
    data: {},
    setState(data) {
        Object.assign(this.data, data);
    }
}