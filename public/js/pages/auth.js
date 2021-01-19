
(function () {
    Vue.component('Auth', {
        data: {
            username: '',
            password: ''
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

                <form class='form'>
                    <h3>Sign In</h3>
                    <div class="input-group mb-3">
                        <input
                            type="text"
                            v-bind:value="$this.username"
                            class="form-control" />
                    </div>
                    <div class="input-group mb-3">
                        <input
                            type="text"
                            v-bind:value="$this.password"
                            class="form-control" />
                    </div>

                    <div style='display: inline-flex;'>
                        <button style='width: 100px' type="button" class="btn btn-primary">
                            Log In
                        </button>
                    </div>
                </form>
            </div>
        </div>
        `
    });
})();
