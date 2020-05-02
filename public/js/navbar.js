
(function () {
  const template = `
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <a class="navbar-brand" href="#">Home</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>

      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav mr-auto">
          <li v-bind:class="httpClass()" v-on:click='selectPageHTTP'>
            <a class="nav-link" href="#">
              HTTP
            </a>
          </li>
          <li v-bind:class="tcpClass()" v-on:click='selectPageTCP'>
            <a class="nav-link" href="#">
              TCP
            </a>
          </li>
          <li v-bind:class="driverClass()" v-on:click='selectPageDriver'>
            <a class="nav-link" href="#">
              Simulator
            </a>
          </li>
        </ul>
        <form class="form-inline my-2 my-lg-0">
          <input 
            v-model:value='search_text'
            v-on:change='signalChange'
            class="form-control mr-sm-2" type="search" 
            placeholder="Search" aria-label="Search">
          <button class="btn btn-outline-success my-2 my-sm-0" 
            v-on:click="$emit('search')" type="button">
            Search
          </button>&nbsp;&nbsp;
          <button class="btn btn-outline-success my-2 my-sm-0" 
            v-on:click="$emit('clear')" type="button">
            Clear
          </button>
        </form>
      </div>
    </nav>
  `;

  Vue.component('navbar', {
    props: [
      'search_text'
    ],
    data: () => store.state,
    methods: {
      signalChange: function (evt) {
        this.$emit("search_text_change", evt.target.value);
      },
      selectPageHTTP() {
        store.setTabPage('HTTP')
      },
      selectPageTCP() {
        store.setTabPage('TCP')
      },
      selectPageDriver() {
        store.setTabPage('DRIVER')
      },
      httpClass() {
        return {
          'nav-item': true,
          'active': this.$data.$tabPage === 'HTTP'
        }
      },
      tcpClass() {
        return {
          'nav-item': true,
          'active': this.$data.$tabPage === 'TCP'
        }
      },
      driverClass() {
        return {
          'nav-item': true,
          'active': this.$data.$tabPage === 'DRIVER'
        }
      }, 
    },
    template: template
  })
})();
