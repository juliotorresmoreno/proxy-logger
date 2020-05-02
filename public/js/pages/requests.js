
(function() {
  const template = `
    <div class="card" style="margin-left: 15px;">
      <div class="card-body" style="padding-bottom: 10px;">
        <div class="row" style='display: flex; flex-direction: row; flex: 1'>
          <div class="col col-md-5" style="display: flex; flex-direction: column; flex: 1">
            <div style="overflow-y: scroll; flex: 1">
              <ol style="list-style: none; padding-left: 0;">
                <log 
                  v-on:view='view' v-for="(item, index) in $data.$filtered_logs" 
                  v-bind:request="item"
                  v-bind:selected="get_request_id()"
                  v-bind:key="index"
                  v-bind:color='item.color'>
                </log>
              </ol>
            </div>
            <div>
              <navigation 
                v-on:page_change='page_change' 
                v-bind:page='$data.$page' 
                v-bind:lastPage='$data.$lastPage' />
            </div>
          </div>
          <div class="col col-md-7" style="flex: 1; display: flex; flex-direction: column;">
            <div v-if='request'>
              <button v-if='replyDisable()'  v-on:click='reply' 
                disabled style='float: right' 
                class="btn btn-secondary">
              Reply
              </button>
              <button v-if='!replyDisable()' v-on:click='reply' 
                style='float: right' class="btn btn-primary">
                Reply
              </button>
              <pre>Fecha/hora: {{ request.Time }}</pre>
            </div>
            <div v-if='request' style='flex: 1; display: flex; flex-direction: column;'>
                <pre style="margin: 0;">{{ request.RawRequest }}</pre>
                <br />
                <pre style="margin: 0;">{{ request.RawResponse }}</pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  `

  Vue.component('requests', {
    props: ['request'],
    data: () => store.state,
    methods: {
      get_request_id: function () {
        return this.$data.$request ? this.$data.$request.ID : null;
      },
      reply: async function () {
        if (!this.$data.$request) return;

        let raw = this.$data.$request.RawRequest;
        let tmp = raw.split('\r\n\r\n');
        let head = tmp[0].split('\r\n')
        let body = tmp[1]
        let method = head[0].split(' ')[0];
        let url = head[0].split(' ')[1];
        let headers = {};
        for (let i = 1; i < head.length; i++) {
          const element = head[i].split(': ');
          if (element[0] !== 'Host' || !headers[element[0]])
            headers[element[0]] = element[1];
        }
        url = url.replace(/http(s)?:\/\/[^/]+/, '')
        url = 'http://' + (headers.Host || headers.host) + url
        /**
         * @type {RequestInit}
         */
        const opts = { url, method, headers }
        if (body) opts.body = body;
        await fetch('/request', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(opts)
        })
      },
      page_change: async function (page) {
        store.setPage(page);
        if (this.$data.$logs.length > 0) {
          store.setRequest(this.$data.$logs[0].ID);
          store.mapColor()
        }
      },
      view: async function (request) {
        if (!request) return;
        if (this.$request !== request)
          await store.setRequest(request.ID)
        // else this.request = null;
      },
      replyDisable() {
        if (!store.state.$request) return 'disabled';
        return store.state.$request.Method === 'CONNECT' ? 'disabled': ''
      }
    },
    template
  })
})();
