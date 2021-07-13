document.addEventListener('DOMContentLoaded', () => {
  // let loc = window.location;
  // let uri = 'ws:';
  //
  // if (loc.protocol === 'https:') {
  //   uri = 'wss:';
  // }
  //
  // uri += '//' + loc.host;
  // uri += loc.pathname + 'ws';
  let uri = 'ws://localhost:8080/ws'
  const input = document.getElementById('input');
  const output = document.getElementById('output');
  const btn = document.querySelector('.btn');
  const ws = new WebSocket(uri);

  ws.onopen = function() {
    console.log('Connected');
  }

  ws.onmessage = function(e) {
    output.innerHTML += e.data + '<br>';
  }

  btn.addEventListener('click', () => {
    ws.send(input.value);
    input.value = '';
  })
});
