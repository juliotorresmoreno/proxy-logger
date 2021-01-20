
(function () {
    Vue.component('Secure', {
        render(createElement) {
            if (!appStore.data.session.token) {
                return createElement('Auth');
            }
            return createElement('Layout');
        },
        updated() {
            attachLinks(this.$el);
        }
    });
})();
