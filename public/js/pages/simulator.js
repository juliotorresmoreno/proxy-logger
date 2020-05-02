
(function() {
  const template = `
    <div class="card" style="margin-left: 15px;">
      <div class="card-body" style="padding-bottom: 10px;">
        <div class="row">
          <div class="col col-md-5" style="display: flex; flex-direction: column;">
            hola
          </div>
          <div class="col col-md-7" style="display: flex; flex-direction: column;">
            mundo
          </div>
        </div>
      </div>
    </div>
  `

  Vue.component('simulator', {
    props: [],
    data: () => store.state,
    methods: {
      
    },
    template
  })
})();
