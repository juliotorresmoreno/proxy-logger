(function () {
  let ws = null;
  $(window).on('load', async function () {
    var app = new Vue({
      el: '#app',
      data: () => store.state,
      methods: {
        search: function () {
          store.filter()
        },
        search_text_change: function (value) {
          store.setSearchText(value)
          store.filter()
        },
        clear: async function () {
          await clear_logs();
          store.setLogs([]);
          store.setState({ $request: null })
        }
      }
    })

    $('#app').css('display', 'block');
  })
})();
