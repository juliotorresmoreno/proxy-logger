

(function () {
  const template = `
    <nav aria-label="navigation">
      <ul class="pagination" style='margin: 0'>
        <li class="page-item">
          <a v-on:click="signalFirstChange" class="page-link" href="javascript: void(0)">
            first
          </a>
        </li>
        <li class="page-item">
          <a v-on:click="signalBackChange" class="page-link" href="javascript: void(0)">
            back
          </a>
        </li>
        <li class="page-item">
          <a v-on:click="signalNextChange" class="page-link" href="javascript: void(0)">
            next
          </a>
        </li>
        <li class="page-item">
          <a v-on:click="signalLastChange" class="page-link" href="javascript: void(0)">
            last
          </a>
        </li>
      </ul>
    </nav>
  `;
  Vue.component('navigation', {
    props: ['page','lastPage'],
    methods: {
      signalFirstChange: function () {
        this.$emit("page_change", 1);
      },
      signalBackChange: function () {
        let page = parseInt(this.page) - 1;
        if (page < 1) {
          page = 1;
        }
        this.$emit("page_change", page);
      },
      signalNextChange: function () {
        let page = parseInt(this.page) + 1;
        let lastPage = parseInt($(this.$el).attr('lastpage'))
        if (page > lastPage) {
          page = lastPage;
        }
        this.$emit("page_change", page);
      },
      signalLastChange: function () {
        this.$emit("page_change", $(this.$el).attr('lastpage'));
      }
    },
    template: template
  })
})();
