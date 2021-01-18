
(function () {
    window.addEventListener('load', function () {
        Vue.component('App', {
            template: '<Secure />'
        });

        window.app = new Vue({
            el: '#app',
            template: '<App />',
            data: () => ({
                appStore: appStore.data
            }),
            mounted() {
                const minPath = (location.protocol + '//' + location.host).length;
                appStore.setState({
                    route: location.href.substr(minPath)
                });
                const links = this.$el.querySelectorAll('a');
                links.forEach((link) => {
                    link.onclick = (evt) => {
                        evt.preventDefault();
                        links.forEach((currentLink) => {
                            let className = currentLink.className;
                            className = className.replace(/\sactive/, '');
                            currentLink.className = className;
                        });
                        link.className += ' active';
                        history.pushState(undefined, undefined, link.href);
                        appStore.setState({
                            route: link.href.substr(minPath)
                        });
                    }
                });
            }
        });
    });
})();
