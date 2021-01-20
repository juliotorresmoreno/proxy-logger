
(function () {
    Vue.component('Secure', {
        render(createElement) {
            if (!appStore.data.session.token) {
                return createElement('Auth');
            }
            return createElement('Layout');
        },
        mounted() {
            attachLinks(this.$el);
        },
        updated() {
            attachLinks(this.$el);
        }
    });
})();
