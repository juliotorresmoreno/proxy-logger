
(function () {
    window.addEventListener('load', function () {
        Vue.component('App', {
            template: '<Secure />'
        });

        window.app = new Vue({
            el: '#app',
            template: '<App />',
            data: {
                appStore: appStore.data,
                homeStore: homeStore.data,
                domainsStore: domainsStore.data,
                historyStore: historyStore.data
            },
            mounted() {
                const minPath = (location.protocol + '//' + location.host).length;
                appStore.setState({
                    route: location.href.substr(minPath)
                });
                attachLinks(this.$el);
            }
        });
    });
})();
