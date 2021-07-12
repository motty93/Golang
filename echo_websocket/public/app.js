document.addEventListener('DOMContentLoaded', () => {
  let loc = window.location;
  let uri = 'ws:';

  if (loc.protocol === 'https:') {
    uri = 'wss:';
  }

  uri += '//' + loc.host;
  uri += loc.pathname + 'ws';

  ws = new WebSocket(uri);
  ws.onopen = function() {
    console.log('Connected');
  }

  ws.onmessage = function(e) {
    const el = document.getElementById('output');
    el.innerHTML += e.data + '<br>';
  }

  const btn = document.querySelector('.btn');
  btn.addEventListener('click', () => {
    const value = document.getElementById('input').value;
    ws.send(value);
    document.getElementById('input').value = '';
  })
});
