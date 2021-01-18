
/**
 * @typedef {import('./appStore').AppStore} AppStore
 */

/**
 * @type {AppStore}
 */
const appStore = {
    data: {
        route: '/',
        session: {
            token: '',
            profile: {
                username: ''
            }
        }
    },
    setState(data) {
        Object.assign(this.data, data);
    }
}