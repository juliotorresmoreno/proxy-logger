
(function () {
  const template = `
  <div
    v-on:click="$emit('view', request)"
    style="cursor: pointer">
    <span v-model='color' v-bind:style="{ color: color }">
      {{request.Method}}
    </span>
    <span v-if='selected == request.ID' style='font-weight: bold'>
      {{request.URI.substr(0, 50)}}
    </span>
    <span v-if='selected != request.ID'>
      {{request.URI.substr(0, 50)}}
    </span> {{request.StatusCode}}
  </div>
  `;

  Vue.component('log', {
    model: {
      event: 'view'
    },
    props: ['request', 'color', 'selected'],
    methods: {

    },
    template: template
  });
})();
