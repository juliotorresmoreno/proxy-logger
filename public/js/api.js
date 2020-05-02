
/**
 * @type {WebSocket}
 */
let ws;

const createConn = function () {
  const host = `${document.location.host}`;
  ws = new WebSocket(`ws://${host}/ws`, 'echo-protocol');
  ws.onmessage = async (message) => {
    try {
      const _message = JSON.parse(message.data);
      if (_message.type === "requests") {
        store.setLogs(_message.data.reverse())
        store.filter()
        store.mapColor()
        store.lastPage()
        if (!store.state.$request && store.state.$logs.length > 0)
          store.setRequest(store.state.$logs[0].ID);
        else
          this.$request = null;
        return
      }
      if (_message.type === "request") {
        store.setRequest(_message.data);
        store.mapColor()
      }
    } catch (error) {
      console.trace(error)
    }
  }
  ws.onerror = (error) => {
    console.log(error);
  }
  ws.onclose = () => {
    setTimeout(() => createConn(), 3000);
    console.log("conexión cerrada!");
  }
  ws.onopen = () => {
    console.log("conexión abierta!");
  }
}

createConn()

/**
 * 
 * @param {String} id 
 */
const get_request = async function (id) {
  return await fetch('/logs/http/' + id)
  /*ws.send({
    'type': 'get_request',
    'request': id
  })*/
}

const clear_logs = async function () {
  await fetch('/clear');
}