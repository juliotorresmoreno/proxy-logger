
(function () {
    async function getData() {
        const response = await fetch('/api/history', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${appStore.data.session.token}`
            }
        });
        await parseError(response);
        return response.json();
    }

    Vue.component('History', {
        template: `
            <div class='row'>
                <div class='col'>history</div>
                <div class='col'>nada</div>
            </div>
        `,
        mounted: () => getData().catch(() => { })
    });
})()
