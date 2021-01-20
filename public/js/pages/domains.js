
(function () {
    async function getData() {
        const response = await fetch('/api/domains', {
            method: 'GET',
            headers: {
                'Authorization': `Bearer ${appStore.data.session.token}`
            }
        });
        await parseError(response);
        return response.json();
    }

    Vue.component('Domains', {
        template: `
            <div>Domains</div>
        `,
        mounted: () => getData().catch(() => { })
    })
})()
