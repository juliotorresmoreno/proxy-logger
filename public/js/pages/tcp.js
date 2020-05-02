(function() {
  const template = `
    <div class="card" style="margin-left: 15px;">
      TCP logs
    </div>
  `

  Vue.component('tcp', {
    props: [],
    data: () => store.state,
    methods: {
      
    },
    template
  })
})();
