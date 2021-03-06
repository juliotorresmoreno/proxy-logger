
(function () {
    /**
     * @typedef {import('./auth').Credentials} Credentials
     * @typedef {import('../models/session').Session} Session
     */

    /**
     * @param {import('./auth').Credentials} data 
     */
    function validateSignIn(data) {
        if (!data.username) throw new ErrorUsernameNotFound();
        if (!data.password) throw new ErrorPasswordNotFound();
    }

    /** 
     * @param {Credentials} data
     * @returns {Promise<Session>}
     */
    async function signIn(credentials) {
        validateSignIn(credentials);
        const response = await fetch('/api/users/sign-in', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(credentials)
        });
        if (!response.ok) throw new ErrorUnauthorized();
        return response.json();
    }

    Vue.component('Auth', {
        data: () => ({
            username: '',
            password: ''
        }),
        methods: {
            async onSignIn(evt) {
                const { username, password } = this;
                const credentials = { username, password };
                const session = await signIn(credentials);
                appStore.setState({ session: session });
            }
        },
        template: `
            <div class="row" style='margin: 0'>
                <div class="col-md-4 offset-md-4">
                    <br />
                    <br />
                    <br />
                    <br />
                    <br />
                    <br />

                    <form v-on:submit.prevent="onSignIn" class='form'>
                        <h3>Sign In</h3>
                        <div class="input-group mb-3">
                            <input
                                type="text" class="form-control"
                                v-model.value="$data.username" />
                        </div>
                        <div class="input-group mb-3">
                            <input
                                type="password" class="form-control"
                                v-model.value="$data.password" />
                        </div>

                        <div style='display: inline-flex;'>
                            <button style='width: 100px' type="submit" class="btn btn-primary">
                                Log In
                            </button>
                        </div>
                    </form>
                </div>
            </div>
        `
    });
})();
