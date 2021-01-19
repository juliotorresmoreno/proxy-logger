
(function () {
    Vue.component('Layout', {
        template: `
            <div>
                <Header />
                <br />
                <div class="container">
                    <Router />
                    <Footer />
                </div>
            </div>
        `,
    });
}) ();
