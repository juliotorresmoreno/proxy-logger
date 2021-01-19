
(function () {
    const routes = {
        '/': 'Home',
        '/history': 'History',
        '/domains': 'Domains'
    };
    Vue.component('Router', {
        render(createElement) {
            const route = routes[appStore.data.route] || 'NotFound'
            return createElement(route)
        }
    });
})();
